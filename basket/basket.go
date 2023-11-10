package basket

import (
	"fmt"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

type quantity struct {
	int
}

func Quantity(v int) (quantity, error) {
	var z quantity
	if v <= 0 {
		return z, fmt.Errorf("Quantity <= 0: %v", v)
	}
	return quantity{v}, nil
}

type knownItem struct {
	item.Id
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
func (my *Basket) Add(itemId knownItem, quantity quantity) int {
	my.items[itemId.Id] += quantity.int
	return my.items[itemId.Id]
}

func (my *Basket) Total() money.Cents {
	var total money.Cents
	for id, qty := range my.items {
		i := my.catalog[id]
		total += i.Price.Mul(qty)
	}
	return total
}

func (my *Basket) KnownItemId(id item.Id) (knownItem, error) {
	if !my.catalogHas(id) {
		return knownItem{}, fmt.Errorf("item not in catalog %q", id)
	}
	return knownItem{id}, nil
}

func (my *Basket) catalogHas(itemId item.Id) bool {
	_, found := my.catalog[itemId]
	return found
}
