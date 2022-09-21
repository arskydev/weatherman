package weather

import (
	"encoding/json"
	"errors"
)

type weatherObject struct {
	Main string `json:"main"`
}

type weatherList []weatherObject

func (w *weatherList) UnmarshalJSON(d []byte) error {
	switch d[0] {
	case '{':
		var v weatherObject
		err := json.Unmarshal(d, &v)
		*w = weatherList{v}
		return err
	case '[':
		var v []weatherObject
		err := json.Unmarshal(d, &v)
		*w = weatherList(v)
		return err
	default:
		return errors.New("unexpected case")
	}
}
