package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var file *os.File

	_, err := os.Stat("expenses.csv")
	if os.IsNotExist(err) {
		file, err = os.Create("expenses.csv")
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		file, err = os.OpenFile("expenses.csv", os.O_RDWR, 0644)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	defer file.Close()

	fmt.Println(file.Name())
}
