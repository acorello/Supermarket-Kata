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

type Basket struct {
	catalog Catalog
	items   map[item.Id]int
}

type Catalog interface {
	Get(id item.Id) (item.Item, bool)
	Has(id item.Id) bool
}

func NewBasket(catalog Catalog) Basket {
	if catalog == nil {
		panic("nil catalog")
	}
	return Basket{
		catalog: catalog,
		items:   make(map[item.Id]int),
	}
}

// Put increments the quantity of itemId by given amount; returns updated quantity.
func (my *Basket) Put(id itemId, qty quantity) int {
	my.items[id.value] += qty.int
	return my.items[id.value]
}

// Remove decrements of given item by given quantity.
// If quantity of items in basket reaches zero the item is removed.
// Removing more items than present in basket is equivalent to removing all of them.
//
// Returns updated quantity.
func (my *Basket) Remove(id itemId, qty quantity) int {
	q := my.items[id.value]
	r := q - qty.int
	if r < 0 {
		delete(my.items, id.value)
		return 0
	} else {
		my.items[id.value] = r
		return r
	}
}

func (my *Basket) Total() money.Cents {
	var total money.Cents
	for id, qty := range my.items {
		i, _ := my.catalog.Get(id)
		total += i.Price.Mul(qty)
	}
	return total
}

type itemId struct {
	value item.Id
}

// [ItemIdInCatalog] returns an [itemId] if given [item.Id] is present in Basket's [item.Catalog], error otherwise.
func (my *Basket) ItemIdInCatalog(id item.Id) (itemId, error) {
	if !my.catalog.Has(id) {
		return itemId{}, fmt.Errorf("item not in catalog %q", id)
	}
	return itemId{id}, nil
}
