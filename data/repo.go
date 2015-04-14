package data

import "fmt"

var currentID int
var todos Todos

func init() {
	t := Todo{Name: "Write Webservice: Go & Gorilla"}
	// fmt.Printf("Todo %s ### original:addr(t) = %p\n", t.Name, &t)
	RepoCreateTodo(t)
	RepoCreateTodo(Todo{Name: "Write Full website: Go & Gorilla"})
}

func RepoListTodos() Todos {
	return todos
}
func RepoCreateTodo(t Todo) Todo {
	// fmt.Printf("Todo %s ### local:addr(t) = %p\n", t.Name, &t)
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	return Todo{}
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[0:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Todo wirth id = %v not found", id)
}
