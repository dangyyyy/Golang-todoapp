package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [название задачи]",
	Short: "Добавить новую задачу",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		todos.Add(args[0])
		if err := storage.Save(todos); err != nil {
			return err
		}
		fmt.Println("Задача добавлена:", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
