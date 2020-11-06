package common

import (
	"regexp"
)

type StringToPromocode struct {
	Regexp string          `json:"regexp"`
	ShopID PromocodeShopID `json:"shop_id"`
	Type   PromocodeType   `json:"type"`
}

// Parse - parses promocodes from string
func Parse(s string) []Promocode {
	var promocodes []Promocode
	for _, v := range stringToPromocodes {
		re := regexp.MustCompile(v.Regexp)
		match := re.FindAllStringSubmatch(string(s), -1)
		for j := range match {
			for i, name := range re.SubexpNames() {
				if name == promocodeNamedRegexp {
					promocode := Promocode{
						ShopID: v.ShopID,
						Type:   v.Type,
						Data:   match[j][i],
					}
					promocodes = append(promocodes, promocode)
				}
			}
		}
	}

	return promocodes
}
