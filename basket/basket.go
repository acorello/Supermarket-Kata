package basket

import (
	"fmt"

	"dev.acorello.it/go/supermarket-kata/item"
)

type Basket struct {
	catalog item.Catalog
	items   map[item.Id]Quantity
}

// must be >= 0
type Quantity uint

func QuantityOfInt(q int) (Quantity, error) {
	if q <= 0 {
		return 0, fmt.Errorf("quantity <= 0: %d", q)
	}
	return Quantity(uint(q)), nil
}

func NewBasket(catalog item.Catalog) Basket {
	if catalog == nil {
		panic("nil ItemCatalog")
	}
	return Basket{
		catalog: catalog,
		items:   make(map[item.Id]Quantity),
	}
}

// Add increments the quantity of itemId by `qty` amount; returns updated quantity.
//
// error != nil if itemId not in catalog
func (my *Basket) Add(itemId item.Id, qty Quantity) (Quantity, error) {
	if !my.catalogHas(itemId) {
		return 0, fmt.Errorf("item not found in catalog: item.Id(%q)", itemId)
	}
	return my.add(itemId, qty), nil
}

func (my *Basket) add(itemId item.Id, qty Quantity) Quantity {
	my.items[itemId] += qty
	return my.items[itemId]
}

func (my *Basket) catalogHas(itemId item.Id) bool {
	_, found := my.catalog[itemId]
	return found
}
