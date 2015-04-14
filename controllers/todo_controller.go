package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/hungson175/learnJSONWebservice/data"
)

type TodoController struct {
	//ListTodos data.Todos
	DataSource *data.DataSource
}

func (ct *TodoController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to JSON Webservice")
}

func (ct *TodoController) TodoIndex(w http.ResponseWriter, r *http.Request) {
	setContentJson(&w)
	w.WriteHeader(http.StatusOK)
	list := ct.DataSource.GetTodos()
	if err := json.NewEncoder(w).Encode(list); err != nil {
		panic(err)
	}
}

func (ct *TodoController) TodoShow(w http.ResponseWriter, r *http.Request) {
	setContentJson(&w)
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["todoID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	todo, err := ct.DataSource.GetTodo(todoID)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func (ct *TodoController) TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo data.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		setContentJson(&w)
		w.WriteHeader(422) //unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	createdTodo, _ := ct.DataSource.CreateTodo(&todo)
	setContentJson(&w)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdTodo); err != nil {
		panic(err)
	}

}

func setContentJson(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
}
