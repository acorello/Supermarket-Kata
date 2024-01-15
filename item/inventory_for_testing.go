package item

import (
	"fmt"
	"log"

	"dev.acorello.it/go/supermarket-kata/money"
)

type InMemoryInventory map[Id]Item

var inventory = newFixedInventory()

func FixedInventory() InMemoryInventory {
	return inventory
}

func newFixedInventory() InMemoryInventory {
	prices := make(map[money.Cents]int, len(fixedItems))
	items_ := make(map[Id]Item, len(fixedItems))
	for _, item_ := range fixedItems {
		items_[item_.Id] = item_
		prices[item_.Price] += 1
	}
	for price, instances := range prices {
		if instances > 1 {
			panic(fmt.Sprintf("some tests rely on items returned by Random having distinct prices; found %v with price %v", instances, price))
		}
	}
	return items_
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
	if count != 0 {
		log.Fatalf("can't return %d items", count)
	}
	return result
}

func (me InMemoryInventory) Knows(id Id) bool {
	_, found := me[id]
	return found
}

func (me InMemoryInventory) Items(ids []Id) (res []Item) {
	for _, id := range ids {
		if itm, found := me[id]; found {
			res = append(res, itm)
		}
	}
	return res
}

func (me InMemoryInventory) Len() int {
	return len(me)
}
