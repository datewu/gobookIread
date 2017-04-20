package proto

import (
	"errors"
	"fmt"
)

// ShirtCloner the prototype interface
type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

// White clour skus
const (
	White = 1
	Black = 2
	Blue  = 3
)

// GetShirtCloner retrieve a new cloner
func GetShirtCloner() ShirtCloner {
	return new(ShirtCache)
}

// ShirtCache implement ShirtCloner interface
type ShirtCache struct{}

// GetClone statisfy ShirtCloner interface
func (*ShirtCache) GetClone(m int) (ItemInfoGetter, error) {
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
	}
	return nil, errors.New("Not implemented yet")
}

// ItemInfoGetter ll
type ItemInfoGetter interface {
	GetInfo() string
}

// ShirtColor as the name suggest
type ShirtColor byte

// Shirt struct the big star
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// GetInfo statisfy ItemInfoGetter interface
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

var whitePrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

// GetPrice normal method
func (s *Shirt) GetPrice() float32 {
	return s.Price
}
