package basket

import (
	"fmt"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

type quantity struct {
	int
}

func Quantity(v int) (*quantity, error) {
	const min, max = 1, 99
	if v < min || v > max {
		return nil, fmt.Errorf("Quantity q<=%d or q>%d. q==%v", v, min, max)
	}
	return &quantity{v}, nil
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
func (my *Basket) Put(id itemId, qty quantity) {
	my.items[id.value] += qty.int
}

// Remove decrements of given item by given quantity.
// If quantity of items in basket reaches zero the item is removed.
// Removing more items than present in basket is equivalent to removing all of them.
//
// Returns updated quantity.
func (my *Basket) Remove(id itemId, qty quantity) {
	q := my.items[id.value]
	r := q - qty.int
	if r < 1 {
		delete(my.items, id.value)
	} else {
		my.items[id.value] = r
	}
}

func (my *Basket) Total() money.Cents {
	var total money.Cents
	for id, qty := range my.items {
		i, _ := my.catalog.Get(id)
		total += money.Cents(int(i.Price) * qty)
	}
	return total
}

type itemId struct {
	value item.Id
}

// [ItemIdInCatalog] returns an [itemId] if given [item.Id] is present in Basket's [item.Catalog], error otherwise.
func (my *Basket) ItemIdInCatalog(id item.Id) (*itemId, error) {
	if !my.catalog.Has(id) {
		return nil, fmt.Errorf("item not in catalog %q", id)
	}
	return &itemId{id}, nil
}
