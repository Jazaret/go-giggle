package model

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db DB

type DB interface {
	Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
}

type dynamoDbClient struct {
	db        *dynamodb.DynamoDB
	tableName *string
}

func (s dynamoDbClient) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if input == nil {
		input = &dynamodb.ScanInput{
			TableName: s.tableName,
		}
	}
	log.Println("scan set")
	return s.db.Scan(input)
}

func InitDatabase(region string, tableName *string) {
	if session, err := session.NewSession(&aws.Config{Region: &region}); err != nil {
		log.Printf("Failed to connect to AWS NewSession: %s\n", err.Error())
		panic(err.Error())
	} else {
		// Create DynamoDB client
		db = &dynamoDbClient{dynamodb.New(session), tableName}
		log.Println("db set")
	}
}
