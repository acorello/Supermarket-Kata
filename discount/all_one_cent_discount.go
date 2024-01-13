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

func (allOneCent) Discount(items item.ItemsQuantities) (output Output) {
	const discountId = DiscountId("all-one-cent")
	for i, q := range items {
		var group item.ItemsQuantities
		group.Add(i, q)
		output.Discounted.Append(DiscountedItems{
			DiscountId: discountId,
			Group:      group,
			Total:      1 * money.Cents(q),
		})
	}
	return output
}
