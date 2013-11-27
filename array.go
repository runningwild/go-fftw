package fftw

// #include <fftw3.h>
import "C"

import (
	"runtime"
	"unsafe"
)

// Contains memory allocated using fftw_malloc.
// Finalizer invokes fftw_free.
// Do not modify location of slice!
type Array1D struct {
	Elems []complex128
}

func (a *Array1D) Len() int {
	return len(a.Elems)
}

func (a *Array1D) Ptr() unsafe.Pointer {
	return unsafe.Pointer(&a.Elems[0])
}

// In-place transform.
func (a *Array1D) FFT(dir Direction, flag Flag) {
	plan1D(a, a, dir, flag).Execute()
}

func NewArray1D(n int) *Array1D {
	elems := allocCmplx(n)
	// Allocate structure with finalizer.
	a := &Array1D{elems}
	runtime.SetFinalizer(a, free1D)
	return a
}

func free1D(x *Array1D) {
	C.fftw_free(x.Ptr())
}

// 2D version of Array1D.
type Array2D struct {
	Elems [][]complex128
}

func (a *Array2D) Dims() (n0, n1 int) {
	return dims2(a.Elems)
}

func (a *Array2D) Ptr() unsafe.Pointer {
	return unsafe.Pointer(&a.Elems[0][0])
}

// In-place transform.
func (a *Array2D) FFT(dir Direction, flag Flag) {
	plan2D(a, a, dir, flag).Execute()
}

func NewArray2D(n0, n1 int) *Array2D {
	elems := allocCmplx(n0 * n1)
	r := make([][]complex128, n0)
	for i := range r {
		r[i] = elems[i*n1 : (i+1)*n1]
	}
	// Allocate structure with finalizer.
	a := &Array2D{r}
	runtime.SetFinalizer(a, free2D)
	return a
}

func free2D(x *Array2D) {
	C.fftw_free(x.Ptr())
}

// 3D version of Array1D.
type Array3D struct {
	Elems [][][]complex128
}

func (a *Array3D) Dims() (n0, n1, n2 int) {
	return dims3(a.Elems)
}

func (a *Array3D) Ptr() unsafe.Pointer {
	return unsafe.Pointer(&a.Elems[0][0][0])
}

// In-place transform.
func (a *Array3D) FFT(dir Direction, flag Flag) {
	plan3D(a, a, dir, flag).Execute()
}

func NewArray3D(n0, n1, n2 int) *Array3D {
	elems := allocCmplx(n0 * n1 * n2)
	r := make([][][]complex128, n0)
	for i := range r {
		b := make([][]complex128, n1)
		for j := range b {
			b[j] = elems[i*(n1*n2)+j*n2 : i*(n1*n2)+(j+1)*n2]
		}
		r[i] = b
	}
	// Allocate structure with finalizer.
	a := &Array3D{r}
	runtime.SetFinalizer(a, free3D)
	return a
}

func free3D(x *Array3D) {
	C.fftw_free(x.Ptr())
}
