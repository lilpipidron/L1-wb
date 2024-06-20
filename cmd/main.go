package main

import (
	"fmt"
	"github.com/lilpipidron/L1-wb/tasks"
)

func main() {
	human := tasks.Human{Name: "John", Age: 25}
	action := tasks.Action{Human: human}
	fmt.Println(action.GetNameAndAge())
}
