package basket_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/basket"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/test_fixtures"
	"github.com/stretchr/testify/assert"
)

var catalog item.Catalog = test_fixtures.Catalog()
var anItem = catalog.RandomItem()

func TestBasket_adding_unkown_item_returns_error(t *testing.T) {
	t.Parallel()

	basket := basket.NewBasket(catalog)

	_, err := basket.Add("item-not-in-catalog", 1)
	assert.Error(t, err)
}

func TestBasket_quantity_less_than_one_errors(t *testing.T) {
	t.Parallel()

	b := basket.NewBasket(catalog)
	var err error

	_, err = b.Add(anItem.Id, 0)
	assert.Error(t, err)

	_, err = b.Add(anItem.Id, -1)
	assert.Error(t, err)
}

func TestBasket_adding_an_item_multiple_times_incr_quantity(t *testing.T) {
	t.Parallel()

	aBasket := basket.NewBasket(catalog)

	var got, err = aBasket.Add(anItem.Id, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, got)

	got, err = aBasket.Add(anItem.Id, 2)
	assert.NoError(t, err)
	assert.Equal(t, 3, got)
}

func TestBasket_Total(t *testing.T) {
	t.Parallel()

	aBasket := basket.NewBasket(catalog)

	const qty = 2
	aBasket.Add(anItem.Id, qty)
	assert.Equal(t, anItem.Price*qty, aBasket.Total())
}
