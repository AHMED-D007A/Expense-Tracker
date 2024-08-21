package main

type Expense struct {
	id          string
	description string
	amount      string
	category    Category
	date        string
}

type Category string

const (
	Transportation Category = "TRANSPORTATION"
	Food           Category = "FOOD"
	Housing        Category = "HOUSING"
	Medicals       Category = "MEDICALS"
	Taxes          Category = "TAXES"
	Other          Category = "OTHER"
)

var Categories = []Category{Transportation, Food, Housing, Medicals, Taxes}

// var Budget int
