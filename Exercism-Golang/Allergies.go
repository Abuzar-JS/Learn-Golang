package main

var products = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	var items []string
	for products, points := range products {
		if allergies&points != 0 {
			items = append(items, products)
		}
	}
	return items
}

func AllergicTo(allergies uint, allergen string) bool {
	panic("Please implement the AllergicTo function")
}
