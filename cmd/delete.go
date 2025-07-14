/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AlexanderZah/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete expense",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := storage.DeleteExpenseByID("expenses.json", int(deleteIdVar))
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Printf("Expense ID=%d deleted successfully", deleteIdVar)
	},
}

var deleteIdVar int

func init() {
	deleteCmd.Flags().IntVar(&deleteIdVar, "id", 0, "Удаляемый айдишник")
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
