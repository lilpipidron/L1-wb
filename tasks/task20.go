package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Изначально мы считываем всю строку через ReadString, читаем до символа конца строки
// Далее убираем этот символ из строки и разбиваем ее по пробела
// Идем в обратном порядке по массиву слов и прикрепляем их к строке-ответу, не забывая про пробелы

func Solve20() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimSuffix(str, "\n")

	words := strings.Split(str, " ")

	ans := ""
	for i := len(words) - 1; i >= 0; i-- {
		ans += words[i]
		if i > 0 {
			ans += " "
		}
	}

	fmt.Println(ans)
}
