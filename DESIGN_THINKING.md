# Design Thinking

## First Analisys

Central concepts: Basket, Item, Discount.

### Basket

- is a collection of items (each having a standard price)
- has a sum total equel to the sum of the individual item prices minus any applied discount

### Item

- has a standard price

### Item-Basket

- an item can be added or removed from the basket

### Discount

- an amount to be subtracted from the total of the basket
- a rule that may apply to:
  - individual items
  - some groups of items
  - may be time sensitive

### Observations

The most complicated concept is the discount: there can be many types each with an arbitrary rule.

It's also the part of the business which will vary the most, because discounts are used to react the the ever-changing supply and demand conditions: we have a lot of X in stock and if we don't sell it it will rot; we're loosing market-share let's attract customers by offering discount on product Y and Z.

**Is there anything in common? Anything that we can implement and will not vary?**

In general:

#### Discounts are a function on the whole basket

```go
func (Basket)Amount()(Amount Money, Applied bool)
```

Given that:

- discounts can apply to an arbitrary number of items
- items are added and removed to the basket in unpredictable order

Then:

- the general discount rule has to take into account the whole basket every time it's updated

It's up to each discount instance to figure out if they find the clusters they are looking for; the general algorithm can't know this.

#### In a large organization, discounts may overlap

I'm thinking that a large supermarket will have a large catalog and discounts may be driven by algorithms… so it's conceivable that when we devise a new discount we either have an algorithm that prevents overlaps discounts in the moment when they are proposed; or we need to account for this occurrence and implement a decision later.

Say that a product is matched by both a "meal deal" and as "3x2" (more abstractly _discountA_ and _discountB_ match two intersecting subsets of items), what do we do?

We need a _DiscountPolicy_. For example:

- choose the discount that matched the bigger/smaller group
- choose the discount that resulted in the larger/smaller amount

I foresee the _DiscountPolicy_ to be a filter on the collection of applicable (or applied discounts), but I'll leave the details of the implementation for later (when we'll have concrete instances of discounts); adding a _DiscountPolicy_ may require opening an existing class (perhaps the basket?) but it should be a simple matter of filtering a collection of _Discounts_ before or after they are applied.

### Recap

The _User_ interacts with the _Basket_ by adding\removing items and viewing the pricing report.

I'm assuming we want to update the total, considering each discount, every time the basket is updated.

So, we have to invoke a collection of discounts every time the Basket is updated.

I'm assuming the discount calculations are fast-enough to be implemented synchronously.

I want to inform the user of each Discount applied and which product they were applied to.

### Skeleton

Let's stub an implementation:

```go
type Basket struct {
    discounters []Discounter
    items []Item
}

type Discounter interface {
    // takes in a list of items (normally the basket contents) and returns a list of all clusters of items to which the discount applied, together with the Amount to be subtracted
    // If the discount did not apply it returns an empty list.
    // If the list is not empty then each item will have a non-empty []Item
    // The discounted amount will be non-zero and less than the total of discounted prices.
    // POST-CONDITION: 0 < Amount <= sum(...item.Price)
    Discounts(...Item) []Discount
}

type Discount struct { []Item; Amount }
```
