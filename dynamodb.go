package common

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Client provides AWS API client
type Client struct {
	DynamoDB dynamodbiface.DynamoDBAPI
}

// DynamoDBRecord provides DynamoDB record data type
type DynamoDBRecord struct {
	ID   string
	Data string
}

// Get reads data record from DynamoDB
func (c *Client) DynamoDBGet(tableName, id string) (string, error) {
	log.Debugf("DynamoDB get table: %s, id: %s", tableName, id)

	result, err := c.DynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {S: aws.String(id)},
		},
	})
	if err != nil {
		return "", err
	}

	item := DynamoDBRecord{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return "", errors.New("failed to unmarshal dynamodb record")
	}

	if item.ID == "" {
		return "", errors.New("id not found")
	}

	log.Debugf("DynamoDB read data: %s", item.Data)
	return item.Data, nil
}

// Put reads data record from DynamoDB
func (c *Client) DynamoDBPut(tableName, id, data string) error {
	log.Debugf("DynamoDB put table: %s, id: %s", tableName, id)

	amap, err := dynamodbattribute.MarshalMap(&DynamoDBRecord{
		ID:   id,
		Data: data,
	})
	if err != nil {
		return err
	}

	_, err = c.DynamoDB.PutItem(&dynamodb.PutItemInput{
		Item:      amap,
		TableName: aws.String(tableName),
	})
	if err != nil {
		return err
	}

	return nil
}
