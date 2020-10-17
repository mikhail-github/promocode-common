package promocode

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	adidasPromocode = `U20-KKLR-Z2GW-RTWG-XZZXM`
	reebokPromocode = `CMFN-RKMB-W3GW-9KD6`
)

func TestParse(t *testing.T) {
	t.Run("Parse Adidas promocode from string 1", func(t *testing.T) {
		str := fmt.Sprintf("%s", adidasPromocode)

		expected := []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode,
			},
		}
		promocode := Parse(str)
		if !reflect.DeepEqual(promocode, expected) {
			t.Logf("expected: %+v, got instead: %+v", expected, promocode)
			t.Fail()
		}
	})

	t.Run("Parse Adidas promocode from string 2", func(t *testing.T) {
		str := fmt.Sprintf("aaa %s bbb", adidasPromocode)

		expected := []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode,
			},
		}
		promocode := Parse(str)
		if !reflect.DeepEqual(promocode, expected) {
			t.Logf("expected: %+v, got instead: %+v", expected, promocode)
			t.Fail()
		}
	})

	t.Run("Parse double Adidas promocode from string", func(t *testing.T) {
		str := fmt.Sprintf("aaa %s bbb %s ccc", adidasPromocode, adidasPromocode)

		expected := []Promocode{
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode,
			},
			Promocode{
				ShopID: AdidasShopID,
				Type:   Adidas20,
				Data:   adidasPromocode,
			},
		}
		promocode := Parse(str)
		if !reflect.DeepEqual(promocode, expected) {
			t.Logf("expected: %+v, got instead: %+v", expected, promocode)
			t.Fail()
		}
	})

	t.Run("Parse Reebok promocode from string 1", func(t *testing.T) {
		str := fmt.Sprintf("%s", reebokPromocode)

		expected := []Promocode{
			Promocode{
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
		}
		promocode := Parse(str)
		if !reflect.DeepEqual(promocode, expected) {
			t.Logf("expected: %+v, got instead: %+v", expected, promocode)
			t.Fail()
		}
	})

	t.Run("Parse Reebok promocode from string 2", func(t *testing.T) {
		str := fmt.Sprintf("aaa %s bbb", reebokPromocode)

		expected := []Promocode{
			Promocode{
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
		}
		promocode := Parse(str)
		if !reflect.DeepEqual(promocode, expected) {
			t.Logf("expected: %+v, got instead: %+v", expected, promocode)
			t.Fail()
		}
	})

	t.Run("Parse double Reebok promocode from string", func(t *testing.T) {
		str := fmt.Sprintf("aaa %s bbb %s ccc", reebokPromocode, reebokPromocode)

		expected := []Promocode{
			Promocode{
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
			Promocode{
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
		}
		promocode := Parse(str)
		if !reflect.DeepEqual(promocode, expected) {
			t.Logf("expected: %+v, got instead: %+v", expected, promocode)
			t.Fail()
		}
	})

	t.Run("Parse multiple promocode from string", func(t *testing.T) {
		str := fmt.Sprintf("aaa %s bbb %s ccc %s ddd %s",
			reebokPromocode, reebokPromocode, adidasPromocode, reebokPromocode)

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
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
			Promocode{
				ShopID: ReebokShopID,
				Type:   Reebok20,
				Data:   reebokPromocode,
			},
		}
		promocode := Parse(str)
		if !reflect.DeepEqual(promocode, expected) {
			t.Logf("expected: %+v, got instead: %+v", expected, promocode)
			t.Fail()
		}
	})
}
