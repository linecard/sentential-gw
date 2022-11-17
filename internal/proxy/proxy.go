package proxy

import (
	"context"
	"encoding/json"
	"net/http"
	"sentential-gw/internal/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type LambdaAPI interface {
	Invoke(context.Context, *lambda.InvokeInput, ...func(*lambda.Options)) (*lambda.InvokeOutput, error)
}

// Emit a Lambda event and return the payload as an HTTP response.
func Handle(w http.ResponseWriter, r *http.Request) {
	util.Log.Info("request", "method", r.Method, "path", r.URL.Path)
	ctx := r.Context()

	event, err := BuildEvent(r)
	if err != nil {
		util.Log.Error("error building event", "error", err)
	}

	payload, err := json.Marshal(event)
	if err != nil {
		util.Log.Error("error marshalling event", "error", err)
	}
	util.Log.Debug("lambda payload constructed", "payload", string(payload))

	client := LambdaAPI(lambda.New(*util.LambdaConfig))
	invoke, err := Invoke(ctx, client, payload)
	if err != nil {
		util.Log.Error("error invoking function", "error", err)
	}

	var response events.APIGatewayV2HTTPResponse
	err = json.Unmarshal(invoke.Payload, &response)
	if err != nil {
		util.Log.Error("error unmarshalling response", "error", err)
	}

	w.Write([]byte(response.Body))
	for name, value := range response.Headers {
		w.Header().Set(name, value)
	}
}

// Invoke a Lambda function and return the payload.
func Invoke(ctx context.Context, api LambdaAPI, payload []byte) (*lambda.InvokeOutput, error) {
	result, err := api.Invoke(
		ctx,
		&lambda.InvokeInput{
			FunctionName: aws.String("function"),
			Payload:      payload,
		},
	)
	if err != nil {
		return nil, err
	}
	util.Log.Debug("lambda response", "response", string(result.Payload))

	return result, nil
}
