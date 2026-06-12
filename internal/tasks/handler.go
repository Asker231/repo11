package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Tasks struct {
	app *http.ServeMux
}

func NewTaskHandler(app *http.ServeMux) *Tasks{
	tasks := &Tasks{
		app: app,
	}
	app.HandleFunc("GET /tasks",tasks.getTasks())
	return  tasks
}

func(t *Tasks)getTasks() http.HandlerFunc{
	var currency Wrapper
	file,err := os.Open("internal/tasks/task.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&currency)
	if err != nil{
		fmt.Println(err.Error())
	}
	return  func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(currency)
		if err != nil{
			fmt.Println(err.Error())
		}
	}
}