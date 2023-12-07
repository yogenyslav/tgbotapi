package tgbotapi

import (
	"fmt"
	"reflect"
)

type Filter struct {
	Type  UpdateType
	Value interface{}
}

func NewFilter(t UpdateType, v interface{}) Filter {
	switch t {
	case CommandUpdateType:
		v, ok := v.(string)
		if !ok {
			panic(InvalidType{reflect.TypeOf(v).Name(), reflect.TypeOf("").Name()})
		}

		v = fmt.Sprintf("/%s", v)
	}

	return Filter{
		Type:  t,
		Value: v,
	}
}
