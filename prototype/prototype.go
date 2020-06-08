package prototype

import (
	"errors"
	"fmt"
)

// 原型模式
type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	showStr := fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %0.2f\n", s.SKU, s.Color, s.Price)
	return showStr
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

var whitePrototype *Shirt = &Shirt{
	Price: 15.0,
	SKU:   "0001",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.0,
	SKU:   "0002",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.0,
	SKU:   "0003",
	Color: Blue,
}

type ShirtCache struct{}

func (s *ShirtCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New(fmt.Sprintf("Invalide prototype: %d", m))
	}
}
