package item

import (
	"log"
)

type InMemoryInventory map[Id]Item

var catalog = newFixedInventory()

func FixedInventory() InMemoryInventory {
	return catalog
}

func newFixedInventory() InMemoryInventory {
	items_ := make(map[Id]Item, len(fixedItems))
	for _, item_ := range fixedItems {
		items_[item_.Id] = item_
	}
	return InMemoryInventory(items_)
}

var fixedItems = []Item{
	{Id: "beans", Price: 90, Unit: "u"},
	{Id: "oranges", Price: 200, Unit: "kg"},
}

func (me InMemoryInventory) RandomItems(count int) (result []Item) {
	if count < 0 {
		log.Fatalf("can't return %d items", count)
	}
	if count > me.Len() {
		log.Fatalf("not enough elements in catalog. Wanted %d, got %d", count, me.Len())
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

func (me InMemoryInventory) Has(id Id) bool {
	_, found := me[id]
	return found
}

func (me InMemoryInventory) PricedItems(qtyIds []ItemIdQuantity) PricedItems {
	var its []ItemQuantity
	for _, qi := range qtyIds {
		it, found := me[qi.Id]
		if !found {
			log.Panicf("item %v not found; expected to be called from basket always with  valid ids", qi)
		}
		its = append(its, ItemQuantity{Item: it, Quantity: qi.Quantity})
	}
	return PricedItems{Discounted: nil, FullPrice: its}
}

func (me InMemoryInventory) Len() int {
	return len(me)
}
