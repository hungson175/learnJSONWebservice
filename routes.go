package main

import (
	"net/http"

	"github.com/hungson175/learnJsonWebservice/controllers"
	"github.com/hungson175/learnJsonWebservice/data"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{}

func init() {
	ds, err := data.NewDataSource()
	if err != nil {
		panic(err)
	}
	todoController := &controllers.TodoController{DataSource: ds}
	routes = Routes{Route{
		"Index",
		"GET",
		"/",
		todoController.Index},
		Route{
			"TodoIndex",
			"GET",
			"/todos",
			todoController.TodoIndex},
		Route{
			"TodoShow",
			"GET",
			"/todos/{todoID}",
			todoController.TodoShow,
		},
		Route{
			"TodoCreate",
			"POST",
			"/todos",
			todoController.TodoCreate,
		},
	}

}
