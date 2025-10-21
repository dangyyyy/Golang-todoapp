package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [название задачи]",
	Short: "Добавить новую задачу",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		prior, _ := strconv.Atoi(args[1])
		todos.Add(args[0], prior)
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
