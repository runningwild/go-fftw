package fftw

// #include <fftw3.h>
import "C"

import (
  "runtime"
  "unsafe"
)

type plan struct {
  fftw_p C.fftw_plan
}
func destroyPlan(p *plan) {
  C.fftw_destroy_plan(p.fftw_p)
}
func newPlan(fftw_p C.fftw_plan) *plan {
  np := new(plan)
  np.fftw_p = fftw_p
  runtime.SetFinalizer(np, destroyPlan)
  return np
}
func (p *plan) Execute() {
  C.fftw_execute(p.fftw_p)
}


type Direction int
var Forward  Direction = C.FFTW_FORWARD
var Backward Direction = C.FFTW_BACKWARD

func Alloc1d(n int) []complex128 {
  return make([]complex128, n)
}
func Alloc2d(n0,n1 int) [][]complex128 {
  a := make([]complex128, n0*n1)
  r := make([][]complex128, n0)
  for i := range r {
    r[i] = a[i*n1 : (i+1)*n1]
  }
  return r
}


func PlanDft1d(in,out []complex128, dir Direction) *plan {
  // TODO: check that len(in) == len(out)
  fftw_in := (*C.fftw_complex)((unsafe.Pointer)(&in[0]))
  fftw_out := (*C.fftw_complex)((unsafe.Pointer)(&out[0]))
  p := C.fftw_plan_dft_1d((C.int)(len(in)), fftw_in, fftw_out, C.int(dir), C.FFTW_ESTIMATE)
  return newPlan(p)
}

func PlanDft2d(in,out [][]complex128, dir Direction) *plan {
  // TODO: check that in and out have the same dimensions
  fftw_in := (*C.fftw_complex)((unsafe.Pointer)(&in[0][0]))
  fftw_out := (*C.fftw_complex)((unsafe.Pointer)(&out[0][0]))
  p := C.fftw_plan_dft_2d((C.int)(len(in)), (C.int)(len(in[0])), fftw_in, fftw_out, C.int(dir), C.FFTW_ESTIMATE)
  return newPlan(p)
}

