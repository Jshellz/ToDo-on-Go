package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Completed   bool   `json:"Completed"`
}

var Todos = []Todo{
	{Name: "todo 1", Description: "this is todo", Completed: false},
	{Name: "todo 2", Description: "this is todo 2", Completed: true},
}
