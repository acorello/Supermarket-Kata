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

// Add increments the quantity of itemId by given amount; returns updated quantity.
func (my *Basket) Add(itemId knownItem, qty quantity) int {
	my.items[itemId.Id] += qty.int
	return my.items[itemId.Id]
}

// Remove decrements the quantity of itemId by given amount.
// If given amount is greater or equal to the amount in basket the item is comletely removed.
//
// Returns updated quantity.
func (my *Basket) Remove(itemId knownItem, qty quantity) int {
	q := my.items[itemId.Id]
	r := q - qty.int
	if r < 0 {
		delete(my.items, itemId.Id)
		return 0
	} else {
		my.items[itemId.Id] = r
		return r
	}
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
