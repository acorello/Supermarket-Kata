package basket_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/basket"
	"dev.acorello.it/go/supermarket-kata/discount"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const zeroCents = money.Cents(0)

var inventory = item.FixedInventory()
var noDiscounts = discount.NoDiscounts()
var someItems = inventory.RandomItems(2)
var anItem, anotherItem = someItems[0], someItems[1]

func NewBasket() basket.Basket {
	return basket.NewBasket(inventory, noDiscounts)
}

// shorhand for sub-test function signatures
type T = *testing.T

func TestBasket(t *testing.T) {
	// -> verify that no basket can be created with invalid dependencies
	require.Panics(t, func() {
		var nilArg basket.Inventory = nil
		basket.NewBasket(nilArg, noDiscounts)
	}, "panic when given nil inventory")
	require.Panics(t, func() {
		var nilArg basket.Discounts = nil
		basket.NewBasket(inventory, nilArg)
	}, "panic when given nil discounts")

	t.Run("Putting an item not in inventory returns an error", func(t T) {
		t.Parallel()
		b := NewBasket()

		invalidItem := item.Id("item-not-in-inventory")
		err := b.Put(invalidItem, 1)
		require.Error(t, err, "basket seems to have accepted an item not in inventory")
	})

	require.NotEqual(t, anItem, anotherItem, "inventory should return a different item")

	t.Run("Total changes as we change an item quantity", func(t T) {
		// TODO: rename b to basket…
		t.Parallel()
		b := NewBasket()

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
			b := NewBasket()

			require.NoError(t, b.Put(anItem.Id, 1))

			require.Error(t, b.Remove(anotherItem.Id, 1))
			require.Equal(t, anItem.Price, b.Total())
		})

	t.Run("removing a VALID ITEM from an EMPTY BASKET returns error",
		func(t T) {
			t.Parallel()
			b := NewBasket()

			// BASKET NOW EMPTY, so should return an error
			require.Error(t, b.Remove(anItem.Id, 1))
			require.Equal(t, zeroCents, b.Total(),
				"total remains zero after removing items from empty basket")
		})
}

func TestBasket_rejectsPuttingZeroItems(t *testing.T) {
	t.Parallel()
	b := NewBasket()

	initTotal := b.Total()
	require.Error(t, b.Put(anItem.Id, 0))
	require.Equal(t, initTotal, b.Total())
}

func TestBasket_rejectsPuttingTooManyItems(t *testing.T) {
	t.Parallel()
	b := NewBasket()

	require.NoError(t, b.Put(anItem.Id, 1))
	require.Equal(t, anItem.Price, b.Total())

	assert.Error(t, b.Put(anItem.Id, basket.MaxQuantity))
	assert.Equal(t, anItem.Price, b.Total(), "total modified on invalid operation")
}

func TestBasket_Discounts(t *testing.T) {
	t.Run("all one cent discount", func(t T) {
		t.Parallel()

		// GIVEN: basket uses all-one-cent-discount
		b := basket.NewBasket(inventory, discount.AllOneCentDiscount())
		// GIVEN: basket contains an item costing >1cent
		require.Greater(t, anItem.Price, money.Cents(1))
		require.NoError(t, b.Put(anItem.Id, 1))

		// THEN: total is 1-cent
		require.Equal(t, money.Cents(1), b.Total())

		// and 2 cents with two counts of anItem
		require.NoError(t, b.Put(anItem.Id, 1))
		require.Equal(t, money.Cents(2), b.Total())

		// adding another item with a different price still greater than 1-cent
		require.Greater(t, anotherItem.Price, money.Cents(1))
		require.NotEqual(t, anItem.Price, anotherItem.Price)

		require.NoError(t, b.Put(anotherItem.Id, 1))
		require.Equal(t, money.Cents(3), b.Total())
	})
}
