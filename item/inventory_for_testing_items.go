package item

const (
	IdBeans       = "beans"
	IdSandwich    = "plaughman_sandwich"
	IdOrangeJuice = "tropicana_orange_juice"
	IdPorridge    = "porridge"
)

var fixedItems = []Item{
	{Id: IdBeans, Price: 90, Unit: "u"},
	{Id: "oranges", Price: 200, Unit: "kg"},
	{Id: "cherries", Price: 700, Unit: "kg"},
	{Id: IdSandwich, Price: 320, Unit: "u"},
	{Id: IdOrangeJuice, Price: 130, Unit: "u"},
	{Id: "innocent_superberries_120ml", Price: 135, Unit: "u"},
	{Id: IdPorridge, Price: 230, Unit: "u"},
}
