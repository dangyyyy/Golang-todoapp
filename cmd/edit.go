package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [индекс] [новое название задачи]",
	Short: "Изменение названия задачи",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		if err := todos.EditTask(id, args[1]); err != nil {
			return err
		}
		if err := storage.Save(todos); err != nil {
			return err
		}
		fmt.Println("✏Задача обновлена")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
