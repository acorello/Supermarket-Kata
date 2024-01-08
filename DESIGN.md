# Design

## First Analisys

### Central concepts: Basket, Item, Discount

- Basket
  - is a collection of items (each having a standard price)
  - has a sum total equal to the sum of the individual item prices minus any applied discount
- Item
  - has a standard price
- Item-Basket
  - an item can be added or removed from the basket
- Discount
  - an amount to be subtracted from the total of the basket
  - a rule that may apply to:
    - individual items
    - some groups of items
    - may be time sensitive

### Derived Concepts: Discount-Plan

- Discount-Plan: is a collection of discounts with a validity period.

### Observations

The most complicated concept is the discount: there can be many types each with an arbitrary rule.

It's also the part of the business which will vary the most, because discounts are used to react to ever-changing conditions.

**How to implement an evolving discounting-plan?**

#### Discounts

Discounts are grouped in a _discount-plan_, a discount plan has a start and end instant (minute resolution).

When a discount is applied to the basket, it should return two values: the set of item-clusters to which the discount applied and the remaining items; so that we can apply further discounts to the remaining items.

- each discount is a function on the entire Basket returning a cluster of items that matched and the new price for each cluster
- a basket is associated with a _discount-plan_; the basket lifecycle may start within the validity of a _discount-plan_ and extend beyond its end, so we should check if the _discount-plan_ is still valid every time the user interacts with the basket, reset it to the latest one and, if the new _discount-plan_ results in a different pricing, the user should be notified of that.

I'll model the basket as an aggregate root with a _discount-plan_ collaborator.

#### What if discounts overlap?

As the inventory grows and the discount-plans become more sophisticated there may be an _accidental_ or _intentional_ case where an item is matched by multiple configured discounts.

I'm assuming we'll never want to apply more than one discount to an item, so what are we going to about that?

For the moment, I'll assume that the case is _accidental_ and want to just prevent multiple discounts from applying to the same product.

I see two approaches to prevent _accidentally_ overlapping discounts being applied multiple times to the same item.

1) we prevent overlapping discounts from going live in the first place

   Implementing this is trivial as long as discounts are associated to a set of explicit item-ids; once discounts are mapped to implicit sets the checking has to somehow find intersecting sets by property or by expanding them.

2) we don't prevent overlapping discounts from going live but we effectively prevent multiple discounts from being applied to the same items programmatically;

   An implementation may be as simple as (2.1) ensuring a stable ordering of the discounts before they are applied and, by removing items to which a discount has been applied after each pass, we ensures that each item is discounted at most once and deterministically.
  
If the existance of overlapping discounts was _intentional_ we would either use the (2.1) stable sorting strategy detailed above or we'll probably have to revise the implementation to perform some selection algorithm (i.e. use the discount resulting in the higher\lower total).

### Recap

The _User_ interacts with the _Basket_ by adding\removing items and viewing the pricing report; after each change the _Basket_ is queried for the grand total.

I want to inform the user of each Discount applied and which product they were applied to.

Do I need to refresh the basket if it outlives the associated discount-plan lifetime?

- if I do then the basket has an "expiry time" no longer than the discount-plan, and we probably want to make the user aware of that
- we will need to notify the user of what changed (if any) compared to the previous discount-plan
