package tasks

import (
	"fmt"
	"time"
)

// Все решения делают sleep в секундах для удобства
// 1) Решение основано на цикле, пока текующее время - время начала цикла меньше задержки, то цикл будет крутиться
// 2) Решение основано на канале, как только необходимое время пройдет, то он отпустит поток и мы сможем идти дальше
// 3) Решение аналогично 2, но тут мы используем timer

func Solve25_1() {
	var seconds time.Duration
	fmt.Scan(&seconds)

	t := time.Now()
	for time.Since(t) < seconds*time.Second {
	}

	fmt.Println("Solve25_1 complete")
}

func Solve25_2() {
	var seconds time.Duration
	fmt.Scan(&seconds)
	<-time.After(seconds * time.Second)

	fmt.Println("Solve25_2 complete")
}

func Solve25_3() {
	var seconds time.Duration
	fmt.Scan(&seconds)
	t := time.NewTimer(seconds * time.Second)
	<-t.C

	fmt.Println("Solve25_3 complete")
}
