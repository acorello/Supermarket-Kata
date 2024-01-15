package basket

import "dev.acorello.it/go/supermarket-kata/item"

type inventoryContract struct {
	Inventory
}

func (i inventoryContract) Items(ids []item.Id) (result []item.Item) {
	result = i.Inventory.Items(ids)
	// all ids in basket were present in Inventory because we call Inventory.Knows before
	// adding them to the basket. So if the returned list has fewer items than the Inventory
	// implementation has a flaw.
	if len(result) < len(ids) {
		panic("inventory did not return all requested items")
	}
	return result
}
