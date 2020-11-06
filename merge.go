package common

func Merge(promocodes, new []Promocode) []Promocode {
	for _, newP := range new {
		if !isPromocodeInSlice(newP, promocodes) {
			promocodes = append(promocodes, newP)
		}
	}

	return promocodes
}

func isPromocodeInSlice(p Promocode, s []Promocode) bool {
	for _, v := range s {
		if p == v {
			return true
		}
	}

	return false
}
