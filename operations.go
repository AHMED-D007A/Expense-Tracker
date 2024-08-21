package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

	fmt.Printf("Expense added successfully (ID: %v)\n", expense.id)

	// here I will add the feature of warning the user if he exceeds the budget .

	// reader := csv.NewReader(file)

	// records, err := reader.ReadAll()
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }

	// counter, _ := strconv.Atoi(expense.amount)
	// for i, record := range records {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	if time.Now().Month() < 10 {
	// 		if time.Now().Month().String() == time.Month(int(record[4][6])-48).String() {
	// 			val, _ := strconv.Atoi(record[2])
	// 			counter += val
	// 		}
	// 		mm, _ := strconv.Atoi(record[4][5:7])
	// 		if time.Now().Month().String() == time.Month(mm).String() {
	// 			val, _ := strconv.Atoi(record[3])
	// 			counter += val
	// 		}
	// 	}
	// }

	// if counter > Budget {
	// 	fmt.Println("Warning: Total expenses has exceeded the budget, Total:", counter, "budget:", Budget)
	// }
}

func deleteExpense(file *os.File, args ...string) {
	if len(args) < 3 {
		fmt.Println("Usage: ./expense_tracker delete <id>")
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// reading all new records except for the one we are going to delete.
	var newRecords [][]string
	for _, record := range records {
		if record[0] == args[2] {
			continue
		}
		newRecords = append(newRecords, record)
	}

	// cleaning all records
	file.Truncate(0)
	file.Seek(0, 0)

	// writing all records in the newRecords variable.
	writer := csv.NewWriter(file)
	err = writer.WriteAll(newRecords)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer writer.Flush()

	fmt.Println("Expense deleted successfully")
}

func listExpenses(file *os.File, args ...string) {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// checks for the user input.
	if len(args) == 2 {
		for _, record := range records {
			fmt.Println(record)
		}
	} else if len(args) == 4 {
		if args[2] == "--month" || args[2] == "-m" {
			var month string
			if len(args[3]) == 1 {
				month = "0" + args[3]
			}
			for i, record := range records {
				if i == 0 {
					fmt.Println(record)
				} else if record[4][5:7] == month {
					fmt.Println(record)
				}
			}
		} else if args[2] == "--category" || args[2] == "-c" {
			cat := validateCategory(args[3])
			for i, record := range records {
				if i == 0 {
					fmt.Println(record)
				} else if record[3] == string(cat) {
					fmt.Println(record)
				}
			}
		} else {
			fmt.Println("Usage: ./expense_tracker list")
			fmt.Println("Usage[Optional]: ./expense_tracker list -category <category>")
			os.Exit(1)
		}
	} else {
		fmt.Println("Usage: ./expense_tracker list")
		fmt.Println("Usage[Optional]: ./expense_tracker list -category <category>")
		os.Exit(1)
	}
}

func summaryExpenses(file *os.File, args ...string) {
	reader := csv.NewReader(file)
	counter := 0

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if len(args) == 2 {
		for _, record := range records {
			val, _ := strconv.Atoi(record[2])
			counter += val
		}
		fmt.Printf("Total expenses: $%v\n", counter)
	} else if len(args) == 4 {
		if args[2] == "--month" || args[2] == "-m" {
			var month string
			if len(args[3]) == 1 {
				month = "0" + args[3]
			}
			for i, record := range records {
				if i == 0 {
					continue
				}
				if record[4][5:7] == month {
					val, _ := strconv.Atoi(record[2])
					counter += val
				}
			}
			intMonth, _ := strconv.Atoi(month)
			fmt.Printf("Total expenses for %v: $%v\n", time.Month(intMonth).String(), counter)
		} else if args[2] == "--category" || args[2] == "-c" {
			cat := validateCategory(args[3])
			for _, record := range records {
				if record[3] == string(cat) {
					val, _ := strconv.Atoi(record[2])
					counter += val
				}
			}
			fmt.Printf("Total expenses for %v: $%v\n", cat, counter)
		} else {
			fmt.Println("Usage: ./expense_tracker summary")
			fmt.Println("Usage[Optional]: ./expense_tracker summary -category <category>")
			os.Exit(1)
		}
	} else {
		fmt.Println("Usage: ./expense_tracker summary")
		fmt.Println("Usage[Optional]: ./expense_tracker summary -category <category>")
		os.Exit(1)
	}
}

func updateExpense(file *os.File, args ...string) {
	if len(args) < 5 {
		fmt.Println("Usage: ./expense_tracker update <id> --amount <amount>")
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// reading all new records except for the one we are going to delete.
	var newRecords [][]string
	for _, record := range records {
		if record[0] == args[2] {
			if args[3] == "--description" || args[3] == "-d" {
				record[1] = args[4]
			} else if args[3] == "--amount" || args[3] == "-a" {
				record[2] = args[4]
			} else if args[3] == "--category" || args[3] == "-c" {
				record[3] = args[4]
			}
		}
		newRecords = append(newRecords, record)
	}

	// cleaning all records
	file.Truncate(0)
	file.Seek(0, 0)

	// writing all records in the newRecords variable.
	writer := csv.NewWriter(file)
	err = writer.WriteAll(newRecords)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer writer.Flush()

	fmt.Println("Expense updated successfully")
}

// func setBudget(args ...string) {
// 	Budget, err := strconv.Atoi(args[2])
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	fmt.Println("Budget is set to", Budget)
// }
