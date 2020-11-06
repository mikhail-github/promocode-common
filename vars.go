package common

const (
	PromoBotDBName                 = "PromocodeBot"
	AdidasShopID   PromocodeShopID = "adidas"
	Adidas20       PromocodeType   = "20%"

	ReebokShopID PromocodeShopID = "reebok"
	Reebok20     PromocodeType   = "20%"

	promocodeNamedRegexp    = `promocode`
	adidasPromocodeRegexp20 = `(^|\s+)(?P<promocode>U20-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{5})($|\s+)`
	reebokPromocodeRegexp20 = `(^|\s+)(?P<promocode>[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4})($|\s+)`

	ErrorDynamoDBIDNotFound = "id not found"
	ErrorPromocodeNotFound  = "promocode not found"
)

var (
	shopIDList = []PromocodeShopID{AdidasShopID, ReebokShopID}

	stringToPromocodes = []StringToPromocode{
		StringToPromocode{Regexp: adidasPromocodeRegexp20, ShopID: AdidasShopID, Type: Adidas20},
		StringToPromocode{Regexp: reebokPromocodeRegexp20, ShopID: ReebokShopID, Type: Reebok20},
	}
)
