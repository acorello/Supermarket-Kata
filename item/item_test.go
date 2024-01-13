package item_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/item"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var inventory = item.FixedInventory()

func TestFixedCatalog_RandomItems(t *testing.T) {
	wanted := 1
	require.GreaterOrEqual(t, inventory.Len(), wanted)

	items := inventory.RandomItems(wanted)
	got := len(items)

	assert.Equal(t, wanted, got)
}

func TestItemsQuantities_Add(t *testing.T) {
	t.Run("nil accepts new elements", func(t *testing.T) {
		var itemsQtys item.ItemsQuantities
		require.Nil(t, itemsQtys)
		itemsQtys.Add(item.Item{}, 1)
		require.NotNil(t, itemsQtys)
		require.Equal(t, itemsQtys[item.Item{}], item.Quantity(1))
	})
}
