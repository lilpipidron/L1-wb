package tasks

import (
	"fmt"
	"sync"
)

// Решение основывается на том, что второй канал у нас будет буфферезированный. Мы знаем сколько в итоге в нем будет значений
// И это помогает нам избежать deadlock на строке с wg.Wait(), потому что при записи в не буфферезированный канал у нас 2 горутина
// будет ждать пока из него прочтут что-то, но этого не будет, так как чтение будет после завершения всех горутин
// и первая тоже будет ждать пока из 1 канала 2 горутина что-то прочтет и мы никогда не выйдем из этого состояния

func Solve9() {
	var n int
	ch1 := make(chan int)
	wg := sync.WaitGroup{}

	fmt.Scan(&n)
	arr := make([]int, n)
	ch2 := make(chan int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, val := range arr {
			ch1 <- val
		}
		close(ch1)
	}()

	go func() {
		defer wg.Done()
		for val := range ch1 {
			ch2 <- val * val
		}
		close(ch2)
	}()

	wg.Wait()

	for val := range ch2 {
		fmt.Println(val)
	}
}
