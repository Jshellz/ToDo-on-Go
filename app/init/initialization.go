package init

import (
	"ToDo/app/db"
	"ToDo/app/model"
)

func init() {
	db.ConnectToDB()
	db.CreateModel(&model.Todo{
		Name: "todo 1", Description: "this is todo", Completed: false,
		//Name: "todo 2", Description: "this is todo 2", Completed: true,
	})
}
