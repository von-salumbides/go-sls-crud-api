package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/von-salumbides/go-sls-crud-api/utils/logger"
)

type Database struct {
	connection *dynamodb.DynamoDB
	isLogMode  bool
}

type Interface interface {
}

func NewAdapter() Interface {
	return nil
}

func (db *Database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

func (db *Database) FindAll() {

}

func (db *Database) FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		logger.PANIC("Marshal Failure", err)
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       conditionParsed,
	}
	return db.connection.GetItem(input)
}

func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) {
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		logger.PANIC("Marshal Failure", err)
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

func (db *Database) Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		logger.PANIC("Marshal Failure", err)
		return nil, err
	}
	input := &dynamodb.DeleteItemInput{
		Key:       conditionParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.DeleteItem(input)
}
