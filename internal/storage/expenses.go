package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Expense struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

func LoadExpenses(filename string) ([]Expense, error) {
	var expenses []Expense

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return expenses, nil
		}
		return nil, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []Expense{}, nil // пустой список, если файл пуст
	}

	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func saveExpenses(filename string, expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func AddExpense(filename, description string, amount float64) (int, error) {
	expenses, err := LoadExpenses(filename)
	if err != nil {
		return 0, err
	}

	newID := 1
	if len(expenses) > 0 {
		newID = expenses[len(expenses)-1].ID + 1
	}

	newExpense := Expense{
		ID:          newID,
		Date:        time.Now().Format(time.RFC3339),
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, newExpense)

	return newID, saveExpenses(filename, expenses)
}

func DeleteExpenseByID(filename string, id int) error {
	expenses, err := LoadExpenses(filename)
	if err != nil {
		return err
	}

	newExpenses := make([]Expense, 0, len(expenses))
	found := false

	for _, expense := range expenses {
		if expense.ID == id {
			found = true
			continue
		}
		newExpenses = append(newExpenses, expense)
	}

	if !found {
		return fmt.Errorf("expense with ID %d not found", id)
	}

	return saveExpenses(filename, newExpenses)
}
