package basket_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/basket"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/test_fixtures"
	"github.com/stretchr/testify/assert"
)

var catalog item.Catalog = test_fixtures.Catalog()

func TestBasket_ItemCount(t *testing.T) {
	basket := basket.NewBasket(catalog)
	assert.Equalf(t, 0, basket.ItemsCount(), "new basket ItemCount should be zero")
}

func TestBasket_Add(t *testing.T) {
	t.Parallel()

	basket := basket.NewBasket(catalog)
	Add := func(desc string, itemId item.Id, qty int, want int) {
		t.Helper()
		t.Logf(desc, itemId, qty)
		if err := basket.Add(itemId, qty); err != nil {
			t.Error()
		}
		got := basket.ItemsCount()
		if want != got {
			t.Fatalf("\n  want: %d\n   got: %d", want, got)
		}
	}

	const id = "item-not-in-catalog"
	assert.Error(t, basket.Add(id, 1))

	someItems := catalog.FetchRandomItems(2)
	item1 := someItems[0]
	item2 := someItems[1]

	Add("emptyBasket.Add(%#v, %#v)", item1.Id, 1, 1)

	Add("basketWithItem1.Add(%#v, %#v)", item1.Id, 1, 1)

	Add("basketWithItem1x2.Add(%#v, %#v)", item2.Id, 1, 2)
}
