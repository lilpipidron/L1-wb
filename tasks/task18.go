package tasks

import "sync"

// Горутины будут вызывать функцию Increment, counter.value лочится, чтобы сонхронизировать доступ
// В конце вызываем функцию Close, которая вернет нам нынешний счетчик, опять же лочим, чтобы другая горутина
// не изменила счетчик в неподходящий момент

type Counter struct {
	mx    sync.Mutex
	value int
}

func (counter *Counter) Increment() {
	defer counter.mx.Unlock()
	counter.mx.Lock()
	counter.value++
}

func (counter *Counter) Close() int {
	defer counter.mx.Unlock()
	counter.mx.Lock()
	return counter.value
}
