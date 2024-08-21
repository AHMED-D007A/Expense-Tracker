package main

import (
	"math/rand"
	"strconv"
)

func idGen() string {
	id := rand.Intn(900000) + 100000
	strID := strconv.Itoa(id)
	return strID
}

func validateCategory(category string) Category {
	var cat Category = Category(category)
	for _, v := range Categories {
		if cat == v {
			return v
		}
	}
	return Other
}
