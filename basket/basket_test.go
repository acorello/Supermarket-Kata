package basket_test

import (
	"fmt"
	"io"
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
	_basket := basket.NewBasket(catalog)

	t.Log("reports if an item.Id is not in its catalog")
	_, err := _basket.ItemIdInCatalog("item-not-in-catalog")
	require.Error(t, err)

	t.Log("returns a basket.ItemId if an item.Id is in its catalog")
	// assuming anItem is immutable: sub-tests are reading it
	anItem := catalog.RandomItem()
	anItemId_, err := _basket.ItemIdInCatalog(anItem.Id)
	require.NoErrorf(t, err, "Basket rejected %#v despite being in its catalog", anItem.Id)
	anItemId := *anItemId_

	t.Run("rejects invalid quantities", func(t T) {
		invalidQuantities := [...]int{math.MinInt, -1, 0}
		for _, q := range invalidQuantities {
			qPtr, err := basket.Quantity(q)

			require.Panics(t, func() { fmt.Fprint(io.Discard, *qPtr) },
				"returned non-nil quantity for %d", q)

			require.Error(t, err,
				"did not return an error for %d", q)
		}
	})

	t.Log("accepts valid quantities")
	for q := 1; q < 100; q++ {
		_, err := basket.Quantity(q)
		require.NoError(t, err, "Basket has accepted a quantity of %d", q)
	}

	t.Run("Total() changes as we change an item quantity", func(t T) {
		b := _basket
		require.GreaterOrEqual(t, anItem.Price, 1, "selected item has .Price < 1")
		t.Parallel()

		require.Equal(t, zeroCents, b.Total(), "empty basket total should be zero")

		b.Put(anItemId, qty(1))
		require.Equal(t, anItem.Price, b.Total())

		b.Put(anItemId, qty(2))
		require.Equal(t, anItem.Price*3, b.Total())

		b.Remove(anItemId, qty(1))
		require.Equal(t, anItem.Price*2, b.Total())

		b.Remove(anItemId, qty(2))
		require.Equal(t, zeroCents, b.Total())

		b.Remove(anItemId, qty(1))
		require.Equal(t, zeroCents, b.Total(), "removing an item from an empty basket doesn't change the total")
	})

	t.Run("removing an item from an empty basket doesn't change the total", func(t *testing.T) {
		b := _basket
		t.Parallel()

		anItem := catalog.RandomItem()
		anItemId := must.WorkPtr(b.ItemIdInCatalog(anItem.Id))

		b.Remove(anItemId, qty(1))
		require.Equal(t, zeroCents, b.Total())
	})
}
