package main

import (
	"net/http"
	tasks "go-demo1/internal/tasks"

	dict  "go-demo1/internal/dictionary"
)

func main() {
   app := http.NewServeMux()

   _ = tasks.NewTaskHandler(app)
    _ = dict.NewDictionary(app)
   server := http.Server{
	Addr: ":8080",
	Handler: app,
   }

   server.ListenAndServe()
}

