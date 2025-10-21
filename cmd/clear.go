package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Очистить все задачи",
	RunE: func(cmd *cobra.Command, args []string) error {
		todos.Clear()
		if err := storage.Save(todos); err != nil {
			return err
		}
		fmt.Println("Все задачи удалены")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
