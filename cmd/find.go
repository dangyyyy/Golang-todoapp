package cmd

import (
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find [значение]",
	Short: "Поиск по задачам",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := todos.Find(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

}
