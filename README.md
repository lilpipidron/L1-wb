# L1-wb

## Устные задачи

### Задача 1
Какой самый эффективный способ конкатенации строк?

Использование ```strings.Builder()``` 
### Задача 2
Что такое интерфейсы, как они применяются в Go?

Интерфейс нужен, чтобы описать методы, которые должны быть у какого-то типа 
Если структура реализует все методы интерфейса, то она автоматически наследуется от него

### Задача 3
Чем отличаются RWMutex от Mutex?

Mutex блокирует полностью ресурс с помощью Lock()
RWMutex имеет RLock() для того, чтобы много горутин одновременно могли читать разделяемый ресурс
и Lock() для полного блокирования на запись

### Задача 4
Чем отличаются буферизированные и не буферизированные каналы?

не буферизированный канал имеет вместимость 1, и если там уже что-то лежит, то поток записи в него встанет в ожидание
буферизированный канал имеет вместимость n, про ожидание аналогично
### Задача 5

Какой размер у структуры struct{}{}?
Ответ: 0

### Задача 6

Есть ли в Go перегрузка методов или операторов?
Ответ: нет

### Задача 7

В какой последовательности будут выведены элементы map[int]int?

Пример:
m[0]=1
m[1]=124
m[2]=281

Ответ: четкого порядка нет, все зависит от хэширования

### Задача 8
В чем разница make и new?
new используется для выделения памяти и возвращает указатель на тип, инициализируя его нулевым значением.
make используется для создания и инициализации срезов, карт и каналов, возвращая инициализированное значение соответствующего типа, готовое к использованию.


### Задача 9
Сколько существует способов задать переменную типа slice или map?
```go
    var s []int
    var m map[int]int
```
```go
    s := make([]int, 0)
    m := make(map[int]int)
```
```go
    s := new([]int, 0)
    m := new(map[int]int)
```
```go
    s := []int{}
    m:= map[int]int{}
```

### Задача 10

Что выведет данная программа и почему?

```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```
Ответ: 
1, 1
Первая единица, потому что p указывает на a, которая ранва единице
Вторая единица, потому что мы не изменили значение p, которое в мейне, надо было передать указатель на указатель, вот так:
```go
func update(p **int) {
	b := 2
	*p = &b
}
func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(&p)
	fmt.Println(*p)
}
```

### Задача 11
Что выведет данная программа и почему?
```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```

Вывод: от 0 до 5 в случайном порядке и дедлок, потому что мы wg передаем не по ссылке и когда вызываем wg.Done() в горутине, значение в основном потоке не изменяется
корректная реализация:
```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg *sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(&wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```

### Задача 12
Что выведет данная программа и почему?
```go
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
Ответ: 0, потому что изменяем новую переменную, которую создали в области видимости if
Пример, когда будет 2 на выходе
```go
func main() {
	n := 0
	if true {
		n = 1
		n++
	}
	fmt.Println(n)
}
```

### Задача 13
Что выведет данная программа и почему?

```go
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
Ответ: [100 2 3 4 5], 100 будет вместо 1, потому что массив передается по ссылке, и в функции мы изменяем исходный массив. 
6 не будет, потому что append() возвращает новый массив 

### Задача 14
Что выведет данная программа и почему?
```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
Ответ: [b b a] [a a] аналогично заданию 13