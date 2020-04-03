package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"

	"encoding/json"
	"fmt"
	"os"
)

func getEngine(url string) *Engine {
	var e = Engine{url}
	return &e
}

func main() {

	var engine = getEngine("https://pingass.scalefast.ninja/api/sites")

	var data = engine.GetJSON()

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling Pinguer request")
		os.Exit(1)
	}

	regions := []string{
		"eu-north-1",
		"ap-south-1",
		"eu-west-3",
		"eu-west-2",
		"eu-west-1",
		"ap-northeast-2",
		"ap-northeast-1",
		"sa-east-1",
		"ca-central-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"eu-central-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
	}

	var client *lambda.Lambda

	for _, regionName := range regions {
		client = GetLambdaClient(regionName)
		client.Invoke(&lambda.InvokeInput{InvocationType: aws.String("Event"), FunctionName: aws.String("pinger-lambda-production-main"), Payload: payload})
	}

	fmt.Println("Done")

}
