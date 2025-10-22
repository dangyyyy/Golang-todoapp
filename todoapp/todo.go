package todoapp

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	if err := t.validateIndex(index - 1); err != nil {
		return err
	}
	*todos = append(t[:index-1], t[index:]...)
	return nil
}

func (todos *Todos) Clear() {
	*todos = (*todos)[:0]
}

func (todos *Todos) Toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index - 1); err != nil {
		return err
	}
	isCompleted := t[index-1].Completed
	if !isCompleted {
		compeletime := time.Now()
		t[index-1].Completed = true
		t[index-1].CompletedAt = &compeletime
	}
	t[index-1].Completed = !isCompleted
	return nil
}

func (todos *Todos) Find(title string) error {
	var filtered Todos

	for _, t := range *todos {
		if strings.Contains(strings.ToLower(t.Title), strings.ToLower(title)) {
			filtered = append(filtered, t)
		}
	}

	if len(filtered) == 0 {
		return fmt.Errorf("задачи, содержащие '%s', не найдены", title)
	}

	Print(filtered, false, false, false)
	return nil
}
func (todos *Todos) EditTask(index int, title string, priority int) error {
	t := *todos
	if err := t.validateIndex(index - 1); err != nil {
		return err
	}
	if title != "" {
		t[index-1].Title = title
	}
	if priority >= 1 && priority <= 3 {
		t[index-1].Priority = priority
	}

	return nil
}

func Print(todos Todos, high, medium, low bool) {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Задачи", "Приоритет", "Выполнено", "Создано", "Завершено")

	for index, t := range todos {
		if high && t.Priority != 1 {
			continue
		}
		if medium && t.Priority != 2 {
			continue
		}
		if low && t.Priority != 3 {
			continue
		}

		completed := "Нет"
		completedAt := ""
		priority := ""
		switch t.Priority {
		case 1:
			priority = "Высокий"
		case 2:
			priority = "Средний"
		case 3:
			priority = "Низкий"
		}

		if t.Completed {
			completed = "Да"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format("02.01.2006 15:04")
			}
		}

		createdAt := t.CreatedAt.Format("02.01.2006 15:04")

		tbl.AddRow(strconv.Itoa(index+1), t.Title, priority, completed, createdAt, completedAt)
	}

	tbl.Render()
}
