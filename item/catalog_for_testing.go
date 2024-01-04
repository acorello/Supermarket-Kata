package item

import (
	"log"
)

type InMemoryCatalog map[Id]Item

var catalog = newFixedCatalog()

func FixedCatalog() InMemoryCatalog {
	return catalog
}

func newFixedCatalog() InMemoryCatalog {
	items_ := make(map[Id]Item, len(fixedItems))
	for _, item_ := range fixedItems {
		items_[item_.Id] = item_
	}
	return InMemoryCatalog(items_)
}

var fixedItems = []Item{
	{Id: "beans", Price: 90, Unit: "u"},
	{Id: "oranges", Price: 200, Unit: "kg"},
}

func (me InMemoryCatalog) RandomItems(count int) (result []Item) {
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

func (me InMemoryCatalog) Has(id Id) bool {
	_, found := me[id]
	return found
}

func (me InMemoryCatalog) Get(id Id) (Item, bool) {
	i, found := me[id]
	return i, found
}

func (me InMemoryCatalog) Len() int {
	return len(me)
}
