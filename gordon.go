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

	var clientEU = GetLambdaClient("eu-west-1")
	var clientUS = GetLambdaClient("us-east-1")

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling Pinguer request")
		os.Exit(1)
	}

	clientEU.Invoke(&lambda.InvokeInput{InvocationType: aws.String("Event"), FunctionName: aws.String("Pinguer"), Payload: payload})
	clientUS.Invoke(&lambda.InvokeInput{InvocationType: aws.String("Event"), FunctionName: aws.String("Pinguer"), Payload: payload})

	fmt.Println("Done")

}
