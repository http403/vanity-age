package main

import (
	"fmt"
	"testing"

	"filippo.io/age"
)

func doubleReduceAdd(r *[]int, n int, p int) {
	if n <= 1 {
		if len(*r) == 0 {
			*r = append(*r, n)
		}
		return
	}

	p = n + p
	*r = append(*r, p)
	doubleReduceAdd(r, n/2, p)
}

func BenchmarkGenerate(b *testing.B) {
	benchThreads := []int{}
	doubleReduceAdd(&benchThreads, THREADS, 0)

	for _, v := range benchThreads {
		b.Run(fmt.Sprintf("thread_count_%d", v),
			func(b *testing.B) {
				keyChan := make(chan *age.X25519Identity)
				for i := 0; i < v; i++ {
					go generate("age1some*", keyChan)
				}
			})
	}
}
