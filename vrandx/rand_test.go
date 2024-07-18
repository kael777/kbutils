package vrandx

import (
	"math/rand"
	randV2 "math/rand/v2"
	"testing"
)

const Max = 1e9

/*

test results
goos: windows
goarch: amd64
pkg: test/randx
cpu: Intel(R) Core(TM) i5-7400 CPU @ 3.00GHz
BenchmarkRand-4       	62930680	        19.06 ns/op	           0 B/op	       0 allocs/op
BenchmarkRandV2-4     	129878239	         9.106 ns/op	       0 B/op	       0 allocs/op
BenchmarkFastRand-4   	158387926	         8.008 ns/op	       0 B/op	       0 allocs/op
BenchmarkMt19937-4    	180471915	         7.188 ns/op	       0 B/op	       0 allocs/op
BenchmarkMt00001-4    	535954507	         2.195 ns/op	       0 B/op	       0 allocs/op
PASS

*/

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(Max)
	}
}

func BenchmarkRandV2(b *testing.B) {
	// rand v2 after golang 1.22
	for i := 0; i < b.N; i++ {
		randV2.IntN(Max)
	}
}

func BenchmarkFastRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32n(Max)
	}
}

func BenchmarkMt19937(b *testing.B) {
	sp := NewMT19937()
	for i := 0; i < b.N; i++ {
		sp.Int63()
	}
}

func BenchmarkMt00001(b *testing.B) {
	sp := &SplitMix32{State: 11}
	for i := 0; i < b.N; i++ {
		sp.RandomFloat()
	}
}
