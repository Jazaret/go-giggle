package model

import (
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Task struct {
	TaskID      int    `json:"TaskID"`
	User        string `json:"User"`
	Description string `json:"Description"`
	Priority    int    `json:"Priority"`
}

func GetTasks() ([]Task, error) {
	// Read from DynamoDB
	var taskList []Task
	result, err := db.Scan(nil)
	log.Println("scanned")
	if err != nil {
		log.Printf("Failed to Scan: %s\n", err.Error())
		return nil, err
	}
	for _, i := range result.Items {
		task := Task{}
		dynamodbattribute.UnmarshalMap(i, &task)
		taskList = append(taskList, task)
		log.Printf("raw %+v\n", i)
		log.Printf("task %+v\n", task)
	}
	return taskList, nil
}

func addTasks() {

}

func updateTask() {

}

func deleteTask() {

}

func emailTasks() {

}

func sanitizeTask() {

}

func validateTask() {

}

func sortTasks() {

}

func mergeTasks() {

}
