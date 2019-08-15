package chapter5

import (
	"math/rand"
	"time"
)

func Shuffle(A []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := len(A)
	for i, _ := range A {
		j := r.Intn(len-i) + i
		A[i], A[j] = A[j], A[i]
	}
}
