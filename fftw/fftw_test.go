package fftw

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewArray(t *testing.T) {
	d10 := NewArray(10)
	d100 := NewArray(100)
	d1000 := NewArray(1000)
	Convey("Allocates the appropriate memory for 1D arrays.", t, func() {
		So(len(d10.Elems), ShouldEqual, 10)
		So(len(d100.Elems), ShouldEqual, 100)
		So(len(d1000.Elems), ShouldEqual, 1000)
	})
}

// Make sure that the memory allocated by fftw is getting properly GCed
func TestGC(t *testing.T) {
	var tot float64 = 0.0
	for i := 0; i < 1000; i++ {
		d := NewArray(1000000)                  // Allocate a bunch of memory
		d.Elems[10000] = complex(float64(i), 0) // Do something stupid with it so
		tot += real(d.Elems[10000])             // hopefully it doesn't get optimized out
	}
}

func TestNewArray2(t *testing.T) {
	d100x50 := NewArray2(100, 50)
	Convey("Allocates the appropriate memory for 2D arrays.", t, func() {
		n0, n1 := d100x50.Dims()
		So(n0, ShouldEqual, 100)
		So(n1, ShouldEqual, 50)
		var counter float64 = 0.0
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				d100x50.Set(i, j, complex(counter, 0))
				counter += 1.0
			}
		}
		counter = 0.0
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				So(real(d100x50.At(i, j)), ShouldEqual, counter)
				counter += 1.0
			}
		}
	})
}

func TestNewArray3(t *testing.T) {
	d100x20x10 := NewArray3(100, 20, 10)
	Convey("Allocates the appropriate memory for 3D arrays.", t, func() {
		n0, n1, n2 := d100x20x10.Dims()
		So(n0, ShouldEqual, 100)
		So(n1, ShouldEqual, 20)
		So(n2, ShouldEqual, 10)
		var counter float64 = 0.0
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				for k := 0; k < n2; k++ {
					d100x20x10.Set(i, j, k, complex(counter, 0))
					counter += 1.0
				}
			}
		}
		counter = 0.0
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				for k := 0; k < n2; k++ {
					So(real(d100x20x10.At(i, j, k)), ShouldEqual, counter)
					counter += 1.0
				}
			}
		}
	})
}

func peakVerifier(s []complex128) {
	So(real(s[0]), ShouldAlmostEqual, 0.0)
	So(imag(s[0]), ShouldAlmostEqual, 0.0)
	So(real(s[1]), ShouldAlmostEqual, float64(len(s))/2)
	So(imag(s[1]), ShouldAlmostEqual, 0.0)
	for i := 2; i < len(s)-1; i++ {
		So(real(s[i]), ShouldAlmostEqual, 0.0)
		So(imag(s[i]), ShouldAlmostEqual, 0.0)
	}
	So(real(s[len(s)-1]), ShouldAlmostEqual, float64(len(s))/2)
	So(imag(s[len(s)-1]), ShouldAlmostEqual, 0.0)
}

func TestFFT(t *testing.T) {
	signal := NewArray(16)
	new_in := NewArray(16)
	for i := range signal.Elems {
		signal.Elems[i] = complex(float64(i), float64(-i))
		new_in.Elems[i] = signal.Elems[i]
	}

	// A simple real cosine should result in transform with two spikes, one at S[1] and one at S[-1]
	// The spikes should be real and have amplitude equal to len(S)/2 (because fftw doesn't normalize)
	for i := range signal.Elems {
		signal.Elems[i] = complex(float64(math.Cos(float64(i)/float64(len(signal.Elems))*math.Pi*2)), 0)
		new_in.Elems[i] = signal.Elems[i]
	}
	NewPlan(signal, signal, Forward, Estimate).Execute().Destroy()
	Convey("Forward 1D FFT works properly.", t, func() {
		peakVerifier(signal.Elems)
	})
}

func TestFFT2(t *testing.T) {
	signal := NewArray2(64, 8)
	n0, n1 := signal.Dims()
	for i := 0; i < n0; i++ {
		for j := 0; j < n1; j++ {
			signal.Set(i, j, complex(float64(i+j), float64(-i-j)))
		}
	}

	// As long as fx < dx/2 and fy < dy/2, where dx and dy are the lengths in each dimension,
	// there will be 2^n spikes, where n is the number of dimensions.  Each spike will be
	// real and have magnitude equal to dx*dy / 2^n
	dx := n0
	fx := float64(dx) / 4
	dy := n1
	fy := float64(dy) / 4
	for i := 0; i < n0; i++ {
		for j := 0; j < n1; j++ {
			cosx := math.Cos(float64(i) / float64(dx) * fx * math.Pi * 2)
			cosy := math.Cos(float64(j) / float64(dy) * fy * math.Pi * 2)
			signal.Set(i, j, complex(float64(cosx*cosy), 0))
		}
	}
	NewPlan2(signal, signal, Forward, Estimate).Execute().Destroy()
	Convey("Forward 2D FFT works properly.", t, func() {
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				if (i == int(fx) || i == dx-int(fx)) &&
					(j == int(fy) || j == dy-int(fy)) {
					So(real(signal.At(i, j)), ShouldAlmostEqual, float64(dx*dy/4))
					So(imag(signal.At(i, j)), ShouldAlmostEqual, 0.0)
				} else {
					So(real(signal.At(i, j)), ShouldAlmostEqual, 0.0)
					So(imag(signal.At(i, j)), ShouldAlmostEqual, 0.0)
				}
			}
		}
	})
}

func TestFFT3(t *testing.T) {
	signal := NewArray3(32, 16, 8)

	n0, n1, n2 := signal.Dims()
	for i := 0; i < n0; i++ {
		for j := 0; j < n1; j++ {
			for k := 0; k < n2; k++ {
				signal.Set(i, j, k, complex(float64(i+j+k), float64(-i-j-k)))
			}
		}
	}

	// As long as fx < dx/2, fy < dy/2, and fz < dz/2, where dx,dy,dz  are the lengths in
	// each dimension, there will be 2^n spikes, where n is the number of dimensions.
	// Each spike will be real and have magnitude equal to dx*dy*dz / 2^n
	dx := n0
	fx := float64(dx) / 4
	dy := n1
	fy := float64(dy) / 4
	dz := n2
	fz := float64(dz) / 4
	for i := 0; i < n0; i++ {
		for j := 0; j < n1; j++ {
			for k := 0; k < n2; k++ {
				cosx := math.Cos(float64(i) / float64(dx) * fx * math.Pi * 2)
				cosy := math.Cos(float64(j) / float64(dy) * fy * math.Pi * 2)
				cosz := math.Cos(float64(k) / float64(dz) * fz * math.Pi * 2)
				signal.Set(i, j, k, complex(float64(cosx*cosy*cosz), 0))
			}
		}
	}
	NewPlan3(signal, signal, Forward, Estimate).Execute().Destroy()
	Convey("Forward 3D FFT works properly.", t, func() {
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				for k := 0; k < n2; k++ {
					if (i == int(fx) || i == dx-int(fx)) &&
						(j == int(fy) || j == dy-int(fy)) &&
						(k == int(fz) || k == dz-int(fz)) {
						So(real(signal.At(i, j, k)), ShouldAlmostEqual, float64(dx*dy*dz/8))
						So(imag(signal.At(i, j, k)), ShouldAlmostEqual, 0.0)
					} else {
						So(real(signal.At(i, j, k)), ShouldAlmostEqual, 0.0)
						So(imag(signal.At(i, j, k)), ShouldAlmostEqual, 0.0)
					}
				}
			}
		}
	})
}
