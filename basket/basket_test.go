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

	b := basket.NewBasket(catalog)

	_, err := b.Add("item-not-in-catalog", basket.Qty(1))
	assert.Error(t, err)
}

func TestBasket_quantity_less_than_one_errors(t *testing.T) {
	t.Parallel()

	var err error

	_, err = basket.Quantity(0)
	assert.Error(t, err)

	_, err = basket.Quantity(-1)
	assert.Error(t, err)
}

func TestBasket_adding_an_item_multiple_times_incr_quantity(t *testing.T) {
	t.Parallel()

	aBasket := basket.NewBasket(catalog)

	var got, err = aBasket.Add(anItem.Id, basket.Qty(1))
	assert.NoError(t, err)
	assert.Equal(t, 1, got)

	got, err = aBasket.Add(anItem.Id, basket.Qty(2))
	assert.NoError(t, err)
	assert.Equal(t, 3, got)
}

func TestBasket_Total(t *testing.T) {
	t.Parallel()

	aBasket := basket.NewBasket(catalog)

	q, _ := aBasket.Add(anItem.Id, basket.Qty(2))
	assert.Equal(t, anItem.Price.Mul(q), aBasket.Total())
}
