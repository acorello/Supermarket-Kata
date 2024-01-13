package discount

import (
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

type DiscountId string

type Output struct {
	Discounted DiscountedGroups
	Rest       item.ItemsQuantities
}

type DiscountedGroups []DiscountedItems

func (me *DiscountedGroups) Append(d DiscountedItems) {
	*me = append(*me, d)
}

type DiscountedItems struct {
	DiscountId
	Group item.ItemsQuantities
	Total money.Cents
}
