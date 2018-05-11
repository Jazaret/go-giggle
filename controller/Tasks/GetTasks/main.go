package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jazaret/go-giggle/model"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	region    = os.Getenv("AWS_REGION")
	tableName = aws.String(os.Getenv("TABLE_NAME"))
)

func init() {
	log.Println("Calling init")
	log.Printf("Region: %s, Table %s\n", region, *tableName)
	model.InitDatabase(region, tableName)
}

// HandlerGetTasks is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func HandlerGetTasks(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	log.Print("Request = " + request.RequestContext.RequestID)

	tasks, err := model.GetTasks()

	if err != nil {
		log.Panic(err)
		panic(err)
	}
	outputTasks, _ := json.Marshal(tasks)

	response := events.APIGatewayProxyResponse{
		Body:       string(outputTasks),
		StatusCode: 200,
	}

	log.Printf("%+v\n", response)

	return response, nil

}

func main() {
	lambda.Start(HandlerGetTasks)
}
