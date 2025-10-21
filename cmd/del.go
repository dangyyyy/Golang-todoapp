package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del [id]",
	Short: "Удалить задачу",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		if err := todos.Delete(id); err != nil {
			return err
		}
		if err := storage.Save(todos); err != nil {
			return err
		}
		fmt.Println("Задача удалена")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
