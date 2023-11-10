package basket_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/basket"
	"dev.acorello.it/go/supermarket-kata/must"
	"dev.acorello.it/go/supermarket-kata/test_fixtures"
	"github.com/stretchr/testify/assert"
)

var catalog = test_fixtures.Catalog()

var qty = must.Fn(basket.Quantity)

func TestBasketKnownItem(t *testing.T) {
	t.Parallel()

	b := basket.NewBasket(catalog)

	_, err := b.KnownItemId("item-not-in-catalog")
	assert.Error(t, err)
}

func TestBasket_adding_an_item_multiple_times_incr_quantity(t *testing.T) {
	t.Parallel()

	aBasket := basket.NewBasket(catalog)
	anItem := catalog.RandomItem()
	anItemId := must.Work(aBasket.KnownItemId(anItem.Id))

	var got = aBasket.Add(anItemId, qty(1))
	assert.Equal(t, 1, got)

	got = aBasket.Add(anItemId, qty(2))
	assert.Equal(t, 3, got)
}

func TestBasketRemove(t *testing.T) {
	t.Parallel()

	aBasket := basket.NewBasket(catalog)
	anItem := catalog.RandomItem()
	anItemId := must.Work(aBasket.KnownItemId(anItem.Id))

	var got = aBasket.Add(anItemId, qty(2))
	assert.Equal(t, 2, got)

	got = aBasket.Remove(anItemId, qty(1))
	assert.Equal(t, 1, got)

	got = aBasket.Remove(anItemId, qty(1))
	assert.Equal(t, 0, got)

	got = aBasket.Remove(anItemId, qty(1))
	assert.Equal(t, 0, got)
}

func TestBasket_Total(t *testing.T) {
	t.Parallel()

	aBasket := basket.NewBasket(catalog)
	anItem := catalog.RandomItem()
	anItemId := must.Work(aBasket.KnownItemId(anItem.Id))

	q := aBasket.Add(anItemId, qty(2))
	assert.Equal(t, anItem.Price.Mul(q), aBasket.Total())
}
