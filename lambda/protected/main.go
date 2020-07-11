package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       "Something went wrong :(",
		}, nil
	}

	cc := lc.ClientContext

	fmt.Println("client")
	fmt.Println(cc.Client)

	fmt.Println("env")
	fmt.Println(cc.Env)

	fmt.Println("custom")
	fmt.Println(cc.Custom)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, wait",
	}, nil
}

func main() {
	lambda.Start(handler)
}
