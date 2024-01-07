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
	catalog Inventory
	items   map[item.Id]item.Quantity
}

type Inventory interface {
	Discounts([]item.ItemIdQuantity) item.Discounting
	Has(id item.Id) bool
}

func NewBasket(catalog Inventory) Basket {
	if catalog == nil {
		panic("nil catalog")
	}
	return Basket{
		Id:      Id(uuid.New()),
		catalog: catalog,
		items:   make(map[item.Id]item.Quantity),
	}
}

const MaxQuantity = 99

// Put increments the quantity of itemId by given amount.
//
// Returns error:
//   - if item is not in catalog
//   - if total quantity has exceeded [MaxQuantity]
func (my *Basket) Put(id item.Id, qty item.Quantity) error {
	if !my.catalog.Has(id) {
		return fmt.Errorf("item id %v not found in catalog", id)
	}
	if qty == 0 {
		return fmt.Errorf("can't put zero items")
	}
	basketQty := my.items[id]
	newQty := basketQty + qty
	if newQty > MaxQuantity {
		return fmt.Errorf("too many items; max allowed: %d", MaxQuantity)
	}
	my.items[id] = newQty
	return nil
}

// Remove decrements of given item by given quantity.
// If quantity of items in basket reaches zero the item is removed.
//
// Returns error:
//   - if item not found in basket
//   - if quantity removed is greater than quantity in basket
func (my *Basket) Remove(id item.Id, qty item.Quantity) error {
	basketQty, found := my.items[id]
	if !found {
		return fmt.Errorf("item id %v not in basket", id)
	}
	if qty > basketQty {
		return fmt.Errorf("can not remove %d items; basket contains only %d", qty, basketQty)
	}
	newQty := basketQty - qty
	if newQty == 0 {
		delete(my.items, id)
	} else {
		my.items[id] = newQty
	}
	return nil
}

func (my *Basket) Total() money.Cents {
	var list []item.ItemIdQuantity
	for id, qty := range my.items {
		list = append(list, item.ItemIdQuantity{Id: id, Quantity: qty})
	}
	discounting := my.catalog.Discounts(list)
	var total money.Cents
	for _, discounted := range discounting.DiscountedItems {
		total += discounted.Total()
	}
	for _, fullPrice := range discounting.FullPriceItems {
		total += fullPrice.Total()
	}
	return total
}
