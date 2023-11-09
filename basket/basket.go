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
type Quantity struct {
	uint
}

func (q Quantity) Add(o Quantity) Quantity {
	return Quantity{q.uint + o.uint}
}

func QuantityOf(q int) (Quantity, error) {
	if q <= 0 {
		return Quantity{0}, fmt.Errorf("quantity <= 0: %d", q)
	}
	return Quantity{uint(q)}, nil
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
		return Quantity{0}, fmt.Errorf("item not found in catalog: item.Id(%q)", itemId)
	}
	q := my.items[itemId].Add(qty)
	my.items[itemId] = q
	return q, nil
}

func (my *Basket) catalogHas(itemId item.Id) bool {
	_, found := my.catalog[itemId]
	return found
}
