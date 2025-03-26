package models

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var Items []Item

func init() {
	Items = append(Items, Item{ID: 1, Name: "Item 1", Price: 100.00})
	Items = append(Items, Item{ID: 2, Name: "Item 2", Price: 200.00})
	Items = append(Items, Item{ID: 3, Name: "Item 3", Price: 300.00})
	Items = append(Items, Item{ID: 4, Name: "Item 4", Price: 400.00})
	Items = append(Items, Item{ID: 5, Name: "Item 5", Price: 500.00})
}
