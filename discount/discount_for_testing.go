package discount

import "dev.acorello.it/go/supermarket-kata/item"

func NoDiscounts() noDiscounts {
	return noDiscounts{}
}

type noDiscounts struct {
}

func (noDiscounts) Discount(items ...item.ItemQuantity) []DiscountedItems {
	return nil
}
