package fftw_test

import (
  "gospec"
  "testing"
)


func TestAllSpecs(t *testing.T) {
  r := gospec.NewRunner()
  r.AddSpec(Alloc1dSpec)
  r.AddSpec(Alloc2dSpec)
  r.AddSpec(FFT1dSpec)
  gospec.MainGoTest(r, t)
}
