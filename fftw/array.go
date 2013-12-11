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
type Array struct {
	Elems []complex128
}

func (a *Array) Len() int {
	return len(a.Elems)
}

func (a *Array) Ptr() unsafe.Pointer {
	return unsafe.Pointer(&a.Elems[0])
}

func NewArray(n int) *Array {
	elems := allocCmplx(n)
	// Allocate structure with finalizer.
	a := &Array{elems}
	runtime.SetFinalizer(a, free1)
	return a
}

// 2D version of Array.
type Array2 struct {
	Elems [][]complex128
}

func (a *Array2) Dims() (n0, n1 int) {
	return dims2(a.Elems)
}

func (a *Array2) Ptr() unsafe.Pointer {
	return unsafe.Pointer(&a.Elems[0][0])
}

func NewArray2(n0, n1 int) *Array2 {
	elems := allocCmplx(n0 * n1)
	r := make([][]complex128, n0)
	for i := range r {
		r[i] = elems[i*n1 : (i+1)*n1]
	}
	// Allocate structure with finalizer.
	a := &Array2{r}
	runtime.SetFinalizer(a, free2)
	return a
}

// 3D version of Array.
type Array3 struct {
	Elems [][][]complex128
}

func (a *Array3) Dims() (n0, n1, n2 int) {
	return dims3(a.Elems)
}

func (a *Array3) Ptr() unsafe.Pointer {
	return unsafe.Pointer(&a.Elems[0][0][0])
}

func NewArray3(n0, n1, n2 int) *Array3 {
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
	a := &Array3{r}
	runtime.SetFinalizer(a, free3)
	return a
}

type array interface {
	Ptr() unsafe.Pointer
}

func free(x array) {
	C.fftw_free(x.Ptr())
}

func free1(x *Array) {
	free(x)
}

func free2(x *Array2) {
	free(x)
}

func free3(x *Array3) {
	free(x)
}
