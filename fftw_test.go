package fftw_test

import (
  "fftw"
  "gospec"
  . "gospec"
  "math"
)


func Alloc1dSpec(c gospec.Context) {
  d10   := fftw.Alloc1d(10)
  d100  := fftw.Alloc1d(100)
  d1000 := fftw.Alloc1d(1000)
  c.Specify("Allocates the appropriate memory for 1d arrays.", func() {
    c.Expect(len(d10), Equals, 10)
    c.Expect(len(d100), Equals, 100)
    c.Expect(len(d1000), Equals, 1000)
  })
}

func Alloc2dSpec(c gospec.Context) {
  d1000x500 := fftw.Alloc2d(1000, 500)
  c.Specify("Allocates the appropriate memory for 2d arrays.", func() {
    c.Expect(len(d1000x500), Equals, 1000)
    for _,v := range d1000x500 {
      c.Expect(len(v), Equals, 500)
    }
  })
}

func FFT1dSpec(c gospec.Context) {
  signal := fftw.Alloc1d(16)
  for i := range signal {
    signal[i] = cmplx(float64(i), float64(-i))
  }
  forward := fftw.PlanDft1d(signal, signal, fftw.Forward, fftw.Estimate)
  c.Specify("Creating a plan doesn't overwrite an existing array if fftw.Estimate is used.", func() {
    for i := range signal {
      c.Expect(signal[i], Equals, cmplx(float64(i), float64(-i)))
    }
  })

  // A simple real cosine should result in transform with two spikes, one at S[1] and one at S[-1]
  // The spikes should be real and have amplitude equal to len(S) (because fftw doesn't normalize)
  for i := range signal {
    signal[i] = cmplx(math.Cos(float64(i) / float64(len(signal)) * math.Pi * 2), 0)
  }
  forward.Execute()
  c.Specify("Forward 1d FFT works properly.", func() {
    c.Expect(real(signal[0]), IsWithin(1e-9), float64(0))
    c.Expect(imag(signal[0]), IsWithin(1e-9), float64(0))
    c.Expect(real(signal[1]), IsWithin(1e-9), float64(len(signal))/2)
    c.Expect(imag(signal[1]), IsWithin(1e-9), float64(0))
    for i := 2; i < len(signal) - 1; i++ {
      c.Expect(real(signal[i]), IsWithin(1e-9), float64(0))
      c.Expect(imag(signal[i]), IsWithin(1e-9), float64(0))
    }
    c.Expect(real(signal[len(signal)-1]), IsWithin(1e-9), float64(len(signal))/2)
    c.Expect(imag(signal[len(signal)-1]), IsWithin(1e-9), float64(0))
  })
}

