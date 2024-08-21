# Expense Tracker
- simple command line interface (CLI) to fetch the recent activity of a GitHub user and display it in the terminal.
- This project is from the project based road map, you can find it [here](https://roadmap.sh/backend/projects).
- You can read more about the project form [here](https://roadmap.sh/projects/expense-tracker).

### Features
- Users can add an expense with a description, amount and category. 
- Users can delete an expense.
- Users can view all expenses.
- Users can view all expenses for a specific category or month (of current year).
- Users can view a summary of all expenses.
- Users can view a summary of expenses for a specific category or month (of current year).
- 

### Usage
```
Usage: ./expense_tracker add --description <description> --amount <amuount>  --category <category>

Usage: ./expense_tracker delete <id>

Usage: ./expense_tracker list
Usage[Optional]: ./expense_tracker list -category <category>

Usage: ./expense_tracker summary
Usage[Optional]: ./expense_tracker summary -category <category>
```
