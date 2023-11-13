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

var qty = must.Fn(basket.Quantity)

// shorhand for sub-test function signatures
type T = *testing.T

func TestBasket(t *testing.T) {
	_basket := basket.NewBasket(catalog)

	t.Log("reports if an item.Id is not in its catalog")
	_, err := _basket.ItemIdInCatalog("item-not-in-catalog")
	require.Error(t, err)

	t.Log("returns a basket.ItemId if an item.Id is in its catalog")
	anItem := catalog.RandomItem()
	// assuming anItem is immutable: sub-tests are reading it
	anItemId, err := _basket.ItemIdInCatalog(anItem.Id)
	require.NoErrorf(t, err, "Basket rejected %#v despite being in its catalog", anItem.Id)

	t.Log("rejects invalid quantities")
	for _, q := range []int{math.MinInt, -1, 0} {
		_, err := basket.Quantity(q)
		require.Error(t, err, "Basket has accepted a quantity of %d", q)
	}

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
		anItemId := must.Work(b.ItemIdInCatalog(anItem.Id))

		b.Remove(anItemId, qty(1))
		require.Equal(t, zeroCents, b.Total())
	})

	t.Run("panic when using a nullish quantity", func(t *testing.T) {
		b := _basket
		t.Parallel()

		anItem := catalog.RandomItem()
		anItemId := must.Work(b.ItemIdInCatalog(anItem.Id))

		doubiousQuantity, _ := basket.Quantity(-1)

		require.Panics(t, func() { b.Put(anItemId, doubiousQuantity) }, "put panics")
		require.Panics(t, func() { b.Remove(anItemId, doubiousQuantity) }, "remove panics")
	})
}
