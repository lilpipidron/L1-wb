package tasks

// Источник информации: https://refactoring.guru/ru/design-patterns/adapter
// Базовая реализация паттерна адаптер
// Struct2 реализующая интерфейс Interface2, оборачивается в структуру Adapter
// которая, реализует Interface1, также как и структура Struct1

type Interface1 interface {
	DoSmth1(a, b int) int
}

type Struct1 struct{}

func (str *Struct1) DoSmth1(a, b int) int {
	return a + b
}

type Interface2 interface {
	DoSmth2(a, b float64) float64
}

type Struct2 struct{}

func (str *Struct2) DoSmth2(a, b float64) float64 {
	return a + b
}

type Adapter struct {
	str2 *Struct2
}

func (adapter *Adapter) DoSmth1(a, b int) int {
	result := adapter.str2.DoSmth2(float64(a), float64(b))
	return int(result)
}
