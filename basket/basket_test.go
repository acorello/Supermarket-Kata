package basket_test

import (
	"math"
	"testing"

	"dev.acorello.it/go/supermarket-kata/basket"
	"dev.acorello.it/go/supermarket-kata/money"
	"dev.acorello.it/go/supermarket-kata/must"
	"dev.acorello.it/go/supermarket-kata/test_fixtures"
	"github.com/stretchr/testify/require"
)

const zeroCents = money.Cents(0)

var catalog = test_fixtures.Catalog()

var qty = must.DereFn(basket.Quantity)

// shorhand for sub-test function signatures
type T = *testing.T

func TestBasket(t *testing.T) {
	// a basket depends on a catalog
	// -> verify that no basket can be created with invalid dependencies
	t.Log("verify constructor rejects invalid arguments")
	require.Panics(t, func() {
		var nilCatalog basket.Catalog = nil
		basket.NewBasket(nilCatalog)
	}, "panic when given nil catalog")

	// create a basket
	_basket := basket.NewBasket(catalog)

	// basket method arguments must be validated using provided validation functions
	// -> verify that validation functions reject invalid inputs
	// -> verify that validation functions accept valid inputs
	rejectsInvalidItemId(t, _basket)

	// assuming anItem is immutable: sub-tests are reading it
	anItem := catalog.RandomItem()
	_anItemId, err := _basket.ItemIdInCatalog(anItem.Id)
	require.NoErrorf(t, err, "%#v rejected despite being in its catalog", anItem.Id)
	anItemId := *_anItemId

	const minQty, maxQty = 1, 99
	rejectsInvalidQuantity(t, minQty, maxQty)
	acceptsValidQuantity(t, minQty, maxQty)

	t.Run("Total() changes as we change an item quantity", func(t T) {
		b := _basket // make copy, we're about to run concurrently!
		t.Parallel()

		require.Greater(t, anItem.Price, 1, "following tests rely on item.Price > 1")

		require.Equal(t, zeroCents, b.Total(), "total of empty basket should be zero")

		b.Put(anItemId, qty(1))
		require.Equal(t, anItem.Price, b.Total())

		b.Put(anItemId, qty(2))
		require.Equal(t, anItem.Price*3, b.Total())

		b.Remove(anItemId, qty(1))
		require.Equal(t, anItem.Price*2, b.Total())

		b.Remove(anItemId, qty(2))
		require.Equal(t, zeroCents, b.Total())

		b.Remove(anItemId, qty(1))
		require.Equal(t, zeroCents, b.Total(),
			"total remains zero after removing items from empty basket")
	})

	t.Run("removing an item from an empty basket doesn't change the total", func(t T) {
		b := _basket
		t.Parallel()

		anItem := catalog.RandomItem()
		anItemId := must.WorkPtr(b.ItemIdInCatalog(anItem.Id))

		b.Remove(anItemId, qty(1))
		require.Equal(t, zeroCents, b.Total())
	})
}

func acceptsValidQuantity(t *testing.T, minQty int, maxQty int) {
	t.Logf("accepts all quantities q: %d <= q <= %d", minQty, maxQty)
	for q := minQty; q <= maxQty; q++ {
		qty, err := basket.Quantity(q)

		require.NoError(t, err,
			"returned an error for %d", q)

		require.NotPanics(t, func() { evalAndDiscard(*qty) },
			"returned nil quantity for %d", q)
	}
}

func rejectsInvalidQuantity(t *testing.T, minQty int, maxQty int) {
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

func rejectsInvalidItemId(t *testing.T, b basket.Basket) {
	itemId, err := b.ItemIdInCatalog("item-not-in-catalog")
	require.Error(t, err, "invalid item.Id passed validation")
	require.Panics(t, func() { evalAndDiscard(*itemId) },
		"non-nil basket.itemId returned for invalid item.Id")
}

func evalAndDiscard(any) {}
