package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var changePriorityCmd = &cobra.Command{
	Use:   "changeP [индекс] [новый приоритет задачи (1-3)]",
	Short: "Изменение приоритета задачи",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		priority, _ := strconv.Atoi(args[1])
		if err := todos.ChangePriority(id, priority); err != nil {
			return err
		}
		if err := storage.Save(todos); err != nil {
			return err
		}
		fmt.Println("✏Приоритет задачи обновлен")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(changePriorityCmd)
}
