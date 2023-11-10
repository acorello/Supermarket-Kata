package basket

import (
	"fmt"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

type quantity struct {
	int
}

func Qty(v int) quantity {
	if q, e := Quantity(v); e != nil {
		panic(e)
	} else {
		return q
	}
}

func Quantity(v int) (quantity, error) {
	var z quantity
	if v <= 0 {
		return z, fmt.Errorf("Quantity <= 0: %v", v)
	}
	return quantity{v}, nil
}

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
func (my *Basket) Add(itemId item.Id, quantity quantity) (int, error) {
	if !my.catalogHas(itemId) {
		return 0, fmt.Errorf("item not found in catalog: %#v", itemId)
	}
	my.items[itemId] += quantity.int
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
