package utils

import (
	"github.com/IgorCastilhos/toolkit/v2"
	"math/rand"
	"time"
)

var tools toolkit.Tools

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt gera um inteiro aleatório entre o min e o max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0->max-min
}

func RandomOwner() string {
	length := 10
	return tools.RandomString(length)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"BRL"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
