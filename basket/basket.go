package basket

import (
	"fmt"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
	"github.com/google/uuid"
)

// TODO: extract and expose known domain errors (e.g. itemId not in catalog), to help clients handle errors.

type Id uuid.UUID

type Basket struct {
	Id
	catalog Catalog
	items   map[item.Id]quantity
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
		items:   make(map[item.Id]quantity),
	}
}

// Put increments the quantity of itemId by given amount.
//
// Returns error:
//   - if item is not in catalog
//   - if total quantity has exceeded [MaxQuantity]
func (my *Basket) Put(id item.Id, qty quantity) error {
	if !my.catalog.Has(id) {
		return fmt.Errorf("item id %v not found in catalog", id)
	}
	basketQty := my.items[id]
	if newQty, err := Quantity(basketQty.int + qty.int); err != nil {
		return err
	} else {
		my.items[id] = *newQty
		return nil
	}
}

// Remove decrements of given item by given quantity.
// If quantity of items in basket reaches zero the item is removed.
//
// Returns error:
//   - if item not found in basket
//   - if quantity removed is greater than quantity in basket
func (my *Basket) Remove(id item.Id, qtyToRemove quantity) error {
	basketQty, found := my.items[id]
	if !found {
		return fmt.Errorf("item id %v not in basket", id)
	}
	if qtyToRemove == basketQty {
		delete(my.items, id)
		return nil
	}
	if newQty, err := Quantity(basketQty.int - qtyToRemove.int); err != nil {
		return fmt.Errorf("can not remove requested quantity %d: %v", qtyToRemove.int, err)
	} else {
		my.items[id] = *newQty
	}
	return nil
}

func (my *Basket) Total() money.Cents {
	var total money.Cents
	for id, qty := range my.items {
		i, _ := my.catalog.Get(id)
		total += money.Cents(int(i.Price) * qty.int)
	}
	return total
}

const MinQuantity, MaxQuantity = 1, 99

type quantity struct {
	int
}

func Quantity(v int) (*quantity, error) {
	if v < MinQuantity {
		return nil, fmt.Errorf("Quantity %d smaller than minimum: %d", v, MinQuantity)
	}
	if v > MaxQuantity {
		return nil, fmt.Errorf("Quantity %d greater than maximum: %d", v, MaxQuantity)
	}
	return &quantity{v}, nil
}
