package discount

import "dev.acorello.it/go/supermarket-kata/item"

type noDiscounts struct {
}

func NoDiscounts() noDiscounts {
	return noDiscounts{}
}

func (noDiscounts) Discount(items []item.ItemQuantity) (res []DiscountedItems, fullPrice []item.ItemQuantity) {
	return nil, items
}
