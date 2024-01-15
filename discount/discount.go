package discount

import (
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

type Id string

type Output struct {
	Discounted DiscountedGroups
	Rest       item.ItemsQuantities
}

type DiscountedGroups []DiscountedGroup

func (me *DiscountedGroups) Append(d DiscountedGroup) {
	*me = append(*me, d)
}

type DiscountedGroup struct {
	Id
	Group item.ItemsQuantities
	Total money.Cents
}
