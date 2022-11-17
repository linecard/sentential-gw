package util

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/hashicorp/go-hclog"
)

var (
	Log          hclog.Logger
	Port         string
	LambdaConfig *lambda.Options
)

// Initialize the configuration.
func init() {
	level := GetEnv("SNTL_LOG", "INFO")
	Log = hclog.New(&hclog.LoggerOptions{
		Level: hclog.LevelFromString(level),
	})
	Port = fmt.Sprintf(":%s", GetEnv("PORT", "8081"))
	LambdaConfig = &lambda.Options{
		Region: GetEnv("AWS_REGION", "us-west-2"),
		EndpointResolver: lambda.EndpointResolverFromURL(
			GetEnv("LAMBDA_ENDPOINT", "http://localhost:9000"),
		),
	}
}
