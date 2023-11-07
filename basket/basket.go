package basket

import (
	"dev.acorello.it/go/supermarket-kata/discount"
	"dev.acorello.it/go/supermarket-kata/item"
)

type Basket struct {
	discounters []discount.Discounter
	items       []item.Item
}
