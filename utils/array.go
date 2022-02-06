package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomArray(n int) []int {
	return rand.Perm(n)
}

func RandomBetween(min, max int) int {
	return rand.Intn(max-min) + min
}
