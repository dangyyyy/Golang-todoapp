package cmd

import (
	"todoapp/todoapp"

	"github.com/spf13/cobra"
)

var filteredHigh bool
var filteredMedium bool
var filteredLow bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Вывести в консоль все задачи",
	RunE: func(cmd *cobra.Command, args []string) error {
		todoapp.Print(todos, filteredHigh, filteredMedium, filteredLow)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&filteredHigh, "high", false, "Показать только задачи с высоким приоритетом")
	listCmd.Flags().BoolVar(&filteredMedium, "medium", false, "Показать только задачи со средним приоритетом")
	listCmd.Flags().BoolVar(&filteredLow, "low", false, "Показать только задачи с низким приоритетом")
}
