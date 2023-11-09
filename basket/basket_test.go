package basket_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/basket"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/test_fixtures"
	"github.com/stretchr/testify/assert"
)

var catalog item.Catalog = test_fixtures.Catalog()

func TestBasket_Add(t *testing.T) {
	t.Parallel()

	basket := basket.NewBasket(catalog)

	Add := func(desc string, itemId item.Id, qty int, want int) {
		t.Helper()
		t.Logf(desc, itemId, qty)
		if got, err := basket.Add(itemId, qty); err != nil {
			t.Error()
		} else if got != want {
			t.Fatalf("\n  want: %d\n   got: %d", want, got)
		}
	}

	_, err := basket.Add("item-not-in-catalog", 1)
	assert.Error(t, err)

	someItems := catalog.FetchRandomItems(2)
	item1 := someItems[0]
	item2 := someItems[1]

	Add("emptyBasket.Add(%#v, %#v)", item1.Id, 1, 1)

	Add("basketWithItem1.Add(%#v, %#v)", item1.Id, 1, 2)

	Add("basketWithItem1x2.Add(%#v, %#v)", item2.Id, 1, 1)
}
