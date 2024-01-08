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

type Quantity uint
