package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Вывести в консоль все задачи",
	RunE: func(cmd *cobra.Command, args []string) error {
		todos.Print()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
