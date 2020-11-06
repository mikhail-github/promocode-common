package common

import (
	"math/rand"
	"time"
)

type PromocodeShopID string
type PromocodeType string

type Promocode struct {
	ShopID PromocodeShopID `json:"shop_id"`
	Type   PromocodeType   `json:"type,omitempty"`
	Data   string          `json:"data"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
