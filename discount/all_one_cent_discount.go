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

const IdAllOneCent = Id("all-one-cent")

func (allOneCent) Discount(items item.ItemsQuantities) (output Output) {
	for i, q := range items {
		var group item.ItemsQuantities
		group.Add(i, q)
		output.Discounted.Append(DiscountedGroup{
			Id:    IdAllOneCent,
			Group: group,
			Total: 1 * money.Cents(q),
		})
	}
	return output
}
