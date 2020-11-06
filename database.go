package common

import (
	"encoding/json"
	"errors"
	"math/rand"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type DB struct {
	Client    dynamodbiface.DynamoDBAPI
	TableName string
	Prefix    string
}

// Get - reads and return a promocode from DB
func (db *DB) Get(shopId *PromocodeShopID, promoType *PromocodeType) (*Promocode, error) {
	client := Client{DynamoDB: db.Client}

	// read Promocodes from DynamoDB
	jstr, err := client.DynamoDBGet(db.TableName, db.Prefix+PromoBotDBName)
	if err != nil {
		if err.Error() != ErrorDynamoDBIDNotFound { //get was unsuccessfull
			log.Errorf("can not read from DynamodDB: %s", err.Error())
			return nil, err
		}

		log.Debug("DynamoDB item not found")
		jstr = `[]`
	}

	// unmarshall DB content to Promocodes
	var dbPromocodes []Promocode
	if err := json.Unmarshal([]byte(jstr), &dbPromocodes); err != nil {
		log.Errorf("can not unmarshal promocodes DynamoDB string: %s error: %s", jstr, err.Error())
		return nil, err
	}
	log.Debugf("Read from database: %+v", dbPromocodes)

	// collect promocodes for particular shop to slice
	var shopPromocodes []Promocode
	for _, p := range dbPromocodes {
		if shopId != nil && p.ShopID == *shopId && promoType != nil && p.Type == *promoType {
			shopPromocodes = append(shopPromocodes, p)
		}
	}

	// return nil if not found
	if len(shopPromocodes) == 0 {
		log.Debugf("promocode for shop: %v and type: %v not found", shopId, promoType)
		return nil, errors.New(ErrorPromocodeNotFound)
	}

	// choose promocode from slice
	n := rand.Intn(len(shopPromocodes))
	return &shopPromocodes[n], nil
}

// Add - merges promocode to DB
func (db *DB) Add(promo *Promocode) error {

	return nil
}

// Delete - removes a promocode from DB
func (db *DB) Delete(promo *Promocode) error {

	return nil
}

// List - lists all promocodes
func (db *DB) List() ([]Promocode, error) {
	var promocodes []Promocode

	return promocodes, nil
}
