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
		itemCatalog: catalog,
		items:       make(map[item.Id]uint),
	}
}

type Basket struct {
	itemCatalog item.Catalog
	items       map[item.Id]uint
}

func (my Basket) ItemsCount() int {
	return len(my.items)
}

func (my *Basket) Add(itemId item.Id, qty uint) error {
	if _, found := my.itemCatalog[itemId]; !found {
		return fmt.Errorf("item not found in catalog: item.Id(%q)", itemId)
	}
	my.items[itemId] += qty
	return nil
}
func (my *Basket) MustAdd(n item.Id, qty uint) {
	err := my.Add(n, qty)
	if err != nil {
		log.Fatal(err)
	}
}
