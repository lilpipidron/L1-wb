package tasks

import (
	"context"
	"fmt"
	"time"
)

// способ 1 подходит для завершения только 1 горутины, первая которая из канала вычитает, та и завершится
// способ 2 уже для любого количества горутин подходит, как только у нас закроется канал, цикл завершится, а за ним горутина
// способ 3 основан на глобальной перемене
// способ 4 контекст с таймаутом
// способ 5 контекст с дедлайном

func Solve6_1() {
	done := make(chan bool)

	go WayToClose1(done)

	time.Sleep(1 * time.Second)
	done <- true
}

func WayToClose1(done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println(1)
		}
	}
}

func Solve6_2() {
	done := make(chan bool)

	go WayToClose2(done)

	time.Sleep(1 * time.Second)
	close(done)
}

func WayToClose2(done chan bool) {
	for <-done {
		fmt.Println(1)
	}
}

var working bool

func Solve6_3() {
	go WayToClose3()
	working = true
	time.Sleep(1 * time.Second)
	working = false
}

func WayToClose3() {
	for working {
		fmt.Println(1)
	}
}

func Solve6_4() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go WayToClose4(ctx)
	<-ctx.Done()
}

func WayToClose4(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(1)
		}
	}
}

func Solve6_5() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel()
	go WayToClose5(ctx)
	<-ctx.Done()
}

func WayToClose5(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(1)
		}
	}
}
