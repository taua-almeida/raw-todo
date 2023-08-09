package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	Id          int
	Name        string
	Description string
	Completed   bool
}

type Todos []Todo

var lastGreaterId int = 0

func NewTodo(name string, description string) Todo {
	return Todo{
		Name:        name,
		Description: description,
		Completed:   false,
	}
}

func (t *Todo) Complete() {
	t.Completed = true
}

func (t *Todo) Uncomplete() {
	t.Completed = false
}

func AddTodos(tds Todos, td Todo) Todos {
	if len(tds) == 0 {
		if lastGreaterId == 0 {
			td.Id = 1
			lastGreaterId = 1
		} else {
			td.Id = lastGreaterId
		}

	} else {
		td.Id = lastGreaterId + 1
		lastGreaterId = td.Id
	}
	return append(tds, td)
}

func DeleteTodo(tds Todos, id int) Todos {
	for idx, td := range tds {
		if td.Id == id {
			if id > lastGreaterId {
				lastGreaterId = id
			}
			tds = append(tds[:idx], tds[idx+1:]...)
		}
	}
	return tds
}

func UncompleteTodoById(tds Todos, id int) Todos {
	for idx, td := range tds {
		if td.Id == id {
			tds[idx].Uncomplete()
		}
	}
	return tds
}

func CompleteTodoById(tds Todos, id int) Todos {
	for idx, td := range tds {
		if td.Id == id {
			tds[idx].Complete()
		}
	}
	return tds
}

func ListTodos(tds Todos) {
	pp, err := json.MarshalIndent(tds, "", "  ")
	if err == nil {
		fmt.Println(string(pp))
	}
}

func ListCompletedTodos(tds Todos) Todos {
	completedTodos := make(Todos, 0)
	for _, todo := range tds {
		if todo.Completed {
			completedTodos = append(completedTodos, todo)
		}
	}
	return completedTodos
}

func ListUncompletedTodos(tds Todos) Todos {
	uncompletedTodos := make(Todos, 0)
	for _, todo := range tds {
		if !todo.Completed {
			uncompletedTodos = append(uncompletedTodos, todo)
		}
	}
	return uncompletedTodos
}
