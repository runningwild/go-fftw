package fftw

// #include <fftw3.h>
import "C"

// Data for a 1D signal.
type Array struct {
	Elems []complex128
}

func (a *Array) Len() int {
	return len(a.Elems)
}

func (a *Array) At(i int) complex128 {
	return a.Elems[i]
}

func (a *Array) Set(i int, x complex128) {
	a.Elems[i] = x
}

func (a *Array) ptr() *complex128 {
	return &a.Elems[0]
}

// Allocates memory using fftw_malloc.
func NewArray(n int) *Array {
	elems := make([]complex128, n)
	return &Array{elems}
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

func (a *Array2) ptr() *complex128 {
	return &a.Elems[0]
}

func NewArray2(n0, n1 int) *Array2 {
	elems := make([]complex128, n0*n1)
	return &Array2{[...]int{n0, n1}, elems}
}

// 3D version of Array.
type Array3 struct {
	N     [3]int
	Elems []complex128
}

func (a *Array3) Dims() (n0, n1, n2 int) {
	return a.N[0], a.N[1], a.N[2]
}

func (a *Array3) ptr() *complex128 {
	return &a.Elems[0]
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
	elems := make([]complex128, n0*n1*n2)
	return &Array3{[...]int{n0, n1, n2}, elems}
}

// N-dimensional version of Array.
type ArrayN struct {
	N     []int
	Elems []complex128
}

func (a *ArrayN) Dims() (n []int) {
	return a.N
}

func (a *ArrayN) ptr() *complex128 {
	return &a.Elems[0]
}

func (a *ArrayN) At(i []int) complex128 {
	return a.Elems[a.index(i)]
}

func (a *ArrayN) Set(i []int, x complex128) {
	a.Elems[a.index(i)] = x
}

func (a *ArrayN) index(i []int) int {
	var m int
	for d := range a.N {
		m = m*a.N[d] + i[d]
	}
	return m
}

func NewArrayN(n []int) *ArrayN {
	var a ArrayN
	a.Elems = make([]complex128, prod(n))
	a.N = make([]int, len(n))
	copy(a.N, n)
	return &a
}

func prod(x []int) int {
	t := 1
	for _, xi := range x {
		t *= xi
	}
	return t
}
