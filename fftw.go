package fftw

// #cgo LDFLAGS: -lfftw3 -lm
// #include <fftw3.h>
import "C"

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// Object whose finalizer calls fftw_free.
type Array1d struct {
	Elems []complex128
}

// Object whose finalizer calls fftw_free.
type Array2d struct {
	Elems [][]complex128
}

// Object whose finalizer calls fftw_free.
type Array3d struct {
	Elems [][][]complex128
}

type Plan struct {
	fftw_p C.fftw_plan
}

func destroyPlan(p *Plan) {
	C.fftw_destroy_plan(p.fftw_p)
}

func newPlan(fftw_p C.fftw_plan) *Plan {
	np := new(Plan)
	np.fftw_p = fftw_p
	runtime.SetFinalizer(np, destroyPlan)
	return np
}

func (p *Plan) Execute() {
	C.fftw_execute(p.fftw_p)
}

func (p *Plan) ExecuteNewArray(in, out []complex128) {
	fftw_in := (*C.fftw_complex)((unsafe.Pointer)(&in[0]))
	fftw_out := (*C.fftw_complex)((unsafe.Pointer)(&out[0]))
	C.fftw_execute_dft(p.fftw_p, fftw_in, fftw_out)
}

type Direction int

var Forward Direction = C.FFTW_FORWARD
var Backward Direction = C.FFTW_BACKWARD

type Flag uint

var Estimate Flag = C.FFTW_ESTIMATE
var Measure Flag = C.FFTW_MEASURE

func alloc(n int) []complex128 {
	// Try to allocate memory.
	buffer, err := C.fftw_malloc(C.size_t(16 * n))
	if err != nil {
		// If malloc failed, trigger garbage collector and try again.
		runtime.GC()
		buffer, err = C.fftw_malloc(C.size_t(16 * n))
		if err != nil {
			// If it still failed, then panic.
			panic(fmt.Errorf("could not malloc %d elems: %s", n, err.Error()))
		}
	}

	// Create a slice header for the memory.
	var elems []complex128
	header := (*reflect.SliceHeader)(unsafe.Pointer(&elems))
	*header = reflect.SliceHeader{uintptr(buffer), n, n}
	// Initialize all memory to zero.
	for i := range elems {
		elems[i] = 0
	}
	return elems
}

func Alloc1d(n int) *Array1d {
	elems := alloc(n)
	// Allocate structure with finalizer.
	a := &Array1d{elems}
	runtime.SetFinalizer(a, free1d)
	return a
}

func free1d(x *Array1d) {
	C.fftw_free(unsafe.Pointer(&x.Elems[0]))
}

func Alloc2d(n0, n1 int) *Array2d {
	elems := alloc(n0 * n1)
	r := make([][]complex128, n0)
	for i := range r {
		r[i] = elems[i*n1 : (i+1)*n1]
	}
	// Allocate structure with finalizer.
	a := &Array2d{r}
	runtime.SetFinalizer(a, free2d)
	return a
}

func free2d(x *Array2d) {
	C.fftw_free(unsafe.Pointer(&x.Elems[0][0]))
}

func Alloc3d(n0, n1, n2 int) *Array3d {
	elems := alloc(n0 * n1 * n2)
	r := make([][][]complex128, n0)
	for i := range r {
		b := make([][]complex128, n1)
		for j := range b {
			b[j] = elems[i*(n1*n2)+j*n2 : i*(n1*n2)+(j+1)*n2]
		}
		r[i] = b
	}
	// Allocate structure with finalizer.
	a := &Array3d{r}
	runtime.SetFinalizer(a, free3d)
	return a
}

func free3d(x *Array3d) {
	C.fftw_free(unsafe.Pointer(&x.Elems[0][0][0]))
}

func PlanDft1d(in, out *Array1d, dir Direction, flag Flag) *Plan {
	// TODO: check that len(in) == len(out)
	fftw_in := (*C.fftw_complex)(unsafe.Pointer(&in.Elems[0]))
	fftw_out := (*C.fftw_complex)(unsafe.Pointer(&out.Elems[0]))
	p := C.fftw_plan_dft_1d(C.int(len(in.Elems)), fftw_in, fftw_out, C.int(dir), C.uint(flag))
	return newPlan(p)
}

func PlanDft2d(in, out *Array2d, dir Direction, flag Flag) *Plan {
	// TODO: check that in and out have the same dimensions
	fftw_in := (*C.fftw_complex)(unsafe.Pointer(&in.Elems[0][0]))
	fftw_out := (*C.fftw_complex)(unsafe.Pointer(&out.Elems[0][0]))
	n0 := len(in.Elems)
	n1 := len(in.Elems[0])
	p := C.fftw_plan_dft_2d(C.int(n0), C.int(n1), fftw_in, fftw_out, C.int(dir), C.uint(flag))
	return newPlan(p)
}

func PlanDft3d(in, out *Array3d, dir Direction, flag Flag) *Plan {
	// TODO: check that in and out have the same dimensions
	fftw_in := (*C.fftw_complex)(unsafe.Pointer(&in.Elems[0][0][0]))
	fftw_out := (*C.fftw_complex)(unsafe.Pointer(&out.Elems[0][0][0]))
	n0 := len(in.Elems)
	n1 := len(in.Elems[0])
	n2 := len(in.Elems[0][0])
	p := C.fftw_plan_dft_3d(C.int(n0), C.int(n1), C.int(n2), fftw_in, fftw_out, C.int(dir), C.uint(flag))
	return newPlan(p)
}

//	// TODO: Once we can create go arrays out of pre-existing data we can do these real-to-complex and complex-to-real
//	//       transforms in-place.
//	// The real-to-complex and complex-to-real transforms save roughly a factor of two in time and space, with
//	// the following caveats:
//	// 1. The real array is of size N, the complex array is of size N/2+1.
//	// 2. The output array contains only the non-redundant output, the complete output is symmetric and the last half
//	//    is the complex conjugate of the first half.
//	// 3. Doing a complex-to-real transform destroys the input signal.
//	func PlanDftR2C1d(in []float64, out []complex128, flag Flag) *Plan {
//		// TODO: check that in and out have the appropriate dimensions
//		fftw_in := (*C.double)(unsafe.Pointer(&in.Elems[0]))
//		fftw_out := (*C.fftw_complex)(unsafe.Pointer(&out.Elems[0]))
//		p := C.fftw_plan_dft_r2c_1d(C.int(len(in.Elems)), fftw_in, fftw_out, C.uint(flag))
//		return newPlan(p)
//	}
//
//	// Note: Executing this plan will destroy the data contained by in
//	func PlanDftC2R1d(in []complex128, out []float64, flag Flag) *Plan {
//		// TODO: check that in and out have the appropriate dimensions
//		fftw_in := (*C.fftw_complex)(unsafe.Pointer(&in.Elems[0]))
//		fftw_out := (*C.double)(unsafe.Pointer(&out.Elems[0]))
//		p := C.fftw_plan_dft_c2r_1d(C.int(len(out.Elems)), fftw_in, fftw_out, C.uint(flag))
//		return newPlan(p)
//	}
