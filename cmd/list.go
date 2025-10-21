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
		var filtered todoapp.Todos
		for _, t := range todos {
			if filteredHigh && t.Priority != 1 {
				continue
			}
			if filteredMedium && t.Priority != 2 {
				continue
			}
			if filteredLow && t.Priority != 3 {
				continue
			}
			filtered = append(filtered, t)
		}
		filtered.Print()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&filteredHigh, "high", false, "Показать только задачи с высоким приоритетом")
	listCmd.Flags().BoolVar(&filteredMedium, "medium", false, "Показать только задачи с средним приоритетом")
	listCmd.Flags().BoolVar(&filteredLow, "low", false, "Показать только задачи с низким приоритетом")

}
