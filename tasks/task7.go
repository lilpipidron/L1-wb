package tasks

import (
	"fmt"
	"sync"
)

// Есть 2 возможных выхода.
// 1) Своя собственная структура с RWMutex для чтения сразу несколькими горутинами при RLock при
// отсутсвиив в данный момент Lock
// 2) Исользуем sync.Map, которая позволяет сделать тоже самое, только все уже под копотом

type ConcurrencyMap struct {
	mx sync.RWMutex
	m  map[string]int
}

func (c *ConcurrencyMap) Init() {
	c.m = make(map[string]int)
}

func (c *ConcurrencyMap) Add(key string, value int) {
	c.mx.Lock()
	c.m[key] = value
	c.mx.Unlock()
}

func (c *ConcurrencyMap) Get(key string) int {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.m[key]
}

func SolveWithSyncMap() {
	var m sync.Map
	m.Store("a", 1)
	fmt.Println(m.Load("a"))
}
