package tasks

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Имеем 2 канала: для записи и для отлова сигнала ctrl+c
// также имеем wg, чтобы наши воркеры достали все числа из канала
// и только после этого мы завершаем окончательно программу
// также мы закрываем канал сразу после сигнала, чтобы больше не писать и
// наши воркеры понимали, что в канал больше ничего не будет записано
// и как только он опустеет они завершатся и defer wg.Done() будет выполнено

func Solve4() {
	ch := make(chan int)
	var numWorkers int
	wg := &sync.WaitGroup{}

	fmt.Scan(&numWorkers)
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go Worker(ch, wg)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case ch <- rand.Int() % 1000:
		case <-sigChan:
			close(ch)
			wg.Wait()
			return
		}
	}
}

func Worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println(val)
	}
}
