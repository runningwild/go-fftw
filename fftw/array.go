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
	N     [2]int
	Elems []complex128
}

func (a *Array2) Dims() (n0, n1 int) {
	return a.N[0], a.N[1]
}

func (a *Array2) At(i0, i1 int) complex128 {
	return a.Elems[a.index(i0, i1)]
}

func (a *Array2) Set(i0, i1 int, x complex128) {
	a.Elems[a.index(i0, i1)] = x
}

func (a *Array2) index(i0, i1 int) int {
	return i1 + a.N[1]*i0
}

func (a *Array2) Ptr() unsafe.Pointer {
	return unsafe.Pointer(&a.Elems[0])
}

func NewArray2(n0, n1 int) *Array2 {
	elems := allocCmplx(n0 * n1)
	// Allocate structure with finalizer.
	a := &Array2{[...]int{n0, n1}, elems}
	runtime.SetFinalizer(a, free2)
	return a
}

// 3D version of Array.
type Array3 struct {
	N     [3]int
	Elems []complex128
}

func (a *Array3) Dims() (n0, n1, n2 int) {
	return a.N[0], a.N[1], a.N[2]
}

func (a *Array3) Ptr() unsafe.Pointer {
	return unsafe.Pointer(&a.Elems[0])
}

func (a *Array3) At(i0, i1, i2 int) complex128 {
	return a.Elems[a.index(i0, i1, i2)]
}

func (a *Array3) Set(i0, i1, i2 int, x complex128) {
	a.Elems[a.index(i0, i1, i2)] = x
}

func (a *Array3) index(i0, i1, i2 int) int {
	return i2 + a.N[2]*(i1+i0*a.N[1])
}

func NewArray3(n0, n1, n2 int) *Array3 {
	elems := allocCmplx(n0 * n1 * n2)
	// Allocate structure with finalizer.
	a := &Array3{[...]int{n0, n1, n2}, elems}
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
