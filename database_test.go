package common

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

//
type mockDynamoDBClientEmpty struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamoDBClientEmpty) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	dbContent := []Promocode{}

	av, err := dynamodbattribute.MarshalMap(dbContent)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}

	r := dynamodb.GetItemOutput{
		Item: av,
	}

	return &r, nil
}

func TestGet(t *testing.T) {
	t.Run("Get promocode from empty db", func(t *testing.T) {
		mock := &mockDynamoDBClientEmpty{}
		db := DB{Client: mock}
		shop := AdidasShopID

		promo, err := db.Get(&shop, nil)
		if promo != nil || err.Error() != ErrorPromocodeNotFound {
			t.Logf("Expected nil, got: %+v\n error expected: %s, got: %s",
				promo, ErrorPromocodeNotFound, err.Error())
			t.Fail()
		}
	})

	t.Run("List promocodes from empty db", func(t *testing.T) {
		mock := &mockDynamoDBClientEmpty{}
		db := DB{Client: mock}

		promocodes, err := db.List()
		if len(promocodes) != 0 || err != nil {
			t.Logf("Expected: nil, got: %s\n Error expected nil, got: %v", promocodes, err)
			t.Fail()
		}
	})
}
