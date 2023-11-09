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

type Catalog map[Id]Item

func (me Catalog) FetchRandomItems(count int) (result []Item) {
	if count < 0 {
		log.Fatalf("can't return %d items", count)
	}
	if count > me.len() {
		log.Fatalf("not enough elements in catalog. Wanted %d, got %d", count, me.len())
	}
	for _, item := range me {
		result = append(result, item)
	}
	return result
}

func (me Catalog) len() int {
	return len(me)
}

func NewCatalog(items ...Item) Catalog {
	catalog := make(map[Id]Item, len(items))
	for _, i := range items {
		catalog[i.Id] = i
	}
	return catalog
}
