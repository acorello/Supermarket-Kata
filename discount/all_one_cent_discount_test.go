package discount_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/discount"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var discounter = discount.AllOneCentDiscount()

var inventory = item.FixedInventory()
var someItems = inventory.RandomItems(2)
var anItem, anotherItem = someItems[0], someItems[1]

func TestAllOneCentDiscount(t *testing.T) {
	require.Greater(t, anItem.Price, money.Cents(1))
	require.Greater(t, anotherItem.Price, money.Cents(1))

	discounted, undiscounted := discounter.Discount([]item.ItemQuantity{
		{
			Item:     anItem,
			Quantity: 1,
		},
	})
	assert.Nil(t, undiscounted)
	assert.Len(t, discounted, 1)
	discountedPrice := discounted[0].Total
	require.Equal(t, money.Cents(1), discountedPrice)

	discounted, undiscounted = discounter.Discount([]item.ItemQuantity{
		{
			Item:     anItem,
			Quantity: 2,
		},
	})
	assert.Nil(t, undiscounted)
	assert.Len(t, discounted, 1)
	discountedPrice = discounted[0].Total
	require.Equal(t, money.Cents(2), discountedPrice)

	discounted, undiscounted = discounter.Discount([]item.ItemQuantity{
		{
			Item:     anItem,
			Quantity: 2,
		},
		{
			Item:     anotherItem,
			Quantity: 2,
		},
	})
	assert.Nil(t, undiscounted)
	assert.Len(t, discounted, 2)

	discountedPrice = discounted[0].Total
	require.Equal(t, money.Cents(2), discountedPrice)
	discountedPrice = discounted[1].Total
	require.Equal(t, money.Cents(2), discountedPrice)
}
