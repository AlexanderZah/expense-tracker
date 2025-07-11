/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AlexanderZah/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		newID, err := storage.AddExpense("expenses.json", descVar, amountVar)
		if err != nil {
			fmt.Println("Ошибка при добавлении:", err)
			return
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", newID)
		fmt.Println("add called")
	},
}

var descVar string
var amountVar float64

func init() {
	addCmd.Flags().StringVarP(&descVar, "description", "d", "", "На что трата")
	addCmd.Flags().Float64VarP(&amountVar, "amount", "a", 0.0, "Сколько потрачено")

	rootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
