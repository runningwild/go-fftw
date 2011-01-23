package fftw

// #include <fftw3.h>
import "C"

import "unsafe"

type Plan struct {
  plan C.fftw_plan
}

func Alloc(n int) []complex128 {
  size := (C.size_t)(n * 16)
  return *(*[]complex128)((unsafe.Pointer)(C.fftw_malloc(size)))
}
