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
	Priority    int
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string, priority int) {
	if priority < 1 || priority > 3 {
		err := errors.New("invalid priority, choose from 1 to 3")
		fmt.Println(err)
		os.Exit(1)
	}
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		Priority:    priority,
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
func (todos *Todos) EditTask(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}
func (todos *Todos) ChangePriority(index int, priority int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Priority = priority
	return nil
}

func (todos *Todos) Print() {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Задачи", "Приоритет", "Выполнено", "Создано", "Завершено")

	for index, t := range *todos {
		completed := "Нет"
		completedAt := ""
		priority := ""
		switch t.Priority {
		case 1:
			priority = "Низкий"
		case 2:
			priority = "Средний"
		case 3:
			priority = "Высокий"
		}

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
			priority,
			completed,
			createdAt,
			completedAt,
		)
	}

	tbl.Render()
}
