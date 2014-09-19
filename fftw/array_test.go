package fftw

import (
	"testing"
)

func TestArray2_Slice(t *testing.T) {
	m, n := 5, 8
	arr := NewArray2(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr.Set(i, j, complex(float64(i), float64(j)))
		}
	}
	for i, si := range arr.Slice() {
		for j := range si {
			want := complex(float64(i), float64(j))
			if si[j] != want {
				t.Errorf("at (%d, %d): want %v, got %v", i, j, si[j], want)
			}
		}
	}
}

func TestArray3_Slice(t *testing.T) {
	m, n, p := 5, 8, 6
	arr := NewArray3(m, n, p)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < p; k++ {
				// Fill with something non-commutative.
				arr.Set(i, j, k, complex(float64(((i+1)*j+1)*k), 0))
			}
		}
	}
	for i, si := range arr.Slice() {
		for j := range si {
			for k := range si[j] {
				want := complex(float64(((i+1)*j+1)*k), 0)
				if si[j][k] != want {
					t.Errorf("at (%d, %d, %d): want %v, got %v", i, j, k, si[j][k], want)
				}
			}
		}
	}
}
