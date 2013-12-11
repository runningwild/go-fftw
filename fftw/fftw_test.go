package fftw

import (
	"github.com/orfjackal/gospec/src/gospec"
	"math"
)

func NewArraySpec(c gospec.Context) {
	d10 := NewArray(10)
	d100 := NewArray(100)
	d1000 := NewArray(1000)
	c.Specify("Allocates the appropriate memory for 1D arrays.", func() {
		c.Expect(len(d10.Elems), gospec.Equals, 10)
		c.Expect(len(d100.Elems), gospec.Equals, 100)
		c.Expect(len(d1000.Elems), gospec.Equals, 1000)
	})
}

// Make sure that the memory allocated by fftw is getting properly GCed
func GCSpec(c gospec.Context) {
	tot := 0.0
	for i := 0; i < 1000; i++ {
		d := NewArray(100000000)              // Allocate a bunch of memory
		d.Elems[10000] = complex(float64(i), 0) // Do something stupid with it so
		tot += real(d.Elems[10000])             // hopefully it doesn't get optimized out
	}
}

func NewArray2Spec(c gospec.Context) {
	d100x50 := NewArray2(100, 50)
	c.Specify("Allocates the appropriate memory for 2D arrays.", func() {
		n0, n1 := d100x50.Dims()
		c.Expect(n0, gospec.Equals, 100)
		c.Expect(n1, gospec.Equals, 50)
		counter := 0.0
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				d100x50.Set(i, j, complex(counter, 0))
				counter += 1.0
			}
		}
		counter = 0.0
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				c.Expect(real(d100x50.At(i, j)), gospec.Equals, counter)
				counter += 1.0
			}
		}
	})
}

func NewArray3Spec(c gospec.Context) {
	d100x20x10 := NewArray3(100, 20, 10)
	c.Specify("Allocates the appropriate memory for 3D arrays.", func() {
		n0, n1, n2 := d100x20x10.Dims()
		c.Expect(n0, gospec.Equals, 100)
		c.Expect(n1, gospec.Equals, 20)
		c.Expect(n2, gospec.Equals, 10)
		counter := 0.0
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
					c.Expect(real(d100x20x10.At(i, j, k)), gospec.Equals, counter)
					counter += 1.0
				}
			}
		}
	})
}

func peakVerifier(s []complex128, c gospec.Context) {
	c.Expect(real(s[0]), gospec.IsWithin(1e-9), 0.0)
	c.Expect(imag(s[0]), gospec.IsWithin(1e-9), 0.0)
	c.Expect(real(s[1]), gospec.IsWithin(1e-9), float64(len(s))/2)
	c.Expect(imag(s[1]), gospec.IsWithin(1e-9), 0.0)
	for i := 2; i < len(s)-1; i++ {
		c.Expect(real(s[i]), gospec.IsWithin(1e-9), 0.0)
		c.Expect(imag(s[i]), gospec.IsWithin(1e-9), 0.0)
	}
	c.Expect(real(s[len(s)-1]), gospec.IsWithin(1e-9), float64(len(s))/2)
	c.Expect(imag(s[len(s)-1]), gospec.IsWithin(1e-9), 0.0)
}

func FFTSpec(c gospec.Context) {
	signal := NewArray(16)
	new_in := NewArray(16)
	for i := range signal.Elems {
		signal.Elems[i] = complex(float64(i), float64(-i))
		new_in.Elems[i] = signal.Elems[i]
	}

	// A simple real cosine should result in transform with two spikes, one at S[1] and one at S[-1]
	// The spikes should be real and have amplitude equal to len(S)/2 (because fftw doesn't normalize)
	for i := range signal.Elems {
		signal.Elems[i] = complex(math.Cos(float64(i)/float64(len(signal.Elems))*math.Pi*2), 0)
		new_in.Elems[i] = signal.Elems[i]
	}
	MakePlan1(signal, signal, Forward, Estimate).Execute()
	c.Specify("Forward 1D FFT works properly.", func() {
		peakVerifier(signal.Elems, c)
	})
}

func FFT2Spec(c gospec.Context) {
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
			signal.Set(i, j, complex(cosx*cosy, 0))
		}
	}
	MakePlan2(signal, signal, Forward, Estimate).Execute()
	c.Specify("Forward 2D FFT works properly.", func() {
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				if (i == int(fx) || i == dx-int(fx)) &&
					(j == int(fy) || j == dy-int(fy)) {
					c.Expect(real(signal.At(i, j)), gospec.IsWithin(1e-7), float64(dx*dy/4))
					c.Expect(imag(signal.At(i, j)), gospec.IsWithin(1e-7), 0.0)
				} else {
					c.Expect(real(signal.At(i, j)), gospec.IsWithin(1e-7), 0.0)
					c.Expect(imag(signal.At(i, j)), gospec.IsWithin(1e-7), 0.0)
				}
			}
		}
	})
}

func FFT3Spec(c gospec.Context) {
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
				signal.Set(i, j, k, complex(cosx*cosy*cosz, 0))
			}
		}
	}
	MakePlan3(signal, signal, Forward, Estimate).Execute()
	c.Specify("Forward 3D FFT works properly.", func() {
		for i := 0; i < n0; i++ {
			for j := 0; j < n1; j++ {
				for k := 0; k < n2; k++ {
					if (i == int(fx) || i == dx-int(fx)) &&
						(j == int(fy) || j == dy-int(fy)) &&
						(k == int(fz) || k == dz-int(fz)) {
						c.Expect(real(signal.At(i, j, k)), gospec.IsWithin(1e-7), float64(dx*dy*dz/8))
						c.Expect(imag(signal.At(i, j, k)), gospec.IsWithin(1e-7), 0.0)
					} else {
						c.Expect(real(signal.At(i, j, k)), gospec.IsWithin(1e-7), 0.0)
						c.Expect(imag(signal.At(i, j, k)), gospec.IsWithin(1e-7), 0.0)
					}
				}
			}
		}
	})
}
