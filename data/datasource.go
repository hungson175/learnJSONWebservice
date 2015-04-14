package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DataSource struct {
	db *sql.DB
}

type NotImplementedError struct {
	name string
}

func (err NotImplementedError) Error() string {
	return err.name + ": Not implemeted yet"
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
		dateString := ""
		err := rows.Scan(&t.ID, &t.Name, &t.Completed, &dateString)
		if err != nil {
			log.Fatal(err)
			continue
		}
		t.Due, _ = time.Parse("2006-01-02", dateString)
		list = append(list, t)
	}
	return list
}

func (ds *DataSource) GetTodo(ID int) (Todo, error) {
	t := Todo{}
	var dateString string
	err := ds.db.QueryRow("select * from todos where id = ?", ID).Scan(&t.ID, &t.Name, &t.Completed, &dateString)
	if err != nil {
		return Todo{}, err
	}
	t.Due, err = time.Parse("2006-01-02", dateString)
	return t, err
}

//UpdateTodo
//Update the content of the Todo with id = ID in DB with content of t (except for the ID)
func (ds *DataSource) UpdateTodo(ID int, t *Todo) error {
	dateString := fmt.Sprintf("%d-%d-%d", t.Due.Year(), t.Due.Month(), t.Due.Day())
	_, err := ds.db.Exec("update todos SET name=?, completed = ?, due = ? where ID = ?", t.Name, t.Completed, dateString, ID)
	return err
}

func (ds *DataSource) DeleteTodo(ID int) error {
	_, err := ds.db.Exec("delete from todos where id = ?", ID)
	return err
}

func (ds *DataSource) ClearData() error {
	_, err := ds.db.Exec("delete from todos")
	return err
}

func (ds *DataSource) Close() error {
	return ds.db.Close()
}
