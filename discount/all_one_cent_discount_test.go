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

	output := discounter.Discount(item.ItemsQuantities{
		anItem: 1,
	})
	assert.Nil(t, output.Rest)
	assert.Len(t, output.Discounted, 1)
	discountedPrice := output.Discounted[0].Total
	require.Equal(t, money.Cents(1), discountedPrice)

	output = discounter.Discount(item.ItemsQuantities{
		anItem: 2,
	})
	assert.Nil(t, output.Rest)
	assert.Len(t, output.Discounted, 1)
	discountedPrice = output.Discounted[0].Total
	require.Equal(t, money.Cents(2), discountedPrice)

	output = discounter.Discount(item.ItemsQuantities{
		anItem:      2,
		anotherItem: 2,
	})
	assert.Nil(t, output.Rest)
	assert.Len(t, output.Discounted, 2)

	discountedPrice = output.Discounted[0].Total
	require.Equal(t, money.Cents(2), discountedPrice)
	discountedPrice = output.Discounted[1].Total
	require.Equal(t, money.Cents(2), discountedPrice)
}
