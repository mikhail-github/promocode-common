package common

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Run("Promocodes merge existing promo", func(t *testing.T) {
		promocodes := []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode,
			},
			Promocode{
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
		}

		new := []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode,
			},
		}

		expected := []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode,
			},
			Promocode{
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
		}

		res := Merge(promocodes, new)
		if !reflect.DeepEqual(res, expected) {
			t.Logf("expected: %+v, got instead: %+v", expected, res)
			t.Fail()
		}
	})

	t.Run("Promocodes merge new promo", func(t *testing.T) {
		promocodes := []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode,
			},
			Promocode{
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
		}

		new := []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode2,
			},
		}

		expected := []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode,
			},
			Promocode{
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode2,
			},
		}

		res := Merge(promocodes, new)
		if !reflect.DeepEqual(res, expected) {
			t.Logf("expected: %+v, got instead: %+v", expected, res)
			t.Fail()
		}
	})
}
