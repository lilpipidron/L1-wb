package tasks

import (
	"fmt"
	"strconv"
)

// имеем 2 решения
// 1) решение основано на битовых операциях, если мы хотим бит заменить на еденицу, то используем или
// потому что есть 2 случая бит = 1 или бит = 0, и если бы мы сделали &&, когда бит 0, то также бы остался 0
// и аналогично, когда хотим заменить на 0, если бы была 1, и был оператор ||, мы бы получчили также 1, а не 0
// 2) решение основано на строковых функциях. Делаем из числа строку из 1 и 0 с помощью функции FormatInt
// и преобразовываем это в срез рун, потому что строки в го неизменяемы, меняем нуный бит 1 имеет руну 49, а 0 руну 48
// и возвращаем обратно в int64 с помощью ParseInt

func Solve8_1() {
	var num int64
	var bit int
	var val uint8
	fmt.Scan(&num, &bit, &val)

	if val == 1 {
		num |= 1 << bit
	} else {
		num &= ^(1 << bit)
	}

	fmt.Println(num)
}

func Solve8_2() {
	var num int64
	var bit int
	var val int
	fmt.Scan(&num, &bit, &val)

	str := []rune(strconv.FormatInt(num, 2))
	if val == 0 {
		str[len(str)-bit-1] = 48
	} else {
		str[len(str)-bit-1] = 49
	}
	ans, _ := strconv.ParseInt(string(str), 2, 64)
	fmt.Println(ans)
}
