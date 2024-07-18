package vrandx

const MaxUint32 = 4294967295

// yet another simple pseudo-random number
type SplitMix32 struct {
	State uint32 // need set the state as first time seed
}

func (s *SplitMix32) Random() uint32 {
	s.State += 0x9e3779b9
	z := s.State
	z = (z ^ (z >> 16)) * 0x21f0aaad
	z = (z ^ (z >> 15)) * 0x735a2d97
	return z ^ (z >> 15)
}

func (s *SplitMix32) RandomFloat() float64 {
	return float64(s.Random()) / float64(MaxUint32)
}
