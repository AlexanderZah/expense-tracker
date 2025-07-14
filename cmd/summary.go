/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/AlexanderZah/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "summary expense",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		total := 0

		expenses, err := storage.LoadExpenses("expenses.json")
		if err != nil {
			fmt.Println("Error loading expenses:", err)
			return
		}

		for _, expense := range expenses {
			parsedTime, err := time.Parse(time.RFC3339, expense.Date)
			if err != nil {
				fmt.Println("Ошибка парсинга даты:", err)
				return
			}
			month := parsedTime.Month()
			if monthVar != 0 {
				if monthVar == int(month) {
					total += int(expense.Amount)
				}
			} else {
				total += int(expense.Amount)
			}

		}
		fmt.Printf("Total expenses $%d", total)
	},
}

var monthVar int

func init() {
	summaryCmd.Flags().IntVar(&monthVar, "month", 0, "Месяц для подсчета")
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
