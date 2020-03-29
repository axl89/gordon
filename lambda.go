package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

//GetLambdaClient returns a lambda client
func GetLambdaClient(region string) *lambda.Lambda {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	client := lambda.New(sess, &aws.Config{Region: aws.String(region)})

	return client
}

func callPinger(client *lambda.Lambda) {

}
