package discount

import "dev.acorello.it/go/supermarket-kata/item"

type noDiscounts struct {
}

func NoDiscounts() noDiscounts {
	return noDiscounts{}
}

func (noDiscounts) Discount(items item.ItemsQuantities) (output Output) {
	output.Rest = items
	return
}
