package main

import (
	"fmt"
	// "math"
)

func main() {
	var taskOne = "Watch Go crash course!"
	var taskTwo = "Watch full Go tutorial!"
	var taskItems = []string{taskOne, taskTwo}

	// fmt.Println("#### welcome to our Todolist App! ####")
	printTasks(taskItems)

	taskItems = addTask(taskItems, "Learn Go in depth!")
	taskItems = addTask(taskItems, "Build a Go project!")

	printTasks(taskItems)
}

func printTasks(taskItems []string) {
	fmt.Println("List my Todos")
	for index, task := range taskItems {
		// fmt.Println(index+1, ".",  task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	taskItems = append(taskItems, newTask)
	return taskItems
}
