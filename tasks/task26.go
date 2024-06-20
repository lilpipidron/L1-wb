package tasks

import "strings"

// решение строится на мапе. В первую очередь, чтобы функция была регистронезависимой мы приводим все символы к одному регистру
// далее проосто идем по всем символам и если этот символ уже есть в мапе, то возвращаем false, иначе, если ни разу не зашли в тело if'а
// то просто вернем true

func Solve26(str string) bool {
	symbols := make(map[rune]int, 26)
	str = strings.ToLower(str)
	for _, sym := range str {
		if _, ok := symbols[sym]; ok {
			return false
		}
		symbols[sym]++
	}
	return true
}
