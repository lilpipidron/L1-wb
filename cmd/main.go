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

	/*	//task8
		fmt.Println("task8")
		tasks.Solve8_1()
		tasks.Solve8_2()*/

	//task9
	fmt.Println("task9")
	tasks.Solve9()

	//task11
	fmt.Println("task11")
	tasks.Solve11_1()
	tasks.Solve11_2()
	tasks.Solve11_3()

	//task12
	fmt.Println("task12")
	tasks.Solve12()

	//task13
	fmt.Println("task13")
	tasks.Solve13_1()
	tasks.Solve13_2()
	tasks.Solve13_3()

	//task19
	fmt.Println("task19")
	tasks.Solve19()

	//task20
	fmt.Println("task20")
	tasks.Solve20()

	//task22
	fmt.Println("tasks22")
	tasks.Solve22()

	//task24
	fmt.Println("task24")
	point1 := tasks.NewPoint(1, 1)
	point2 := tasks.NewPoint(2, 2)
	fmt.Println(tasks.Solve24(point1, point2))

	//task26
	fmt.Println("task26")
	fmt.Println(tasks.Solve26("abc"))
	fmt.Println(tasks.Solve26("aAbBcC"))
}
