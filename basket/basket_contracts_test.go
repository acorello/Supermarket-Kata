package basket

import (
	"dev.acorello.it/go/supermarket-kata/discount"
	"dev.acorello.it/go/supermarket-kata/item"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
# OVERVIEW

-> ONE LESS OF EACH ITEM: A,B,C
-> ONE MORE OF EACH ITEM: A,B,C
-> EACH MISSING
-> AN EXTRA "ITEM-X" IN EACH GROUP
-> ALL ITEMS IN UND and NIL|EMPTY DIS (2 cases)
-> ALL ITEMS IN ONE DIS and NIL|EMPTY UND (2 cases)
-> AN ITEM IS SPREAD BETWEEN DIS and UND

TODO: IF AN ITEM APPEARS IN MULTIPLE DISCOUNTS WE SHOULD PANIC
TODO? I'm writing a requirement (eg. `require.Contains(t, baseCase.Rest, itemFP)`) just before I
	amend the base case to make it invalid. The check serves to ensure modifications to the test
	data do not accidentally result in false positives/negatives, but that's cumbersome.
	Consider introducing self-checking methods such as `.AddIfPresent(itemId, 1)`
*/

func TestDiscountsContract_Discount(t *testing.T) {
	itemFP := item.Item{Id: "ITEM_FULL_PRICE", Price: 3, Unit: "ml"}
	itemD1 := item.Item{Id: "ITEM_D1", Price: 5, Unit: "kg"}
	itemD2 := item.Item{Id: "ITEM_D2", Price: 7, Unit: "u"}
	itemX := item.Item{Id: "ITEM_X", Price: 17, Unit: "u"}
	input := item.ItemsQuantities{itemFP: 2, itemD1: 4, itemD2: 6}
	baseCase := func() discount.Output {
		return discount.Output{
			Rest: item.ItemsQuantities{itemFP: 2},
			Discounted: discount.DiscountedGroups{
				{
					Id:    discount.Id("DISCOUNT-1"),
					Group: item.ItemsQuantities{itemD1: 4},
				},
				{
					Id:    discount.Id("DISCOUNT-2"),
					Group: item.ItemsQuantities{itemD2: 6},
				},
			},
		}
	}
	tests := map[string]struct {
		output      discount.Output
		shouldPanic bool
	}{
		// BASE NORMAL CASES: all items are present in exact quantities, we tolerate nil\empty collections
		"all items are present in exact quantities, one item full-price, two discounted each by a distinct discounts": {
			shouldPanic: false,
			output:      baseCase(),
		},
		"all items are present as full-price, discount groups is nil": {
			shouldPanic: false,
			output: discount.Output{
				Rest:       input,
				Discounted: nil,
			},
		},
		"all items are present as full-price, discount groups is empty": {
			shouldPanic: false,
			output: discount.Output{
				Rest:       input,
				Discounted: discount.DiscountedGroups{},
			},
		},
		"all items are present in a discount-group, full-price group is nil": {
			shouldPanic: false,
			output: discount.Output{
				Rest:       nil,
				Discounted: discount.DiscountedGroups{{Group: input}},
			},
		},
		"all items are present in a discount-group, full-price group is empty": {
			shouldPanic: false,
			output: discount.Output{
				Rest:       item.ItemsQuantities{},
				Discounted: discount.DiscountedGroups{{Group: input}},
			},
		},
		"an item is partially in FULL_PRICE and partially in DISCOUNT_2": {
			shouldPanic: false,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Discounted[1].Group, itemD2)
				require.Greater(t, baseCase.Discounted[1].Group[itemD2], item.Quantity(1))
				baseCase.Rest[itemD2] = 1
				baseCase.Discounted[1].Group[itemD2] -= 1
			}),
		},
		// FOR EACH GROUP, ALL ITEMS ARE PRESENT BUT ONE TOO MANY OR ONE TOO FEW
		"full-price items contains ITEM_FULL_PRICE but with quantity plus 1": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Rest, itemFP)
				baseCase.Rest[itemFP] += 1
			}),
		},
		"full-price items contains ITEM_FULL_PRICE but with quantity minus 1": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Rest, itemFP)
				baseCase.Rest[itemFP] -= 1
			}),
		},
		"first discount has ITEM_D1 but with quantity plus 1": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Discounted[0].Group, itemD1)
				baseCase.Discounted[0].Group[itemD1] += 1
			}),
		},
		"first discount has ITEM_D1 but with quantity minus 1": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Discounted[0].Group, itemD1)
				baseCase.Discounted[0].Group[itemD1] -= 1
			}),
		},
		"second discount has ITEM_D2 but with quantity plus 1": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Discounted[1].Group, itemD2)
				baseCase.Discounted[1].Group[itemD2] += 1
			}),
		},
		"second discount has ITEM_D2 but with quantity minus 1": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Discounted[1].Group, itemD2)
				baseCase.Discounted[1].Group[itemD2] -= 1
			}),
		},
		// FOR EACH GROUP, AN ITEM IS MISSING
		"ITEM_FULL_PRICE was not returned": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Rest, itemFP)
				delete(baseCase.Rest, itemFP)
			}),
		},
		"ITEM_D1 was not returned": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Discounted[0].Group, itemD1)
				delete(baseCase.Discounted[0].Group, itemD1)
			}),
		},
		"ITEM_D2 was not returned": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				require.Contains(t, baseCase.Discounted[1].Group, itemD2)
				delete(baseCase.Discounted[1].Group, itemD2)
			}),
		},
		// FOR EACH GROUP, THERE IS ONE EXTRA ITEM
		"there is an extra item in FULL_PRICE group": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				baseCase.Rest[itemX] = 1
			}),
		},
		"there is an extra item in DISCOUNT_1 group": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				baseCase.Discounted[0].Group[itemX] = 1
			}),
		},
		"there is an extra item in DISCOUNT_2 group": {
			shouldPanic: true,
			output: But(baseCase(), func(baseCase *discount.Output) {
				baseCase.Discounted[1].Group[itemX] = 1
			}),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			contract := discountsContract{
				Discounts: FixedOutput{tt.output},
			}
			if tt.shouldPanic {
				assert.Panics(t, func() { contract.Discount(input) })
			} else {
				assert.NotPanics(t, func() { contract.Discount(input) })
			}
		})
	}
}

type FixedOutput struct {
	discount.Output
}

func (me FixedOutput) Discount(items item.ItemsQuantities) discount.Output {
	return me.Output
}

func But(base discount.Output, modify func(*discount.Output)) discount.Output {
	modify(&base)
	return base
}
