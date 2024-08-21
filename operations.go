package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func addExpense(file *os.File, args ...string) {
	if len(args) < 6 {
		fmt.Println("Usage: ./expense_tracker add --description <description> --amount <amuount>  --category <category>")
		os.Exit(1)
	}

	var expense Expense
	var argsLine string

	// take the arguments as a line of string to process.
	for _, arg := range args[2:] {
		argsLine += arg + " "
	}

	// process it to check for the description.
	if strings.Contains(argsLine, "--description") || strings.Contains(argsLine, "-d ") {
		// identify the description index so we can retrieve the value after it.
		// we  get the first index of the value by geting the last index of the flag.
		startIndex := strings.Index(argsLine, "--description") + 14
		if startIndex == 14-1 {
			startIndex = strings.Index(argsLine, "-d") + 3
		}
		// we get the last index of the value by getting the first index of the next flag.
		endIndex := startIndex
		for i := endIndex; i < len(argsLine); i++ {
			if argsLine[i] == '-' {
				break
			}
			endIndex++
		}
		expense.description = argsLine[startIndex : endIndex-1]
	} else {
		fmt.Println("Usage: ./expense_tracker add --description <description> --amount <amuount> --category <category>")
		os.Exit(1)
	}

	// process it to check for the amount.
	if strings.Contains(argsLine, "--amount") || strings.Contains(argsLine, "-a ") {
		// identify the amount index so we can retrieve the value after it.
		// we  get the first index of the value by geting the last index of the flag.
		startIndex := strings.Index(argsLine, "--amount") + 9
		if startIndex == 9-1 {
			startIndex = strings.Index(argsLine, "-a") + 3
		}
		// we get the last index of the value by getting the first index of the next flag.
		endIndex := startIndex
		for i := endIndex; i < len(argsLine); i++ {
			if argsLine[i] == '-' {
				break
			}
			endIndex++
		}
		expense.amount = argsLine[startIndex : endIndex-1]
	} else {
		fmt.Println("Usage: ./expense_tracker add --description <description> --amount <amuount> --category <category>")
		os.Exit(1)
	}

	// process it to check for the category.
	if strings.Contains(argsLine, "--category") || strings.Contains(argsLine, "-c ") {
		// identify the category index so we can retrieve the value after it.
		// we  get the first index of the value by geting the last index of the flag.
		startIndex := strings.Index(argsLine, "--category") + 11
		if startIndex == 11-1 {
			startIndex = strings.Index(argsLine, "-c") + 3
		}
		// we get the last index of the value by getting to the end of the line.
		endIndex := startIndex
		for i := endIndex; i < len(argsLine); i++ {
			if argsLine[i] == '-' {
				break
			}
			endIndex++
		}
		category := argsLine[startIndex : endIndex-1]
		cat := validateCategory(category)
		expense.category = cat
	} else {
		fmt.Println("Usage: ./expense_tracker add --description <description> --amount <amuount> --category <category>")
		os.Exit(1)
	}

	expense.id = idGen()
	expense.date = time.Now().Format(time.DateOnly)

	row := []string{expense.id, expense.description, expense.amount, string(expense.category), expense.date}
	writer := csv.NewWriter(file)

	err := writer.Write(row)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer writer.Flush()
}

func deleteExpense(file *os.File, args ...string) {
}

func listExpenses(file *os.File, args ...string) {}

func summaryExpenses(file *os.File, args ...string) {}
