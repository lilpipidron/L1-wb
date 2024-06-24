package tasks

import (
	"fmt"
	"sort"
)

// Делаем структурку где первое поле это десяток градусов, а второе само множество температур
// Далее с помощью мапы массив конвертируем в набор структур
// После переносим это в массив таких стрктур и сортируем, чтобы множества шли в порядке возрастания их "заголовка"

type Set struct {
	Label int
	Arr   []float64
}

func NewSet(label int) Set {
	return Set{Label: label, Arr: *new([]float64)}
}

func Solve10() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sets := make(map[int]Set)
	for _, val := range arr {
		label := int(val) / 10 * 10
		if _, ok := sets[label]; !ok {
			sets[label] = NewSet(label)
		}
		set := sets[label]
		set.Arr = append(set.Arr, val)
		sets[label] = set
	}
	var ans []Set
	for _, val := range sets {
		ans = append(ans, val)
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Label < ans[j].Label
	})

	fmt.Println(ans)
}
