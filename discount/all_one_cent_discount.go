package discount

import (
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

func AllOneCentDiscount() allOneCent {
	return allOneCent{}
}

type allOneCent struct {
}

func (allOneCent) Discount(items []item.ItemQuantity) (
	discounted []DiscountedItems,
	rest []item.ItemQuantity,
) {
	const discountId = DiscountId("all-one-cent")
	for _, i := range items {
		var group item.ItemsQuantities
		group.Add(i.Item, i.Quantity)
		discounted = append(discounted, DiscountedItems{
			DiscountId: discountId,
			Group:      group,
			Total:      1 * money.Cents(i.Quantity),
		})
	}
	return discounted, nil
}
