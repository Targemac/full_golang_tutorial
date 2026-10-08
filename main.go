package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("*** Welcome to my Todo List ***")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	taskItems := []string{"Buy groceries", "Clean the house", "Finish homework"}
	for _, task := range taskItems {
		fmt.Fprintln(w, task)
	}
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	var greeting string
	greeting = "Hello User! Welcome to my Todo List"
	fmt.Fprintln(w, greeting)
}
