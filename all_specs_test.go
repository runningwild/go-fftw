package fftw_test

import (
  "gospec"
  "testing"
)


func TestAllSpecs(t *testing.T) {
  r := gospec.NewRunner()
  r.AddSpec(GCSpec)
  r.AddSpec(Alloc1dSpec)
  r.AddSpec(Alloc2dSpec)
  r.AddSpec(Alloc3dSpec)
  gospec.MainGoTest(r, t)
  
  // TODO: Investigate a less stupid way of doing tests in serial
  // TODO: Investigate what about fftw makes tests need to be serial
  r = gospec.NewRunner()
  r.AddSpec(FFT1dSpec)
  gospec.MainGoTest(r, t)
  r = gospec.NewRunner()
  r.AddSpec(FFT2dSpec)
  gospec.MainGoTest(r, t)
  r = gospec.NewRunner()
  r.AddSpec(FFT3dSpec)
  gospec.MainGoTest(r, t)
  r = gospec.NewRunner()
  r.AddSpec(FFTR2CSpec)
  gospec.MainGoTest(r, t)
  r = gospec.NewRunner()
  r.AddSpec(FFTC2RSpec)
  gospec.MainGoTest(r, t)
}
