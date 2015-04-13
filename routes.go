package main

import (
	"net/http"

	"github.com/hungson175/learnJSONWebservice/controllers"
	"github.com/hungson175/learnJSONWebservice/data"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var todoController = &controllers.TodoController{ListTodos: data.RepoListTodos()}
var routes = Routes{
	Route{
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
