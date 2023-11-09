package test_fixtures

import (
	_ "embed"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

// this path is
var catalog = makeCatalog()

func Catalog() item.Catalog {
	return catalog
}

func makeCatalog() item.Catalog {
	items_ := make([]item.Item, len(catalog_))
	for i, item_ := range catalog_ {
		items_[i] = item.Item{
			Id:    item_.Id,
			Price: item_.Cents,
			Unit:  item_.Unit,
		}
	}
	return item.NewCatalog(items_...)
}

var catalog_ = []struct {
	item.Id
	Unit string
	money.Cents
}{
	{"beans", "u", 90},
	{"oranges", "kg", 200},
}
