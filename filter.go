package tgbotapi

import (
	"fmt"
)

type filter struct {
	Type  UpdateType
	Value string
}

func NewFilter(t UpdateType, v string) filter {
	switch t {
	case CommandUpdateType:
		v = fmt.Sprintf("/%s", v)
	}

	return filter{
		Type:  t,
		Value: v,
	}
}
