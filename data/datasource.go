package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DataSource struct {
	db *sql.DB
}

func NewDataSource() (*DataSource, error) {
	db, err := sql.Open("mysql", "root:dangthaison@tcp(127.0.0.1:3306)/todos_app")
	return &DataSource{db}, err
}

//CRUD
//Create
func (ds *DataSource) CreateTodo(t *Todo) error {
	db := ds.db
	_, err := db.Exec("insert into todos (`name`,`completed`,`due`) values (?,?,?)", t.Name, t.Completed, t.Due)
	return err
}

//Read
func (ds *DataSource) GetTodos() Todos {
	rows, err := ds.db.Query("select * from todos")
	if err != nil {
		return nil
	}
	defer rows.Close()
	list := Todos{}
	for rows.Next() {
		t := Todo{}
		err := rows.Scan(&t.ID, &t.Name, &t.Completed, &t.Due)
		if err != nil {
			log.Fatal(err)
			continue
		}
		list = append(list, t)
	}
	return list
}

func (ds *DataSource) GetTodo(ID int) (*Todo, error) {
	return &Todo{}, nil
}

//Update
func (ds *DataSource) UpdateTodo(ID int, t *Todo) error {
	return nil
}

//Delete
func (ds *DataSource) DeleteTodo(ID int) error {
	return nil
}

func (ds *DataSource) ClearData() error {
	_, err := ds.db.Exec("delete from todos")
	return err
}
