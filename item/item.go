package item

import (
	"log"

	"dev.acorello.it/go/supermarket-kata/money"
)

type Id string

type Item struct {
	Id
	Price money.Cents
	Unit  string
}

type Catalog struct {
	items map[Id]Item
}

func (me Catalog) RandomItem() Item {
	return me.RandomItems(1)[0]
}

func (me Catalog) RandomItems(count int) (result []Item) {
	if count < 0 {
		log.Fatalf("can't return %d items", count)
	}
	if count > me.Len() {
		log.Fatalf("not enough elements in catalog. Wanted %d, got %d", count, me.Len())
	}
	for _, item := range me.items {
		result = append(result, item)
		count--
		if count <= 0 {
			break
		}
	}
	return result
}

func (me Catalog) Has(id Id) bool {
	_, found := me.items[id]
	return found
}

func (me Catalog) Get(id Id) (Item, bool) {
	i, found := me.items[id]
	return i, found
}

func (me Catalog) Len() int {
	return len(me.items)
}

func NewCatalog(items ...Item) Catalog {
	catalog := make(map[Id]Item, len(items))
	for _, i := range items {
		catalog[i.Id] = i
	}
	return Catalog{catalog}
}
