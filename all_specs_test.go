package fftw

import (
	"github.com/orfjackal/gospec/src/gospec"
	"testing"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(GCSpec)
	r.AddSpec(NewArray1DSpec)
	r.AddSpec(NewArray2DSpec)
	r.AddSpec(NewArray3DSpec)
	gospec.MainGoTest(r, t)

	// TODO: Investigate a less stupid way of doing tests in serial
	// TODO: Investigate what about fftw makes tests need to be serial
	r = gospec.NewRunner()
	r.AddSpec(FFT1DSpec)
	gospec.MainGoTest(r, t)
	r = gospec.NewRunner()
	r.AddSpec(FFT2DSpec)
	gospec.MainGoTest(r, t)
	r = gospec.NewRunner()
	r.AddSpec(FFT3DSpec)
	gospec.MainGoTest(r, t)
}
