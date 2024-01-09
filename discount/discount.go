package discount

import (
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

type DiscountId string

type DiscountedItems struct {
	DiscountId
	Group []item.ItemQuantity
	Total money.Cents
}
