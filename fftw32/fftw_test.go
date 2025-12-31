package fftw32

import (
	"math"
	"testing"
)

func TestNewArray(t *testing.T) {
	t.Parallel()

	for _, n := range []int{10, 100, 1000} {
		d := NewArray(n)
		if len(d.Elems) != n {
			t.Errorf("Expected %d elements got %d", n, len(d.Elems))
		}
	}
}

// Make sure that the memory allocated by fftw is getting properly GCed.
func TestGC(t *testing.T) {
	t.Parallel()

	var tot float32 = 0.0

	for i := range 1000 {
		d := NewArray(1000000)                  // Allocate a bunch of memory
		d.Elems[10000] = complex(float32(i), 0) // Do something stupid with it so
		tot += real(d.Elems[10000])             // hopefully it doesn't get optimized out
	}
}

func TestNewArray2(t *testing.T) {
	t.Parallel()

	d100x50 := NewArray2(100, 50)

	dim0, dim1 := d100x50.Dims()
	if dim0 != 100 {
		t.Fatalf("Expected dim[0] = 100, got %d", dim0)
	}

	if dim1 != 50 {
		t.Fatalf("Expected dim[1] = 50, got %d", dim1)
	}

	setArray2(d100x50, dim0, dim1)
	verifyArray2(t, d100x50, dim0, dim1)
}

func TestNewArray3(t *testing.T) {
	t.Parallel()

	d100x20x10 := NewArray3(100, 20, 10)

	dim0, dim1, dim2 := d100x20x10.Dims()
	if dim0 != 100 {
		t.Fatalf("Expected dim[0] = 100, got %d", dim0)
	}

	if dim1 != 20 {
		t.Fatalf("Expected dim[1] = 20, got %d", dim1)
	}

	if dim2 != 10 {
		t.Fatalf("Expected dim[2] = 10, got %d", dim2)
	}

	setArray3(d100x20x10, dim0, dim1, dim2)
	verifyArray3(t, d100x20x10, dim0, dim1, dim2)
}

func peakVerifier(t *testing.T, s []complex64) {
	t.Helper()
	testAlmostEqual(t, real(s[0]), 0.0)
	testAlmostEqual(t, imag(s[0]), 0.0)
	testAlmostEqual(t, real(s[1]), float32(len(s))/2)
	testAlmostEqual(t, imag(s[1]), 0.0)

	for i := 2; i < len(s)-1; i++ {
		testAlmostEqual(t, real(s[i]), 0.0)
		testAlmostEqual(t, imag(s[i]), 0.0)
	}

	testAlmostEqual(t, real(s[len(s)-1]), float32(len(s))/2)
	testAlmostEqual(t, imag(s[len(s)-1]), 0.0)
}

func TestFFT(t *testing.T) {
	t.Parallel()

	signal := NewArray(16)
	newIn := NewArray(16)

	for i := range signal.Elems {
		signal.Elems[i] = complex(float32(i), float32(-i))
		newIn.Elems[i] = signal.Elems[i]
	}

	// A simple real cosine should result in transform with two spikes, one at S[1] and one at S[-1]
	// The spikes should be real and have amplitude equal to len(S)/2 (because fftw doesn't normalize)
	for i := range signal.Elems {
		signal.Elems[i] = complex(float32(math.Cos(float64(i)/float64(len(signal.Elems))*math.Pi*2)), 0)
		newIn.Elems[i] = signal.Elems[i]
	}

	NewPlan(signal, signal, Forward, Estimate).Execute().Destroy()
	peakVerifier(t, signal.Elems)
}

