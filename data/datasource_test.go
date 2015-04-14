package data

import (
	"reflect"
	"testing"
	"time"
)

func TestCreateAndRead(t *testing.T) {
	todoList := []Todo{
		{Name: "Write Interface", Completed: true, Due: time.Date(2015, 4, 14, 0, 0, 0, 0, time.Local)},
		{Name: "Write Tests", Completed: false, Due: time.Date(2015, 4, 14, 0, 0, 0, 0, time.Local)},
		{Name: "Write Implementation", Completed: false, Due: time.Date(2015, 4, 14, 0, 0, 0, 0, time.Local)},
		{Name: "Run test", Completed: false, Due: time.Date(2015, 4, 14, 0, 0, 0, 0, time.Local)},
		{Name: "Get lunch", Completed: false, Due: time.Date(2015, 4, 14, 0, 0, 0, 0, time.Local)},
	}
	ds, err := NewDataSource()
	if err != nil {
		t.Errorf("Cannot create Todos datasource")
	}
	ds.ClearData()
	for _, todo := range todoList {
		err := ds.CreateTodo(&todo)
		if err != nil {
			t.Errorf("Error by creating todo: %v\n", err)
		}
	}

	resTodos := ds.GetTodos()
	if len(resTodos) != len(todoList) {
		t.Errorf("Get/Create todos failed: Inserted %v rows , but Get return %v rows\n", len(todoList), len(resTodos))
	}
	for i := 0; i < len(resTodos); i++ {
		if !reflect.DeepEqual(resTodos[i], todoList[i]) {
			t.Errorf("Expected: %v ### but returned : %v\n", todoList[i], resTodos[i])
		}
	}
}
