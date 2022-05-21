package module1

// Budget stores budget information

type Budge struct {
	Name string
	Max float32
	Items []Item
}

// Item stores item information
type Item struct {
	// Name of the item
	Name string
	Description string
	Price float32
}


