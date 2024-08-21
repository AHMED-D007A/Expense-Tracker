package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var file *os.File

	info, err := os.Stat("expenses.csv")
	if os.IsNotExist(err) {
		file, err = os.Create("expenses.csv")
		if err != nil {
			log.Fatalln(err.Error())
		}
		file.WriteString("id,description,amount,category,date\n")
	} else {
		file, err = os.OpenFile("expenses.csv", os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if info.Size() == 0 {
			file.WriteString("id,description,amount,category,date\n")
		}
	}

	defer file.Close()

	if len(os.Args) == 1 {
		fmt.Println("Usage: [command] [...arguments]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addExpense(file, os.Args...)
	case "delete":
		deleteExpense(file, os.Args...)
	case "list":
		listExpenses(file)
	case "summary":
		summaryExpenses(file)
	}
}