func TestFFT2(t *testing.T) {
	t.Parallel()

	signal := NewArray2(64, 8)

	dim0, dim1 := signal.Dims()
	for i := range dim0 {
		for j := range dim1 {
			signal.Set(i, j, complex(float32(i+j), float32(-i-j)))
		}
	}

	// As long as freqX < lenX/2 and freqY < lenY/2, where lenX and lenY are the lengths in each dimension,
	// there will be 2^n spikes, where n is the number of dimensions.  Each spike will be
	// real and have magnitude equal to lenX*lenY / 2^n
	lenX := dim0
	freqX := float64(lenX) / 4
	lenY := dim1
	freqY := float64(lenY) / 4

	for i := range dim0 {
		for j := range dim1 {
			cosx := math.Cos(float64(i) / float64(lenX) * freqX * math.Pi * 2)
			cosy := math.Cos(float64(j) / float64(lenY) * freqY * math.Pi * 2)
			signal.Set(i, j, complex(float32(cosx*cosy), 0))
		}
	}

	NewPlan2(signal, signal, Forward, Estimate).Execute().Destroy()

	verifyFFT2(t, signal, dim0, dim1, freqX, freqY)
}

func TestFFT3(t *testing.T) {
	t.Parallel()

	signal := NewArray3(32, 16, 8)

	dim0, dim1, dim2 := signal.Dims()
	for i := range dim0 {
		for j := range dim1 {
			for k := range dim2 {
				signal.Set(i, j, k, complex(float32(i+j+k), float32(-i-j-k)))
			}
		}
	}

	// As long as freqX < lenX/2, freqY < lenY/2, and freqZ < lenZ/2, where lenX,lenY,lenZ  are the lengths in
	// each dimension, there will be 2^n spikes, where n is the number of dimensions.
	// Each spike will be real and have magnitude equal to lenX*lenY*lenZ / 2^n
	lenX := dim0
	freqX := float64(lenX) / 4
	lenY := dim1
	freqY := float64(lenY) / 4
	lenZ := dim2
	freqZ := float64(lenZ) / 4

	for i := range dim0 {
		for j := range dim1 {
			for k := range dim2 {
				cosx := math.Cos(float64(i) / float64(lenX) * freqX * math.Pi * 2)
				cosy := math.Cos(float64(j) / float64(lenY) * freqY * math.Pi * 2)
				cosz := math.Cos(float64(k) / float64(lenZ) * freqZ * math.Pi * 2)
				signal.Set(i, j, k, complex(float32(cosx*cosy*cosz), 0))
			}
		}
	}

	NewPlan3(signal, signal, Forward, Estimate).Execute().Destroy()

	verifyFFT3(t, signal, dim0, dim1, dim2, freqX, freqY, freqZ)
}

const almostEqualEpsilon = 0.000001

func almostEqual(v1, v2 float32) bool {
	return math.Abs(float64(v1-v2)) < almostEqualEpsilon
}

func testAlmostEqual(t *testing.T, v1, v2 float32) {
	t.Helper()

	if !almostEqual(v1, v2) {
		t.Fatalf("%f != %f (delta %f)", v1, v2, math.Abs(float64(v1-v2)))
	}
}

func setArray2(a *Array2, dim0, dim1 int) {
	var counter float32

	for i := range dim0 {
		for j := range dim1 {
			a.Set(i, j, complex(counter, 0))

			counter += 1.0
		}
	}
}

func verifyArray2(t *testing.T, a *Array2, dim0, dim1 int) {
	t.Helper()

	var counter float32

	for i := range dim0 {
		for j := range dim1 {
			if v := real(a.At(i, j)); v != counter {
				t.Fatalf("Expected real(%d,%d) = %f, got %f", i, j, counter, v)
			}

			counter += 1.0
		}
	}
}

func setArray3(a *Array3, dim0, dim1, dim2 int) {
	var counter float32

	for i := range dim0 {
		for j := range dim1 {
			for k := range dim2 {
				a.Set(i, j, k, complex(counter, 0))

				counter += 1.0
			}
		}
	}
}

func verifyArray3(t *testing.T, a *Array3, dim0, dim1, dim2 int) {
	t.Helper()

	var counter float32

	for i := range dim0 {
		for j := range dim1 {
			for k := range dim2 {
				if v := real(a.At(i, j, k)); v != counter {
					t.Fatalf("Expected real(%d,%d,%d) = %f, got %f", i, j, k, counter, v)
				}

				counter += 1.0
			}
		}
	}
}

