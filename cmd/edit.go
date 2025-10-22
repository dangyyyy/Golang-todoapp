package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var priorityValue int

var editCmd = &cobra.Command{
	Use:   "edit [индекс] [новое название задачи]",
	Short: "Изменить название или приоритет задачи",
	Long: `Позволяет изменить название или приоритет задачи.
Примеры:
  todo edit 2 "Сходить в зал"   # изменить название
  todo edit 2 -p 1              # изменить приоритет`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("неверный индекс задачи")
		}
		if cmd.Flags().Changed("priority") {
			if priorityValue < 1 || priorityValue > 3 {
				return fmt.Errorf("недопустимый приоритет (1 = высокий, 2 = средний, 3 = низкий)")
			}
			if err := todos.EditTask(id, "", priorityValue); err != nil {
				return err
			}
			fmt.Println("Приоритет задачи обновлён")
		} else {
			if len(args) < 2 {
				return fmt.Errorf("укажите новое название задачи")
			}
			newTitle := args[1]
			if err := todos.EditTask(id, newTitle, 0); err != nil {
				return err
			}
			fmt.Println("Название задачи обновлено")
		}

		if err := storage.Save(todos); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().IntVarP(&priorityValue, "priority", "p", 0, "Изменить приоритет задачи (1 = высокий, 2 = средний, 3 = низкий)")
}
