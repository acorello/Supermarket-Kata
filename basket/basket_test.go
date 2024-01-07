package basket_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/basket"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const zeroCents = money.Cents(0)

var catalog = item.FixedCatalog()
var someItems = catalog.RandomItems(2)
var anItem, anotherItem = someItems[0], someItems[1]

// shorhand for sub-test function signatures
type T = *testing.T

func TestBasket(t *testing.T) {
	// a basket depends on a catalog
	// -> verify that no basket can be created with invalid dependencies
	require.Panics(t, func() {
		var nilCatalog basket.Catalog = nil
		basket.NewBasket(nilCatalog)
	}, "panic when given nil catalog")

	t.Run("Putting an item not in catalog returns an error", func(t T) {
		t.Parallel()
		b := basket.NewBasket(catalog)

		invalidItem := item.Id("item-not-in-catalog")
		err := b.Put(invalidItem, 1)
		require.Error(t, err, "basket seems to have accepted an item not in catalog")
	})

	// assuming anItem is immutable: sub-tests are reading it
	require.NotEqual(t, anItem, anotherItem, "catalog should return a different item")

	t.Run("Total changes as we change an item quantity", func(t T) {
		// TODO: rename b to basket…
		t.Parallel()
		b := basket.NewBasket(catalog)

		// ASSUMPTIONS
		require.Greater(t, anItem.Price, 1, "following tests rely on item.Price > 1")
		require.Equal(t, zeroCents, b.Total(), "total of empty basket should be zero")

		// TESTS
		require.NoError(t, b.Put(anItem.Id, 1))
		require.Equal(t, anItem.Price, b.Total())

		require.NoError(t, b.Put(anItem.Id, 2))
		require.Equal(t, anItem.Price*3, b.Total())

		require.NoError(t, b.Remove(anItem.Id, 1))
		require.Equal(t, anItem.Price*2, b.Total())

		require.NoError(t, b.Remove(anItem.Id, 2))
		require.Equal(t, zeroCents, b.Total())

		// removing a VALID ITEM NOT PRESENT in a NON-EMPTY BASKET returns error
		require.Error(t, b.Remove(anotherItem.Id, 2))
		require.Equal(t, zeroCents, b.Total())

		// BASKET NOW EMPTY, so should return an error
		require.Error(t, b.Remove(anItem.Id, 1))
		require.Equal(t, zeroCents, b.Total(),
			"total remains zero after removing items from empty basket")
	})

	t.Run("removing a VALID ITEM NOT PRESENT in a NON-EMPTY BASKET returns error",
		func(t T) {
			t.Parallel()
			b := basket.NewBasket(catalog)

			require.NoError(t, b.Put(anItem.Id, 1))

			require.Error(t, b.Remove(anotherItem.Id, 1))
			require.Equal(t, anItem.Price, b.Total())
		})

	t.Run("removing a VALID ITEM from an EMPTY BASKET returns error",
		func(t T) {
			t.Parallel()
			b := basket.NewBasket(catalog)

			// BASKET NOW EMPTY, so should return an error
			require.Error(t, b.Remove(anItem.Id, 1))
			require.Equal(t, zeroCents, b.Total(),
				"total remains zero after removing items from empty basket")
		})
}

func TestBasket_rejectsPuttingZeroItems(t *testing.T) {
	t.Parallel()
	b := basket.NewBasket(catalog)

	initTotal := b.Total()
	require.Error(t, b.Put(anItem.Id, 0))
	require.Equal(t, initTotal, b.Total())
}

func TestBasket_rejectsPuttingTooManyItems(t *testing.T) {
	t.Parallel()
	b := basket.NewBasket(catalog)

	require.NoError(t, b.Put(anItem.Id, 1))
	require.Equal(t, anItem.Price, b.Total())

	assert.Error(t, b.Put(anItem.Id, basket.MaxQuantity))
	assert.Equal(t, anItem.Price, b.Total(), "total modified on invalid operation")
}

func TestBasket_rejectsRemovingTooManyItems(t *testing.T) {
	t.Parallel()
	b := basket.NewBasket(catalog)

	require.NoError(t, b.Put(anItem.Id, 1))
	require.Equal(t, anItem.Price, b.Total())

	assert.Error(t, b.Remove(anItem.Id, 2))
	assert.Equal(t, anItem.Price, b.Total(), "total modified on invalid operation")
}
