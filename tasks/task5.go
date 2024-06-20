package tasks

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// тут у нас есть три способа завершения программы по времени
// 1) использовать канал, который нам дает time.After
// 2) в отдельной функции мы вызываем sleep и после os.Exit
// 1 решение гораздо лучше, потому что мы позволяем все из каналов достать, в 2 случае мы просто всю программу завершаем и не даем каналу опустошиться
// 3) Своя реализация способа 1. Как sleep закончится в канал попадет значение и все аккуратно завершится

func Sender1() {
	var n time.Duration
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	fmt.Scan(&n)

	to := time.After(n * time.Second)
	go Reader(ch, wg)
	go func() {
		for {
			select {
			case <-to:
				close(ch)
				return
			default:
				ch <- rand.Int() % 1000
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	wg.Wait()
}

func Ender1(duration time.Duration) {

	time.Sleep(duration * time.Second)

	os.Exit(0)
}

func Sender2() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	var duration time.Duration

	fmt.Scan(&duration)
	wg.Add(1)

	go Ender1(duration)
	go Reader(ch, wg)
	for {

		ch <- rand.Int() % 1000
	}
	wg.Wait()
}

func Ender3(duration time.Duration, done chan bool) {

	time.Sleep(duration * time.Second)

	done <- true
}

func Sender3() {
	ch := make(chan int)
	done := make(chan bool)
	wg := &sync.WaitGroup{}
	var duration time.Duration

	fmt.Scan(&duration)
	wg.Add(1)

	go Ender3(duration, done)
	go Reader(ch, wg)
	go func() {
		for {
			select {
			case <-done:
				close(ch)
				return
			default:
				ch <- rand.Int() % 1000
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	wg.Wait()
}

func Reader(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println(val)
	}
}
