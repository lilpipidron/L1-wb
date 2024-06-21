package tasks

import "fmt"

// Решение, собственным множеством является множетсво, которое не является самим множеством или пустым
// в решении возьмем и получим все уникальные элементы, а после уберем один из них в исходном множестве
// и получим собственное множество
// не сработает, когда у нас всего 1 элемент, так как мы не сможем получить не пустое множество или множество не совпадающее
// с исходным

func Solve12() {
	elems := []string{"cat", "cat", "dog", "cat", "tree"}
	mapOfAmount := make(map[string]int)

	for _, elem := range elems {
		mapOfAmount[elem]++
	}

	var ans []string
	count := 0
	for key := range mapOfAmount {
		if count == 0 {
			count++
			continue
		}
		ans = append(ans, key)
	}

	fmt.Println(ans)
}
