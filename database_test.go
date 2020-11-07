package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

var (
	dbContent = []Promocode{}
	mock      = &mockDynamoDBClient{}
	db        = DB{Client: mock}
)

func (m *mockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	jbyte, _ := json.Marshal(dbContent)
	av, err := dynamodbattribute.MarshalMap(&DynamoDBRecord{
		ID:   "test",
		Data: string(jbyte),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}

	r := dynamodb.GetItemOutput{
		Item: av,
	}
	// fmt.Printf("dynamo: %+v\n", r)

	return &r, nil
}

func (m *mockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	item := DynamoDBRecord{}
	if err := dynamodbattribute.UnmarshalMap(input.Item, &item); err != nil {
		return nil, errors.New("failed to unmarshal dynamodb record")
	}
	err := json.Unmarshal([]byte(item.Data), &dbContent)
	return nil, err
}

func TestGet(t *testing.T) {
	t.Run("Get promocode from empty db", func(t *testing.T) {
		dbContent = []Promocode{}
		shop := AdidasShopID

		promo, err := db.Get(&shop, nil)
		if promo != nil || err.Error() != ErrorPromocodeNotFound {
			t.Logf("Expected nil, got: %+v\n error expected: %s, got: %s",
				promo, ErrorPromocodeNotFound, err.Error())
			t.Fail()
		}
	})

	t.Run("Get promocode from db", func(t *testing.T) {
		dbContent = []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Data:   adidasPromocode,
			},
			Promocode{
				ShopID: ReebokShopID,
				Data:   reebokPromocode,
			},
		}
		shop := AdidasShopID

		promo, err := db.Get(&shop, nil)
		if err != nil {
			t.Logf("Expected error: nil, got: %v", err)
			t.Fail()
		} else if promo.Data != adidasPromocode || promo.ShopID != AdidasShopID {
			t.Logf("Expected adidas promocode, got: %v", *promo)
			t.Fail()
		}
	})
}

func TestList(t *testing.T) {
	t.Run("List promocodes from empty db", func(t *testing.T) {
		dbContent = []Promocode{}
		mock := &mockDynamoDBClient{}
		db := DB{Client: mock}

		promocodes, err := db.List()
		if len(promocodes) != 0 || err != nil {
			t.Logf("Expected: nil, got: %s\n Error expected nil, got: %v", promocodes, err)
			t.Fail()
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add promocode to empty db", func(t *testing.T) {
		dbContent = []Promocode{}
		mock := &mockDynamoDBClient{}
		db := DB{Client: mock}
		promo := Promocode{
			ShopID: AdidasShopID,
			Data:   adidasPromocode,
		}

		err := db.Add(&promo)
		if err != nil {
			t.Logf("Expected error: nil, got: %v", err)
			t.Fail()
		} else if !reflect.DeepEqual(dbContent[0], promo) {
			t.Logf("Expected: %+v, got: %+v", promo, dbContent[0])
			t.Fail()
		}
	})

	t.Run("Rewrite existing promocode in db", func(t *testing.T) {
		dbContent = []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Data:   adidasPromocode,
			},
			Promocode{
				ShopID: ReebokShopID,
				Data:   reebokPromocode,
			},
		}
		expected := []Promocode{
			Promocode{
				ShopID: ReebokShopID,
				Data:   reebokPromocode,
			},
			Promocode{
				ShopID: AdidasShopID,
				Data:   adidasPromocode2,
			},
		}

		mock := &mockDynamoDBClient{}
		db := DB{Client: mock}
		promo := Promocode{
			ShopID: AdidasShopID,
			Data:   adidasPromocode2,
		}

		err := db.Add(&promo)
		if err != nil {
			t.Logf("Expected error: nil, got: %v", err)
			t.Fail()
		} else if !reflect.DeepEqual(dbContent, expected) {
			t.Logf("Expected: %+v, got: %+v", expected, dbContent)
			t.Fail()
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete promocode from empty db", func(t *testing.T) {
		dbContent = []Promocode{}
		mock := &mockDynamoDBClient{}
		db := DB{Client: mock}
		promo := Promocode{
			ShopID: AdidasShopID,
			Data:   adidasPromocode,
		}

		err := db.Delete(&promo)
		if err != nil {
			t.Logf("Error expected nil, got: %v", err)
			t.Fail()
		}
	})

	t.Run("Delete promocode", func(t *testing.T) {
		dbContent = []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Data:   adidasPromocode,
			},
			Promocode{
				ShopID: ReebokShopID,
				Data:   reebokPromocode,
			},
		}
		expected := []Promocode{
			Promocode{
				ShopID: ReebokShopID,
				Data:   reebokPromocode,
			},
		}
		mock := &mockDynamoDBClient{}
		db := DB{Client: mock}
		promo := Promocode{
			ShopID: AdidasShopID,
			Data:   adidasPromocode,
		}

		err := db.Delete(&promo)
		if err != nil {
			t.Logf("Error expected nil, got: %v", err)
			t.Fail()
		} else if !reflect.DeepEqual(dbContent, expected) {
			t.Logf("Expected: %+v, got: %+v", expected, dbContent)
			t.Fail()
		}
	})
}
