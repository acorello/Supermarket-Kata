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

type ItemsQuantities map[Item]Quantity

func (m *ItemsQuantities) Add(i Item, q Quantity) {
	if *m == nil {
		*m = make(ItemsQuantities)
	}
	(*m)[i] += q
}

type ItemQuantity struct {
	Item // TODO: embedded Item exposes Price; may look like the price of whole
	Quantity
}

func (me ItemQuantity) Total() money.Cents {
	return me.Price * money.Cents(me.Quantity)
}

type Quantity uint
