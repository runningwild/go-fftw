package fftw32

import (
	"github.com/orfjackal/gospec/src/gospec"
	"testing"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(GCSpec)
	r.AddSpec(NewArraySpec)
	r.AddSpec(NewArray2Spec)
	r.AddSpec(NewArray3Spec)
	gospec.MainGoTest(r, t)

	// TODO: Investigate a less stupid way of doing tests in serial
	// TODO: Investigate what about fftw makes tests need to be serial
	r = gospec.NewRunner()
	r.AddSpec(FFTSpec)
	gospec.MainGoTest(r, t)
	r = gospec.NewRunner()
	r.AddSpec(FFT2Spec)
	gospec.MainGoTest(r, t)
	r = gospec.NewRunner()
	r.AddSpec(FFT3Spec)
	gospec.MainGoTest(r, t)
}
