package common

import (
	"encoding/json"
	"errors"
	"math/rand"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	log "github.com/sirupsen/logrus"
)

type DB struct {
	Client    dynamodbiface.DynamoDBAPI
	TableName string
	Prefix    string
}

// Get - reads and return a promocode from DB
func (db *DB) Get(shopId *PromocodeShopID, promoType *PromocodeType) (*Promocode, error) {
	dbPromocodes, err := db.read()
	if err != nil {
		log.Errorf("can not read from DynamodDB: %s", err.Error())
		return nil, err
	}
	log.Debugf("Read from database: %+v", dbPromocodes)

	// collect promocodes for particular shop to slice
	var shopPromocodes []Promocode
	for _, p := range dbPromocodes {
		if (shopId == nil || shopId != nil && p.ShopID == *shopId) &&
			(promoType == nil || promoType != nil && p.Type == *promoType) {
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

// Add - adds promocode to DB
func (db *DB) Add(promo *Promocode) error {
	// read Promocodes from DynamoDB
	dbPromocodes, err := db.read()
	if err != nil {
		log.Errorf("can not read from DynamodDB: %s", err.Error())
		return err
	}
	log.Debugf("Read from database: %+v", dbPromocodes)

	// supposed to be merge
	// remove particular shop promocodes
	var savePromocodes []Promocode
	for _, p := range dbPromocodes {
		if promo.ShopID != "" && promo.ShopID != p.ShopID {
			savePromocodes = append(savePromocodes, p)
		}
	}
	savePromocodes = append(savePromocodes, *promo)

	// save promocodes to DB
	if err := db.save(savePromocodes); err != nil {
		log.Errorf("can not save promocodes to DynamoDB: %s", err.Error())
		return err
	}

	return nil
}

// Delete - removes a promocode from DB
func (db *DB) Delete(promo *Promocode) error {
	// read Promocodes from DynamoDB
	dbPromocodes, err := db.read()
	if err != nil {
		log.Errorf("can not read from DynamodDB: %s", err.Error())
		return err
	}
	log.Debugf("Read from database: %+v", dbPromocodes)

	// remove promocode
	var savePromocodes []Promocode
	for _, p := range dbPromocodes {
		if !((promo.ShopID != "" && promo.ShopID == p.ShopID) &&
			(promo.Type == "" || promo.Type != "" && promo.Type == p.Type) &&
			(promo.Data == p.Data)) {
			savePromocodes = append(savePromocodes, p)
		}
	}

	// save promocodes to DB
	if err := db.save(savePromocodes); err != nil {
		log.Errorf("can not save promocodes to DynamoDB: %s", err.Error())
		return err
	}

	return nil
}

// List - lists all promocodes
func (db *DB) List() ([]Promocode, error) {
	// read Promocodes from DynamoDB
	dbPromocodes, err := db.read()
	if err != nil {
		log.Errorf("can not read from DynamodDB: %s", err.Error())
		return nil, err
	}
	log.Debugf("Read from database: %+v", dbPromocodes)

	return dbPromocodes, nil
}

// Read from DB and unmarshal
func (db *DB) read() ([]Promocode, error) {
	client := Client{DynamoDB: db.Client}

	// read Promocodes from DynamoDB
	jstr, err := client.DynamoDBGet(db.TableName, db.Prefix+PromoBotDBName)
	if err != nil {
		if err.Error() != ErrorDynamoDBIDNotFound { //get was unsuccessfull
			return nil, err
		}

		log.Debug("DynamoDB item not found")
		jstr = `[]`
	}

	// unmarshall DB content to Promocodes
	var promocodes []Promocode
	if err := json.Unmarshal([]byte(jstr), &promocodes); err != nil {
		log.Errorf("can not unmarshal promocodes DynamoDB string: %s error: %s", jstr, err.Error())
		return nil, err
	}

	// fmt.Printf("read: %+v\n", promocodes)
	return promocodes, nil
}

// Marshal and save promocodes to DB
func (db *DB) save(promocodes []Promocode) error {
	jbyte, err := json.Marshal(promocodes)
	if err != nil {
		return err
	}

	client := Client{DynamoDB: db.Client}
	err = client.DynamoDBPut(db.TableName, db.Prefix+PromoBotDBName, string(jbyte))

	// fmt.Printf("save: %+v\n", promocodes)
	return err
}
