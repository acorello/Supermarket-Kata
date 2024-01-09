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

func (allOneCent) Discount(items []item.ItemQuantity) (res []DiscountedItems, fullPrice []item.ItemQuantity) {
	const discountId = DiscountId("all-one-cent")
	for _, i := range items {
		res = append(res, DiscountedItems{
			DiscountId: discountId,
			Group:      []item.ItemQuantity{i},
			Total:      1 * money.Cents(i.Quantity),
		})
	}
	return res, nil
}
