package tasks

import "fmt"

// Я так понимаю, тут хотят увидеть сет
// В сете у нас все элементы уникальны, просто сделаем мапу, в которой ключом будет строка, а значение - пустая структура для экономии памяти

func Solve12() {
	elems := []string{"cat", "cat", "dog", "cat", "tree"}
	mapOfElems := make(map[string]struct{})

	for _, elem := range elems {
		mapOfElems[elem] = *new(struct{})
	}

	var ans []string
	for key := range mapOfElems {
		ans = append(ans, key)
	}

	fmt.Println(ans)
}
