# Promocode package

## Types
```
type Promocode struct {
	ShopID PromocodeShopID `json:"shop_id"`
	Type   PromocodeType   `json:"type,omitempty"`
	Data   string          `json:"data"`
}
```

## Functions
### .Parse(string) []Promocode  
parses promocodes from string