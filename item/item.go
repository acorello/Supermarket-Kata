package item

import (
	"dev.acorello.it/go/supermarket-kata/money"
)

type Id string

type Item struct {
	Id
	Price money.Cents
	Unit  string
}

type ItemIdQuantity struct {
	Id
	Quantity
}

type ItemQuantity struct {
	Item
	Quantity
}

func (me ItemQuantity) Total() money.Cents {
	return me.Price * money.Cents(me.Quantity)
}

type DiscountedItems struct {
	Discount string
	Items    []Item
}

func (me DiscountedItems) Total() money.Cents {
	panic("not implemented")
}

type Discounting struct {
	DiscountedItems []DiscountedItems
	FullPriceItems  []ItemQuantity
}

type Quantity uint
