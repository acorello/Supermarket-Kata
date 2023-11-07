package discount

import (
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

type Discounter interface {
	// Discounter.Discounts looks for individual or single items to which it applies. If the discount did not apply it returns an empty list.
	//
	// Each of the returned discounts contains a cluster of items to which the discount applied, together with the amount to be discounted.
	//
	// If the list is not empty then each item will have a non-empty `[]Item`.
	// The discounted amount will be non-zero and less than the total
	//
	// # POST-CONDITION
	//		0 < sum(result..Amount) <= sum(items..Price)
	Discounts(items ...item.Item) []Discount
}

type Discount struct {
	Matches []item.Item
	Amount  money.Cents
}
