package item

import (
	"log"
)

type InMemoryInventory map[Id]Item

var inventory = newFixedInventory()

func FixedInventory() InMemoryInventory {
	return inventory
}

func newFixedInventory() InMemoryInventory {
	items_ := make(map[Id]Item, len(fixedItems))
	for _, item_ := range fixedItems {
		items_[item_.Id] = item_
	}
	return InMemoryInventory(items_)
}

func (me InMemoryInventory) RandomItems(count int) (result []Item) {
	if count < 0 {
		log.Fatalf("can't return %d items", count)
	}
	if count > me.Len() {
		log.Fatalf("not enough elements in inventory. Wanted %d, got %d", count, me.Len())
	}
	for _, item := range me {
		result = append(result, item)
		count--
		if count <= 0 {
			break
		}
	}
	return result
}

func (me InMemoryInventory) Knows(id Id) bool {
	_, found := me[id]
	return found
}

func (me InMemoryInventory) PricedItems(qtyIds []ItemIdQuantity) []ItemQuantity {
	var items []ItemQuantity
	for _, qi := range qtyIds {
		anItem, found := me[qi.Id]
		if !found {
			log.Panicf("item %v not found; client should invoke [Knows] to validate item.Id", qi)
		}
		items = append(items, ItemQuantity{Item: anItem, Quantity: qi.Quantity})
	}
	return items
}

func (me InMemoryInventory) Len() int {
	return len(me)
}
