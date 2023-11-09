package basket

import (
	"fmt"
	"log"

	"dev.acorello.it/go/supermarket-kata/item"
)

func NewBasket(catalog item.Catalog) Basket {
	if catalog == nil {
		panic("nil ItemCatalog")
	}
	return Basket{
		catalog: catalog,
		items:   make(map[item.Id]int),
	}
}

type Basket struct {
	catalog item.Catalog
	items   map[item.Id]int
}

func (my Basket) ItemsCount() int {
	return len(my.items)
}

func (my *Basket) Add(itemId item.Id, qty int) error {
	if _, found := my.catalog[itemId]; !found {
		return fmt.Errorf("item not found in catalog: item.Id(%q)", itemId)
	}
	if qty < 1 {
		return fmt.Errorf("quantity must be at least 1, got: %d", qty)
	}
	my.items[itemId] += qty
	return nil
}

func (my *Basket) MustAdd(n item.Id, qty int) {
	err := my.Add(n, qty)
	if err != nil {
		log.Fatal(err)
	}
}
