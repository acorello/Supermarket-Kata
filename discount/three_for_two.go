package discount

import (
	"slices"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

const IdThreeForTwo = DiscountId("three_for_two")

func ThreeForTwo(ids ...item.Id) threeForTwo {
	// set item.Ids and sort
	sorted := slices.Clone(ids)
	slices.Sort(sorted)
	return threeForTwo{
		itemIds: sorted,
	}
}

type threeForTwo struct {
	itemIds []item.Id
}

func (my threeForTwo) Discount(basket item.ItemsQuantities) (output Output) {
	const ElgibleQuantity = 3
	const EquivalentQuantity = 2
	var matchingItemQuantities item.ItemsQuantities
	for basketItem, qty := range basket {
		if _, found := slices.BinarySearch(my.itemIds, basketItem.Id); !found {
			output.Rest.Add(basketItem, qty)
			continue
		}
		// I'm assuming the given input can contain duplicates
		matchingItemQuantities.Add(basketItem, qty)
	}
	for matchingItem, qty := range matchingItemQuantities {
		rem := qty % ElgibleQuantity
		if rem > 0 {
			output.Rest.Add(matchingItem, rem)
		}
		tripletsCount := qty / ElgibleQuantity
		if tripletsCount > 0 {
			var discountedItems item.ItemsQuantities
			discountedItems.Add(matchingItem, qty-rem)
			discountedTotal := matchingItem.Price *
				EquivalentQuantity *
				// TODO? model quantity so that it can be multiplied to money without cast
				// perhaps interface { ~uint }, or type Quantity = uint, or bare uint
				money.Cents(tripletsCount)
			discountedGroup := DiscountedItems{
				DiscountId: IdThreeForTwo,
				Group:      discountedItems,
				Total:      discountedTotal,
			}
			output.Discounted.Append(discountedGroup)
		}
	}
	return output
}
