package tasks

import (
	"strconv"
)

//В go нет явной имплементации интерфейса
//структура имплементирует интерфейс, если реализует все методы интерфейса
//также в go нет явного наследования, по сути встраивание методов происходит
//когда в дочерней структуре есть поле, которое является родительской структурой
//как показано ниже. В таком случае мы можем вызывать методы родительской структуру из дочерней

type Human struct {
	Name string
	Age  int
}

type CoolInterface interface {
	GetNameAndAge() string
}

func (h Human) GetNameAndAge() string {
	return h.Name + strconv.Itoa(h.Age)
}

type Action struct {
	Human
}
