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

	/*//task 4
	fmt.Println("task4")
	tasks.Solve4()

	//task 5
	fmt.Println("task5")
	tasks.Sender1()
	tasks.Sender2()
	tasks.Sender3()*/

	/*	//task 6
		fmt.Println("task6")
		tasks.Solve6_1()
		tasks.Solve6_2()
		tasks.Solve6_3()
		tasks.Solve6_4()
		tasks.Solve6_5()*/

	//task26
	fmt.Println("task26")
	fmt.Println(tasks.Solve26("abc"))
	fmt.Println(tasks.Solve26("aAbBcC"))
}
