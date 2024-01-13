package discount_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/discount"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
	"github.com/stretchr/testify/assert"
)

const eligibleQuantity = 3
const equivalentQuantity = 2

var eligibleItem = item.Item{
	Id:    item.Id("I_AM_ELIGIBLE"),
	Price: money.Cents(5),
	Unit:  "u",
}

var ineligibleItem = item.Item{
	Id:    item.Id("I_AM_NOT_ELIGIBLE"),
	Price: money.Cents(7),
	Unit:  "u",
}

type ItemIdQuantity item.ItemIdQuantity

func TestThreeForTwoDiscount_Basket_With(t *testing.T) {
	threeFor2 := discount.ThreeForTwo(eligibleItem.Id)

	for name, testCase := range tests {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			got := threeFor2.Discount(tc.input)
			assert.Equal(t, tc.output, got)
		})
	}
}

type testCase struct {
	input  item.ItemsQuantities
	output discount.Output
}

var tests = map[string]testCase{
	"nil": {
		input:  nil,
		output: discount.Output{},
	},
	"empty": {
		input:  item.ItemsQuantities{},
		output: discount.Output{},
	},
	"one eligible item": {
		input: item.ItemsQuantities{
			eligibleItem: 1,
		},
		output: discount.Output{
			Discounted: nil,
			Rest: item.ItemsQuantities{
				eligibleItem: 1,
			},
		},
	},
	"one ineligible item": {
		input: item.ItemsQuantities{
			ineligibleItem: 1,
		},
		output: discount.Output{
			Discounted: nil,
			Rest: item.ItemsQuantities{
				ineligibleItem: 1,
			},
		},
	},
	"eligible-quantity of eligible item": {
		input: item.ItemsQuantities{
			eligibleItem: eligibleQuantity,
		},
		output: discount.Output{
			Discounted: discount.DiscountedGroups{
				{
					DiscountId: discount.IdThreeForTwo,
					Total:      eligibleItem.Price * equivalentQuantity,
					Group: item.ItemsQuantities{
						eligibleItem: eligibleQuantity,
					},
				},
			},
			Rest: nil,
		},
	},
	"eligible-quantity minus one of eligible item": {
		input: item.ItemsQuantities{
			eligibleItem: eligibleQuantity - 1,
		},
		output: discount.Output{
			Discounted: nil,
			Rest: item.ItemsQuantities{
				eligibleItem: eligibleQuantity - 1,
			},
		},
	},
	"eligible-quantity plus one of eligible-item": {
		input: item.ItemsQuantities{
			eligibleItem: eligibleQuantity + 1,
		},
		output: discount.Output{
			Discounted: discount.DiscountedGroups{
				{
					DiscountId: discount.IdThreeForTwo,
					Total:      eligibleItem.Price * equivalentQuantity,
					Group: item.ItemsQuantities{
						eligibleItem: eligibleQuantity,
					},
				},
			},
			Rest: item.ItemsQuantities{
				eligibleItem: 1,
			},
		},
	},
	"twice the eligible-quantity of eligible-item": {
		input: item.ItemsQuantities{
			eligibleItem: eligibleQuantity * 2,
		},
		output: discount.Output{
			Discounted: discount.DiscountedGroups{
				{
					DiscountId: discount.IdThreeForTwo,
					Total:      eligibleItem.Price * equivalentQuantity * 2,
					Group: item.ItemsQuantities{
						eligibleItem: eligibleQuantity * 2,
					},
				},
			},
			Rest: nil,
		},
	},
	"twice the eligible-quantity plus one of eligible-item": {
		input: item.ItemsQuantities{
			eligibleItem: eligibleQuantity*2 + 1,
		},
		output: discount.Output{
			Discounted: discount.DiscountedGroups{
				{
					DiscountId: discount.IdThreeForTwo,
					Total:      eligibleItem.Price * equivalentQuantity * 2,
					Group: item.ItemsQuantities{
						eligibleItem: eligibleQuantity * 2,
					},
				},
			},
			Rest: item.ItemsQuantities{
				eligibleItem: 1,
			},
		},
	},
	"twice the eligible-quantity minus one of eligible-item": {
		input: item.ItemsQuantities{
			eligibleItem: eligibleQuantity*2 - 1,
		},
		output: discount.Output{
			Discounted: discount.DiscountedGroups{
				{
					DiscountId: discount.IdThreeForTwo,
					Total:      eligibleItem.Price * equivalentQuantity,
					Group: item.ItemsQuantities{
						eligibleItem: eligibleQuantity,
					},
				},
			},
			Rest: item.ItemsQuantities{
				eligibleItem: eligibleQuantity - 1,
			},
		},
	},
	"twice the eligible-quantity minus one of eligible-item; one of ineligible-item": {
		input: item.ItemsQuantities{
			eligibleItem:   eligibleQuantity*2 - 1,
			ineligibleItem: 1,
		},
		output: discount.Output{
			Discounted: discount.DiscountedGroups{
				{
					DiscountId: discount.IdThreeForTwo,
					Total:      eligibleItem.Price * equivalentQuantity,
					Group: item.ItemsQuantities{
						eligibleItem: eligibleQuantity,
					},
				},
			},
			Rest: item.ItemsQuantities{
				eligibleItem:   eligibleQuantity - 1,
				ineligibleItem: 1,
			},
		},
	},
	"twice the eligible-quantity minus one of eligible-item; eligible-quantity of ineligible-item": {
		input: item.ItemsQuantities{
			eligibleItem:   eligibleQuantity*2 - 1,
			ineligibleItem: eligibleQuantity,
		},
		output: discount.Output{
			Discounted: discount.DiscountedGroups{
				{
					DiscountId: discount.IdThreeForTwo,
					Total:      eligibleItem.Price * equivalentQuantity,
					Group: item.ItemsQuantities{
						eligibleItem: eligibleQuantity,
					},
				},
			},
			Rest: item.ItemsQuantities{
				eligibleItem:   eligibleQuantity - 1,
				ineligibleItem: eligibleQuantity,
			},
		},
	}}
