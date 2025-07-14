/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AlexanderZah/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list expense",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := storage.LoadExpenses("expenses.json")
		if err != nil {
			fmt.Println("Error loading expenses:", err)
			return
		}
		fmt.Printf("%-4s %-12s %-15s %-6s\n", "ID", "Date", "Description", "Amount")
		for _, expense := range expenses {
			fmt.Printf("%-4d %-12s %-15s $%.2f\n",
				expense.ID, expense.Date[:10], expense.Description, expense.Amount)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
