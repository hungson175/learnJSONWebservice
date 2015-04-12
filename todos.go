package main

import "time"

type Todo struct {
	ID        int       `json:"ID"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
