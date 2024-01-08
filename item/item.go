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
	Items    []ItemQuantity
	total    money.Cents
}

func (me DiscountedItems) Total() money.Cents {
	return me.total
}

type PricedItems struct {
	Discounted []DiscountedItems
	FullPrice  []ItemQuantity
}

type Quantity uint
