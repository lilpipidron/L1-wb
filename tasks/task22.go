package tasks

import (
	"fmt"
	"math/big"
)

// Все решение основано на пакете math/big и использует одноименные функции для операций
// пакет нужен специально для таких больших числе

func Solve22() {
	var a, b big.Int
	mulRes := new(big.Int)
	subRes := new(big.Int)
	divRes := new(big.Int)
	addRes := new(big.Int)

	fmt.Scan(&a, &b)

	fmt.Println(mulRes.Mul(&a, &b))
	fmt.Println(divRes.Div(&a, &b))
	fmt.Println(addRes.Add(&a, &b))
	fmt.Println(subRes.Sub(&a, &b))
}
