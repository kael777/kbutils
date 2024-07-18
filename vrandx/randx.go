package vrandx

// a faster rand int function,use go linkname to use runtime.fastrand function
// it run 30% faster than rand.Intn
// also , rand.Float64 is faster than rand.Intn > rand.Int64n > rand.Int32n

import (
	_ "unsafe"
)

type Seed struct {
	s uint64
}

//go:linkname Uint32 runtime.fastrand
func Uint32() uint32

//go:linkname Uint32n runtime.fastrandn
func Uint32n(max uint32) uint32

// MakeSeed returns a new random seed.
func MakeSeed() Seed {
	var s1, s2 uint64
	for {
		s1 = uint64(Uint32())
		s2 = uint64(Uint32())
		// We use seed 0 to indicate an uninitialized seed/hash,
		// so keep trying until we get a non-zero seed.
		if s1|s2 != 0 {
			break
		}
	}
	return Seed{s: s1<<32 + s2}
}
