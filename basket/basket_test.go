package basket_test

import (
	"math"
	"testing"

	"dev.acorello.it/go/supermarket-kata/basket"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
	"dev.acorello.it/go/supermarket-kata/must"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const zeroCents = money.Cents(0)

var catalog = item.FixedCatalog()
var someItems = catalog.RandomItems(2)
var anItem, anotherItem = someItems[0], someItems[1]

var qty = must.DereFn(basket.Quantity)

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
		err := b.Put(invalidItem, qty(1))
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
		require.NoError(t, b.Put(anItem.Id, qty(1)))
		require.Equal(t, anItem.Price, b.Total())

		require.NoError(t, b.Put(anItem.Id, qty(2)))
		require.Equal(t, anItem.Price*3, b.Total())

		require.NoError(t, b.Remove(anItem.Id, qty(1)))
		require.Equal(t, anItem.Price*2, b.Total())

		require.NoError(t, b.Remove(anItem.Id, qty(2)))
		require.Equal(t, zeroCents, b.Total())

		// removing a VALID ITEM NOT PRESENT in a NON-EMPTY BASKET returns error
		require.Error(t, b.Remove(anotherItem.Id, qty(2)))
		require.Equal(t, zeroCents, b.Total())

		// BASKET NOW EMPTY, so should return an error
		require.Error(t, b.Remove(anItem.Id, qty(1)))
		require.Equal(t, zeroCents, b.Total(),
			"total remains zero after removing items from empty basket")
	})

	t.Run("removing a VALID ITEM NOT PRESENT in a NON-EMPTY BASKET returns error",
		func(t T) {
			t.Parallel()
			b := basket.NewBasket(catalog)

			require.NoError(t, b.Put(anItem.Id, qty(1)))

			require.Error(t, b.Remove(anotherItem.Id, qty(1)))
			require.Equal(t, anItem.Price, b.Total())
		})

	t.Run("removing a VALID ITEM from an EMPTY BASKET returns error",
		func(t T) {
			t.Parallel()
			b := basket.NewBasket(catalog)

			// BASKET NOW EMPTY, so should return an error
			require.Error(t, b.Remove(anItem.Id, qty(1)))
			require.Equal(t, zeroCents, b.Total(),
				"total remains zero after removing items from empty basket")
		})
}

const minQty, maxQty = 1, 99

func TestBasket_acceptsValidQuantity(t *testing.T) {
	t.Logf("accepts all quantities q: %d <= q <= %d", minQty, maxQty)
	for q := minQty; q <= maxQty; q++ {
		qty, err := basket.Quantity(q)

		require.NoError(t, err,
			"returned an error for %d", q)

		require.NotPanics(t, func() { evalAndDiscard(*qty) },
			"returned nil quantity for %d", q)
	}
}

func TestBasket_rejectsInvalidQuantity(t *testing.T) {
	someInvalidQuantities := [...]int{math.MinInt, math.MaxInt,
		minQty - 2, minQty - 1, maxQty + 1, maxQty + 2,
	}
	t.Logf("rejects a few invalid quantities q: q < %d or q > %d", minQty, maxQty)
	for _, q := range someInvalidQuantities {
		qty, err := basket.Quantity(q)

		require.Panics(t, func() { evalAndDiscard(*qty) },
			"returned non-nil quantity for %d", q)

		require.Error(t, err,
			"did not return an error for %d", q)
	}
}

func TestBasket_rejectsPuttingTooManyItems(t *testing.T) {
	t.Parallel()
	b := basket.NewBasket(catalog)

	require.NoError(t, b.Put(anItem.Id, qty(1)))
	require.Equal(t, anItem.Price, b.Total())

	assert.Error(t, b.Put(anItem.Id, qty(maxQty)))
	assert.Equal(t, anItem.Price, b.Total(), "total modified on invalid operation")
}

func evalAndDiscard(any) {}
