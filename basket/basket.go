package basket

import (
	"fmt"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

type Basket struct {
	catalog item.Catalog
	items   map[item.Id]int
}

func NewBasket(catalog item.Catalog) Basket {
	if catalog == nil {
		panic("nil ItemCatalog")
	}
	return Basket{
		catalog: catalog,
		items:   make(map[item.Id]int),
	}
}

// Add increments the quantity of itemId by `qty` amount; returns updated quantity.
//
// error if itemId not in catalog
//
// error if quantity <= 0
func (my *Basket) Add(itemId item.Id, quantity int) (int, error) {
	if !my.catalogHas(itemId) {
		return 0, fmt.Errorf("item not found in catalog: %#v", itemId)
	}
	if quantity <= 0 {
		return 0, fmt.Errorf("quantity <= 0: %d", quantity)
	}
	my.items[itemId] += quantity
	return my.items[itemId], nil
}

func (my *Basket) Total() money.Cents {
	var total money.Cents
	for id, qty := range my.items {
		i := my.catalog[id]
		total += i.Price.Mul(qty)
	}
	return total
}

func (my *Basket) catalogHas(itemId item.Id) bool {
	_, found := my.catalog[itemId]
	return found
}
