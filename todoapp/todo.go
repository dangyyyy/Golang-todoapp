package todoapp

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}
func (todos *Todos) Delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}
func (todos *Todos) Clear() {
	t := *todos
	*todos = t[:0]
	return
}

func (todos *Todos) Toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed
	if !isCompleted {
		compeletime := time.Now()
		t[index].Completed = true
		t[index].CompletedAt = &compeletime
	}
	t[index].Completed = !isCompleted
	return nil
}
func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) Print() {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Задачи", "Выполнено", "Создано", "Завершено")

	for index, t := range *todos {
		completed := "Нет"
		completedAt := ""

		if t.Completed {
			completed = "Да"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format("02.01.2006 15:04")
			}
		}

		createdAt := t.CreatedAt.Format("02.01.2006 15:04")

		tbl.AddRow(
			strconv.Itoa(index),
			t.Title,
			completed,
			createdAt,
			completedAt,
		)
	}

	tbl.Render()
}
