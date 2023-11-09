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

	_, err := basket.Add("item-not-in-catalog", quantity(1))
	assert.Error(t, err)
}

func TestBasket_quantity_less_than_one_errors(t *testing.T) {
	t.Parallel()

	var err error

	_, err = basket.QuantityOf(0)
	assert.Error(t, err)

	_, err = basket.QuantityOf(-1)
	assert.Error(t, err)
}

func TestBasket_adding_an_item_multiple_times_incr_quantity(t *testing.T) {
	t.Parallel()

	aBasket := basket.NewBasket(catalog)

	var got, err = aBasket.Add(anItem.Id, quantity(1))
	assert.NoError(t, err)
	assert.Equal(t, quantity(1), got)

	got, err = aBasket.Add(anItem.Id, quantity(2))
	assert.NoError(t, err)
	assert.Equal(t, quantity(3), got)
}

func TestQuantityOfInt_returns_error_if_LE_zero(t *testing.T) {
	var err error
	_, err = basket.QuantityOf(0)
	assert.Error(t, err)
	_, err = basket.QuantityOf(-1)
	assert.Error(t, err)
}

func quantity(q int) basket.Quantity {
	if v, err := basket.QuantityOf(q); err != nil {
		panic(err)
	} else {
		return v
	}
}
