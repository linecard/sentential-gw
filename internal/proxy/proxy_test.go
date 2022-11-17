package proxy

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type mockLambdaAPI func(ctx context.Context, input *lambda.InvokeInput, opts ...func(*lambda.Options)) (*lambda.InvokeOutput, error)

func (m mockLambdaAPI) Invoke(ctx context.Context, input *lambda.InvokeInput, opts ...func(*lambda.Options)) (*lambda.InvokeOutput, error) {
	return m(ctx, input, opts...)
}

func TestInvoke(t *testing.T) {
	cases := []struct {
		client  func(t *testing.T) LambdaAPI
		payload []byte
		expect  *lambda.InvokeOutput
	}{
		{
			client: func(t *testing.T) LambdaAPI {
				return mockLambdaAPI(func(ctx context.Context, input *lambda.InvokeInput, opts ...func(*lambda.Options)) (*lambda.InvokeOutput, error) {
					t.Helper()
					if input.Payload == nil {
						t.Fatal("expected payload to not be nil")
					}

					return &lambda.InvokeOutput{Payload: []byte("{\"sentential\": \"true\"}")}, nil
				})
			},
			payload: []byte("{\"RawPath\": \"/sentential\"}"),
			expect:  &lambda.InvokeOutput{Payload: []byte("{\"sentential\": \"true\"}")},
		},
	}

	for _, c := range cases {
		content, err := Invoke(context.TODO(), c.client(t), c.payload)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if string(content.Payload) != string(c.expect.Payload) {
			t.Errorf("expected %v, got %v", c.expect.Payload, content.Payload)
		}
	}
}
