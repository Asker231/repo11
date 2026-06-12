package main

import (
	"net/http"
	tasks "go-demo1/internal/tasks"
)

func main() {
   app := http.NewServeMux()

   _ = tasks.NewTaskHandler(app)

   server := http.Server{
	Addr: ":8080",
	Handler: app,
   }

   server.ListenAndServe()
}

