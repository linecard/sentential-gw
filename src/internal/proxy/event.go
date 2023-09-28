package proxy

import (
	"io"
	"net/http"

	"github.com/wheegee/sentential-gw/internal/util"

	"github.com/aws/aws-lambda-go/events"
)

// Build and return an API Gateway v2 request event from a given HTTP request.
func BuildEvent(r *http.Request) (events.APIGatewayV2HTTPRequest, error) {
	var cookies []string
	for _, cookie := range r.Cookies() {
		cookies = append(cookies, cookie.String())
	}

	headers := make(map[string]string)
	for name, values := range r.Header {
		headers[name] = values[0]
	}

	queryStringParameters := make(map[string]string)
	for name, value := range r.URL.Query() {
		queryStringParameters[name] = value[0]
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		util.Log.Error("error reading request body", "error", err)
	}

	return events.APIGatewayV2HTTPRequest{
		Version:               "2.0",
		RouteKey:              "$default",
		RawPath:               r.URL.Path,
		RawQueryString:        r.URL.RawQuery,
		Cookies:               cookies,
		Headers:               headers,
		QueryStringParameters: queryStringParameters,
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method:   r.Method,
				Path:     r.URL.Path,
				SourceIP: r.RemoteAddr,
			},
		},
		Body: string(body),
	}, nil
}
