package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/netlify/gotrue/models"
)

type netlify struct {
	user models.User `json:"user",omitempty`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       "Something went wrong :(",
		}, nil
	}

	cc := lc.ClientContext

	netlifyJSON, _ := base64.StdEncoding.DecodeString(cc.Custom["netlify"])
	netlify := netlify{}
	netlifyUser := json.Unmarshal([]byte(netlifyJson), &netlify)

	if netlify.user == nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Don't know you bro :/",
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hello, %s", netlify.user.Email),
	}, nil
}

func main() {
	lambda.Start(handler)
}
