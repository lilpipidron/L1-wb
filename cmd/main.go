package main

import (
	"fmt"
	"github.com/lilpipidron/L1-wb/tasks"
)

func main() {
	//task1
	fmt.Println("task1")
	human := tasks.Human{Name: "John", Age: 25}
	action := tasks.Action{Human: human}
	fmt.Println(action.GetNameAndAge())

	//task2
	fmt.Println("task2")
	tasks.Solve2()

	//task 3
	fmt.Println("task3")
	tasks.Solve3()

	//task 4
	fmt.Println("task4")
	tasks.Solve4()
}
