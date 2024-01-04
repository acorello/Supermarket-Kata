package basket

import (
	"fmt"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
	"github.com/google/uuid"
)

type Id uuid.UUID

type Basket struct {
	Id
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
		Id:      Id(uuid.New()),
		catalog: catalog,
		items:   make(map[item.Id]int),
	}
}

// Put increments the quantity of itemId by given amount.
//
// Returns error:
//   - if item is not in catalog
//   - TODO: if total quantity has exceeded reasonable maximum
func (my *Basket) Put(id item.Id, qty quantity) error {
	if !my.catalog.Has(id) {
		return fmt.Errorf("item id %v not found in catalog", id)
	}
	my.items[id] += qty.int
	return nil
}

// Remove decrements of given item by given quantity.
// If quantity of items in basket reaches zero the item is removed.
//
// Returns error:
//   - if item not found in basket
//   - TODO: if quantity removed is greater than quantity in basket
func (my *Basket) Remove(id item.Id, qty quantity) error {
	q, found := my.items[id]
	if !found {
		return fmt.Errorf("item id %v not in basket", id)
	}
	r := q - qty.int
	if r < 1 {
		delete(my.items, id)
	} else {
		my.items[id] = r
	}
	return nil
}

func (my *Basket) Total() money.Cents {
	var total money.Cents
	for id, qty := range my.items {
		i, _ := my.catalog.Get(id)
		total += money.Cents(int(i.Price) * qty)
	}
	return total
}

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
