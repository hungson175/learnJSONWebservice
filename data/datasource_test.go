package data

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestAll(test *testing.T) {
	log.SetFlags(log.Lshortfile)
	fmt.Println("Run test for datasource")
	ds, err := NewDataSource()
	if err != nil {
		panic(err)
	}
	defer ds.Close()
	todos := ds.GetTodos()
	//restore
	defer func() {
		ds.ClearData()
		ds.InsertWithExactID(todos)
	}()
	testFunctions := []func(ds *DataSource, test *testing.T){
		testCreateAndRead,
		testUpdate,
		testDelete,
	}
	for _, fn := range testFunctions {
		fn(ds, test)
	}
}

//TODO: Wrapper all test to recover the state of the database as original
//TODO: "uncomment" testUpdate, testDelete
// func WrapperTestFuction(fn func(t *testing.T)) func(t *testing.T) {

// }
func testCreateAndRead(ds *DataSource, t *testing.T) {
	// ds, err := NewDataSource()
	// if err != nil {
	// 	t.Errorf("Cannot create Todos datasource")
	// }
	// defer ds.Close()
	todoList, _ := createSampleData(ds, t)
	resTodos := ds.GetTodos()
	if len(resTodos) != len(todoList) {
		t.Errorf("Get/Create todos failed: Inserted %v rows , but Get return %v rows\n", len(todoList), len(resTodos))
	}
	for i := 0; i < len(resTodos); i++ {
		if !todoEqualsExceptID(&resTodos[i], &todoList[i]) {
			t.Errorf("Expected: %v ### but returned : %v\n", todoList[i], resTodos[i])
		}
		item, err := ds.GetTodo(resTodos[i].ID)
		if err != nil {
			t.Errorf("Expected: %v ### error return %v\n", resTodos[i], err)
		} else {
			if !reflect.DeepEqual(item, resTodos[i]) {
				t.Errorf("Expected: %v ### but returned : %v\n", resTodos[i], item)
			}
		}

	}

}

func createSampleData(ds *DataSource, t *testing.T) (Todos, error) {
	todoList := []Todo{
		{Name: "Write Interface", Completed: true, Due: time.Date(2015, 4, 14, 0, 0, 0, 0, time.UTC)},
		{Name: "Write Tests", Completed: false, Due: time.Date(2015, 4, 14, 0, 0, 0, 0, time.UTC)},
		{Name: "Write Implementation", Completed: false, Due: time.Date(2015, 4, 14, 0, 0, 0, 0, time.UTC)},
		{Name: "Run test", Completed: false, Due: time.Date(2015, 4, 16, 0, 0, 0, 0, time.UTC)},
		{Name: "Get lunch", Completed: false, Due: time.Date(2015, 4, 16, 0, 0, 0, 0, time.UTC)},
	}

	var err error
	ds.ClearData()
	for _, todo := range todoList {
		resTodo, err := ds.CreateTodo(&todo)
		if err != nil {
			t.Errorf("Error by creating todo: %v\n", err)
			return todoList, err
		}
		if !todoEqualsExceptID(resTodo, &todo) {
			t.Errorf("Expected %v but result %v", todo, resTodo)
		}
	}
	return todoList, err
}

func testUpdate(ds *DataSource, test *testing.T) {

	// ds, err := NewDataSource()
	// if err != nil {
	// 	test.Errorf("Cannot create Todos datasource")
	// }

	//defer ds.Close()
	createSampleData(ds, test)
	dbTodos := ds.GetTodos()
	for _, oldTodo := range dbTodos {
		ID := oldTodo.ID
		oldTodo.Name += "+Mod"
		err := ds.UpdateTodo(ID, &oldTodo)
		if err != nil {
			test.Errorf("Cannot update %v", err)
		}
		resultTodo, err := ds.GetTodo(ID)
		if err != nil {
			test.Errorf("FAILED: Expected %v but not found any row", oldTodo)
		} else if !reflect.DeepEqual(resultTodo, oldTodo) {
			test.Errorf("FAILED: Expected %v but result %v", oldTodo, resultTodo)
		}
	}
}

func testDelete(ds *DataSource, test *testing.T) {
	// log.SetFlags(log.Lshortfile)
	// ds, err := NewDataSource()
	// if err != nil {
	// 	test.Errorf("Cannot create Todos datasource")
	// }
	// defer ds.Close()
	createSampleData(ds, test)
	dbTodos := ds.GetTodos()
	for i := 0; i < len(dbTodos); i++ {
		ID := dbTodos[i].ID
		err := ds.DeleteTodo(ID)
		if err != nil {
			test.Errorf("Cannot delete ID = %v, error : %v", ID, err)
		}
		list := ds.GetTodos()
		if !reflect.DeepEqual(dbTodos[i+1:], list) {
			test.Errorf("FAIL test: Expected %v but result %v", dbTodos[i+1:], list)
		}
	}

}

func todoEqualsExceptID(t *Todo, u *Todo) bool {
	return t.Name == u.Name && u.Due == t.Due && u.Completed == t.Completed
}