func verifyImpulse2(t *testing.T, a *Array2, dim0, dim1 int, wantVal float32, impulse bool) {
	t.Helper()

	for i := range dim0 {
		for j := range dim1 {
			want := wantVal

			if impulse && (i != 0 || j != 0) {
				want = 0.0
			}

			testAlmostEqual(t, real(a.At(i, j)), want)
			testAlmostEqual(t, imag(a.At(i, j)), 0.0)
		}
	}
}

func verifyImpulse3(t *testing.T, a *Array3, dim0, dim1, dim2 int, wantVal float32, impulse bool) {
	t.Helper()

	for i := range dim0 {
		for j := range dim1 {
			for k := range dim2 {
				want := wantVal

				if impulse && (i != 0 || j != 0 || k != 0) {
					want = 0.0
				}

				testAlmostEqual(t, real(a.At(i, j, k)), want)
				testAlmostEqual(t, imag(a.At(i, j, k)), 0.0)
			}
		}
	}
}

func verifyFFT2(t *testing.T, a *Array2, dim0, dim1 int, freqX, freqY float64) {
	t.Helper()

	for i := range dim0 {
		for j := range dim1 {
			if (i == int(freqX) || i == dim0-int(freqX)) &&
				(j == int(freqY) || j == dim1-int(freqY)) {
				testAlmostEqual(t, real(a.At(i, j)), float32(dim0*dim1/4))
				testAlmostEqual(t, imag(a.At(i, j)), 0.0)
			} else {
				testAlmostEqual(t, real(a.At(i, j)), 0.0)
				testAlmostEqual(t, imag(a.At(i, j)), 0.0)
			}
		}
	}
}

func verifyFFT3(t *testing.T, a *Array3, dim0, dim1, dim2 int, freqX, freqY, freqZ float64) {
	t.Helper()

	for i := range dim0 {
		for j := range dim1 {
			for k := range dim2 {
				if (i == int(freqX) || i == dim0-int(freqX)) &&
					(j == int(freqY) || j == dim1-int(freqY)) &&
					(k == int(freqZ) || k == dim2-int(freqZ)) {
					testAlmostEqual(t, real(a.At(i, j, k)), float32(dim0*dim1*dim2/8))
					testAlmostEqual(t, imag(a.At(i, j, k)), 0.0)
				} else {
					testAlmostEqual(t, real(a.At(i, j, k)), 0.0)
					testAlmostEqual(t, imag(a.At(i, j, k)), 0.0)
				}
			}
		}
	}
}

func TestFFTAndIFFT1D(t *testing.T) {
	t.Parallel()

	const n = 16

	input := NewArray(n)
	for i := range input.Elems {
		input.Elems[i] = complex(float32(i+1), float32(-i))
	}

	fft := FFT(input)
	if fft.Len() != n {
		t.Fatalf("Expected FFT length %d, got %d", n, fft.Len())
	}

	ifft := IFFT(fft)
	for i := range input.Elems {
		testAlmostEqual(t, real(ifft.Elems[i]), float32(n)*real(input.Elems[i]))
		testAlmostEqual(t, imag(ifft.Elems[i]), float32(n)*imag(input.Elems[i]))
	}
}

func TestFFT2AndIFFT2Impulse(t *testing.T) {
	t.Parallel()

	const dim0, dim1 = 8, 4

	input := NewArray2(dim0, dim1)
	input.Set(0, 0, complex(1, 0))

	fft := FFT2(input)

	verifyImpulse2(t, fft, dim0, dim1, 1.0, false)

	ifft := IFFT2(fft)

	verifyImpulse2(t, ifft, dim0, dim1, float32(dim0*dim1), true)
}

func TestFFT3AndIFFT3Impulse(t *testing.T) {
	t.Parallel()

	const dim0, dim1, dim2 = 4, 3, 2

	input := NewArray3(dim0, dim1, dim2)
	input.Set(0, 0, 0, complex(1, 0))

	fft := FFT3(input)

	verifyImpulse3(t, fft, dim0, dim1, dim2, 1.0, false)

	ifft := IFFT3(fft)

	verifyImpulse3(t, ifft, dim0, dim1, dim2, float32(dim0*dim1*dim2), true)
}
