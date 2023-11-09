package basket

import (
	"fmt"

	"dev.acorello.it/go/supermarket-kata/item"
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
// error != nil
//   - if itemId not in catalog
//   - if qty < 1
func (my *Basket) Add(itemId item.Id, qty int) (int, error) {
	if !my.catalogHas(itemId) {
		return 0, fmt.Errorf("item not found in catalog: item.Id(%q)", itemId)
	}
	if qty < 1 {
		return 0, fmt.Errorf("quantity must be at least 1, got: %d", qty)
	}
	return my.add(itemId, qty), nil
}

func (my *Basket) add(itemId item.Id, qty int) int {
	my.items[itemId] += qty
	return my.items[itemId]
}

func (my *Basket) catalogHas(itemId item.Id) bool {
	_, found := my.catalog[itemId]
	return found
}
