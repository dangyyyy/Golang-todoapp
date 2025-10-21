package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var toggleCmd = &cobra.Command{
	Use:   "toggle [id]",
	Short: "Отметить задачу как выполненную/невыполненную",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		if err := todos.Toggle(id); err != nil {
			return err
		}
		if err := storage.Save(todos); err != nil {
			return err
		}
		fmt.Println("Статус задачи изменён")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
