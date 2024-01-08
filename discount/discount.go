package discount

import (
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

type DiscountId string

type DiscountedItems struct {
	DiscountId
	Items []item.ItemQuantity
	total money.Cents
}

func (me DiscountedItems) Total() money.Cents {
	return me.total
}
