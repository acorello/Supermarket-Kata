package basket

import (
	"fmt"

	"dev.acorello.it/go/supermarket-kata/discount"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
	"github.com/google/uuid"
)

// TODO: extract and expose known domain errors (e.g. itemId not in inventory), to help clients handle errors.

type Id uuid.UUID

type Basket struct {
	Id
	inventory Inventory
	discounts Discounts
	items     map[item.Id]item.Quantity
}

type Inventory interface {
	PricedItems([]item.ItemIdQuantity) []item.ItemQuantity
	Knows(id item.Id) bool
}

type Discounts interface {
	Discount(items []item.ItemQuantity) ([]discount.DiscountedItems, []item.ItemQuantity)
}

func NewBasket(inventory Inventory, discounts Discounts) Basket {
	if inventory == nil {
		panic("nil parameter: inventory")
	}
	if discounts == nil {
		panic("nil parameter: discounts")
	}
	return Basket{
		Id:        Id(uuid.New()),
		inventory: inventory,
		discounts: discounts,
		items:     make(map[item.Id]item.Quantity),
	}
}

const MaxQuantity = 99

// Put increments the quantity of itemId by given amount.
//
// Returns error:
//   - if inventory doesn't know item.Id (never had such item, ever)
//   - if total quantity has exceeded [MaxQuantity]
func (my *Basket) Put(id item.Id, qty item.Quantity) error {
	if !my.inventory.Knows(id) {
		return fmt.Errorf("unknown item.Id %v", id)
	}
	if qty == 0 {
		return fmt.Errorf("can't put zero items")
	}
	basketQty := my.items[id]
	if newQty := basketQty + qty; newQty > MaxQuantity {
		return fmt.Errorf("too many items; had %d, plus %d = %d; but max allowed: %d", basketQty, qty, newQty, MaxQuantity)
	} else {
		my.items[id] = newQty
		return nil
	}
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

func (my *Basket) Total() (total money.Cents) {
	var list []item.ItemIdQuantity
	for id, qty := range my.items {
		list = append(list, item.ItemIdQuantity{Id: id, Quantity: qty})
	}
	items := my.inventory.PricedItems(list)
	discounted, fullPrice := my.discounts.Discount(items)
	for i := range fullPrice {
		total += fullPrice[i].Total()
	}
	for i := range discounted {
		total += discounted[i].Total
	}
	return total
}
