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
	Transportation Category = "Transportation"
	Food           Category = "Food"
	Housing        Category = "Housing"
	Medicals       Category = "Medicals"
	Taxes          Category = "Taxes"
	Other          Category = "Other"
)

var Categories = []Category{Transportation, Food, Housing, Medicals, Taxes}

// var Budget int
