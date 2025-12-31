package fftw

// Tests that functions can be called.

import "testing"

func TestFFT_exec(t *testing.T) {
	t.Parallel()

	n := 32
	x := NewArray(n)
	x = FFT(x)
	_ = IFFT(x)
}

func TestFFT2_exec(t *testing.T) {
	t.Parallel()

	n0, n1 := 32, 16
	x := NewArray2(n0, n1)
	x = FFT2(x)
	_ = IFFT2(x)
}

func TestFFT3_exec(t *testing.T) {
	t.Parallel()

	n0, n1, n2 := 32, 16, 8
	x := NewArray3(n0, n1, n2)
	x = FFT3(x)
	_ = IFFT3(x)
}

func TestFFTN_exec(t *testing.T) {
	t.Parallel()

	n := []int{32, 16, 8, 4}
	x := NewArrayN(n)
	x = FFTN(x)
	_ = IFFTN(x)
}
