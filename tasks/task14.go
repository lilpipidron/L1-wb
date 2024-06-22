package tasks

import (
	"fmt"
	"reflect"
)

// 1) Решение пострено на рефлексии из стандартного пакета go reflect
// 2) Решение построено на утверждение типа (type assertion) для определения типа переменной.
// 3) Решение пострено на рефлексии из стандартного пакета go reflect

func Solve14_1(smth interface{}) {
	fmt.Println(reflect.TypeOf(smth))
}

func Solve14_2(smth interface{}) {
	switch smth.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int, chan string, chan bool, chan interface{}:
		fmt.Println("chan")
	}
}

func Solve14_3(smth interface{}) {
	switch reflect.ValueOf(smth).Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("chan")
	default:
		fmt.Println("unknown type")
	}
}
