// snippet-comment:[These are tags for the AWS doc team's sample catalog. Do not remove.]
// snippet-sourceauthor:[Doug-AWS]
// snippet-sourcedescription:[Runs a Lambda function.]
// snippet-keyword:[AWS Lambda]
// snippet-keyword:[Invoke function]
// snippet-keyword:[Go]
// snippet-sourcesyntax:[go]
// snippet-service:[lambda]
// snippet-keyword:[Code Sample]
// snippet-sourcetype:[full-example]
// snippet-sourcedate:[2018-03-16]
/*
 Copyright 2010-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

 This file is licensed under the Apache License, Version 2.0 (the "License").
 You may not use this file except in compliance with the License. A copy of the
 License is located at

 http://aws.amazon.com/apache2.0/

 This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
 OF ANY KIND, either express or implied. See the License for the specific
 language governing permissions and limitations under the License.
*/

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

	var engine = getEngine("https://hackalog.scalefast.ninja/50.json")

	var data = engine.GetJSON()

	var clientEU = GetLambdaClient("eu-west-1")
	var clientUS = GetLambdaClient("us-east-1")

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling Pinguer request")
		os.Exit(1)
	}

	//result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("Pinguer"), Payload: payload})
	clientEU.Invoke(&lambda.InvokeInput{InvocationType: aws.String("Event"), FunctionName: aws.String("Pinguer"), Payload: payload})
	clientUS.Invoke(&lambda.InvokeInput{InvocationType: aws.String("Event"), FunctionName: aws.String("Pinguer"), Payload: payload})

	fmt.Println("Done")

}
