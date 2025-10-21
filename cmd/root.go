package cmd

import (
	"todoapp/todoapp"

	"github.com/spf13/cobra"
)

var storage *todoapp.Storage[todoapp.Todos]
var todos todoapp.Todos
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Минималистичный todo CLI",
	Long:  "Простое приложение для управления списком задач с хранением в JSON.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		storage = todoapp.NewStorage[todoapp.Todos]("todos.json")
		return storage.Load(&todos)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
