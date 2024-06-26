package tasks

import "fmt"

// Полученную строку преобразвываем в массив рун, потому что у нас может использоваться символ Юникода
// обычный символ латиницы занимает 1 байт, в таком случае можно было сделать массив byte, но в нашем случае
// это привело бы к некорректной работе программы, потому что символы кириллицы занимают 2 байта
// поэтому в данном случае работа с рунами лучше, потому что каждый элемент массива будет представлять один символ Юникода,
// независимо от того, сколько байтов он занимает

func Solve19() {
	var str, ans string
	fmt.Scan(&str)

	symbols := []rune(str)
	for i := len(symbols) - 1; i >= 0; i-- {
		ans += string(symbols[i])
	}

	fmt.Println(ans)
}
