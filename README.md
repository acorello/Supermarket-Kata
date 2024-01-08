# Supermarket Kata

Taking inspiration from the [Supermarket Kata] this code implements a basket allowing you to add and remove items; and generate a total, with details of prices and applied discounts.

When adding an item the basket checks if it's a valid item.Id; and if the quantity hasn't exceeded a maximum.

When generating a total; basket delegates to inventory the to check the stock has enough items, and to return the prices and applied discounts.

[Supermarket Kata]: http://codekata.com/kata/kata01-supermarket-pricing/
