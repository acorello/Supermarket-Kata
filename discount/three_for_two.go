package discount

import (
	"slices"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

const IdThreeForTwo = DiscountId("three_for_two")

func NewThreeForTwo(ids ...item.Id) threeForTwo {
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
	var candidateItemsQtys item.ItemsQuantities
	for basketItem, qty := range basket {
		if !my.isDiscounted(basketItem) {
			output.Rest.Add(basketItem, qty)
			continue
		}
		candidateItemsQtys.Add(basketItem, qty)
	}
	for candidateItem, qty := range candidateItemsQtys {
		rem := qty % ElgibleQuantity
		if rem > 0 {
			output.Rest.Add(candidateItem, rem)
		}
		eligibleQty := qty / ElgibleQuantity
		if eligibleQty > 0 {
			var discountedItems item.ItemsQuantities
			discountedItems.Add(candidateItem, qty-rem)
			discountedGroup := DiscountedItems{
				DiscountId: IdThreeForTwo,
				Group:      discountedItems,

				// TODO? model quantity so that it can be multiplied to money without cast
				// perhaps interface { ~uint }, or type Quantity = uint, or bare uint
				Total: candidateItem.Price * EquivalentQuantity *
					money.Cents(eligibleQty),
			}
			output.Discounted.Append(discountedGroup)
		}
	}
	return output
}

func (my threeForTwo) isDiscounted(i item.Item) bool {
	_, found := slices.BinarySearch(my.itemIds, i.Id)
	return found
}
