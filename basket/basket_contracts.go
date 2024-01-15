package basket

import (
	"cmp"
	"dev.acorello.it/go/supermarket-kata/discount"
	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/maps_util"
	"maps"
	"slices"
)

type inventoryContract struct{ Inventory }

func (i inventoryContract) Items(ids []item.Id) (result []item.Item) {
	// TODO: test contract violations are detected
	result = i.Inventory.Items(ids)
	// Result should contain all requested ids, exactly.
	// We should never be the case that basket requests items the inventory doesn't know, such items
	// should not even be in the basket; and to ensure that we should call Inventory.Knows.
	slices.Sort(ids)
	slices.SortFunc(result, func(a, b item.Item) int {
		return cmp.Compare(a.Id, b.Id)
	})
	identical := slices.EqualFunc(ids, result, func(id item.Id, i item.Item) bool {
		return id == i.Id
	})
	if !identical {
		panic("result contains different item.Id/s than requested")
	}
	return result
}

type discountsContract struct{ Discounts }

func (me discountsContract) Discount(items item.ItemsQuantities) (result discount.Output) {
	// TODO: test contract violations are detected
	sum := func(_ item.Item, l, r item.Quantity) item.Quantity { return l + r }

	result = me.Discounts.Discount(items)

	// combined output quantities should equal input quantities
	aggregateQuantities := make(item.ItemsQuantities, len(items))
	maps_util.Merge(sum, aggregateQuantities, result.Rest)
	for _, d := range result.Discounted {
		maps_util.Merge(sum, aggregateQuantities, d.Group)
	}
	if !maps.Equal(aggregateQuantities, items) {
		panic("missing items or quantities from discount.Output")
	}
	return result
}
