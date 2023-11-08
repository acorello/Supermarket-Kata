package basket_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/basket"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/test_fixtures"
	"github.com/stretchr/testify/assert"
)

var catalog item.Catalog = test_fixtures.Catalog()

func TestBasket_Add_then_ItemCount(t *testing.T) {
	t.Parallel()

	basket := basket.NewBasket(catalog)

	t.Run("empty basket initial state", func(t *testing.T) {
		assert.Equalf(t, 0, basket.ItemsCount(), "should have ItemCount of zero")
	})

	t.Run("error when adding item not in catalog", func(t *testing.T) {
		const id = "non-existent-item-id"
		if err := basket.Add(id, 1); err == nil {
			t.Errorf("Expected error when adding non existent element: item.Id(%s)", id)
		}
	})

	someItems := catalog.FetchRandomItems(2)
	item1 := someItems[0]
	item2 := someItems[1]
	t.Run("add first item", func(t *testing.T) {
		basket.MustAdd(item1.Id, 1)
		assert.Equalf(t, 1, basket.ItemsCount(),
			"adding an item to empty basket should result in ItemCount of one")
	})
	t.Run("add same item", func(t *testing.T) {
		basket.MustAdd(item1.Id, 1)
		assert.Equalf(t, 1, basket.ItemsCount(),
			"adding item already in basket should not increase ItemCount")
	})
	t.Run("add new item", func(t *testing.T) {
		basket.MustAdd(item2.Id, 1)
		assert.Equalf(t, 2, basket.ItemsCount(),
			"adding new item to non-empty basket should increase ItemCount")
	})
}
