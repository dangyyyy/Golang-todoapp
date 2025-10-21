package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var priority int

var addCmd = &cobra.Command{
	Use:   "add [название задачи]",
	Short: "Добавить новую задачу",
	Long:  "Добавляет новую задачу с возможностью указать приоритет",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]
		todos.Add(title, priority)
		if err := storage.Save(todos); err != nil {
			return err
		}
		fmt.Println("Задача добавлена:", title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Приоритет задачи (1 = высокий, 2 = средний, 3 = низкий)")
}
