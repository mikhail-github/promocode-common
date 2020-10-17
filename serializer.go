package promocode

import (
	"encoding/json"
)

func Marshal(list []Promocode) ([]byte, error) {
	j, err := json.Marshal(list)

	return j, err
}

func Unmarshal(j []byte) ([]Promocode, error) {
	var list []Promocode
	err := json.Unmarshal(j, &list)

	return list, err
}
