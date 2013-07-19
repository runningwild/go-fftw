package fftw

// #cgo LDFLAGS: -lfftw3 -lm
// #include <fftw3.h>
import "C"

import (
	"reflect"
	"runtime"
	"unsafe"
)

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

func Alloc1d(n int) []complex128 {
	buffer := (unsafe.Pointer)(C.fftw_malloc((C.size_t)(16 * n)))
	var slice []complex128
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	*header = reflect.SliceHeader{uintptr(buffer), n, n}
	return slice
}

func Alloc2d(n0, n1 int) [][]complex128 {
	a := Alloc1d(n0 * n1)
	r := make([][]complex128, n0)
	for i := range r {
		r[i] = a[i*n1 : (i+1)*n1]
	}
	return r
}

func Alloc3d(n0, n1, n2 int) [][][]complex128 {
	a := Alloc1d(n0 * n1 * n2)
	r := make([][][]complex128, n0)
	for i := range r {
		b := make([][]complex128, n1)
		for j := range b {
			b[j] = a[i*(n1*n2)+j*n2 : i*(n1*n2)+(j+1)*n2]
		}
		r[i] = b
	}
	return r
}

func PlanDft1d(in, out []complex128, dir Direction, flag Flag) *Plan {
	// TODO: check that len(in) == len(out)

	fftw_in := (*C.fftw_complex)((unsafe.Pointer)(&in[0]))
	fftw_out := (*C.fftw_complex)((unsafe.Pointer)(&out[0]))
	p := C.fftw_plan_dft_1d((C.int)(len(in)), fftw_in, fftw_out, C.int(dir), C.uint(flag))
	return newPlan(p)
}

func PlanDft2d(in, out [][]complex128, dir Direction, flag Flag) *Plan {
	// TODO: check that in and out have the same dimensions
	fftw_in := (*C.fftw_complex)((unsafe.Pointer)(&in[0][0]))
	fftw_out := (*C.fftw_complex)((unsafe.Pointer)(&out[0][0]))
	p := C.fftw_plan_dft_2d((C.int)(len(in)), (C.int)(len(in[0])), fftw_in, fftw_out, C.int(dir), C.uint(flag))
	return newPlan(p)
}

func PlanDft3d(in, out [][][]complex128, dir Direction, flag Flag) *Plan {
	// TODO: check that in and out have the same dimensions
	fftw_in := (*C.fftw_complex)((unsafe.Pointer)(&in[0][0][0]))
	fftw_out := (*C.fftw_complex)((unsafe.Pointer)(&out[0][0][0]))
	p := C.fftw_plan_dft_3d((C.int)(len(in)), (C.int)(len(in[0])), (C.int)(len(in[0][0])), fftw_in, fftw_out, C.int(dir), C.uint(flag))
	return newPlan(p)
}

// TODO: Once we can create go arrays out of pre-existing data we can do these real-to-complex and complex-to-real
//       transforms in-place.
// The real-to-complex and complex-to-real transforms save roughly a factor of two in time and space, with
// the following caveats:
// 1. The real array is of size N, the complex array is of size N/2+1.
// 2. The output array contains only the non-redundant output, the complete output is symmetric and the last half
//    is the complex conjugate of the first half.
// 3. Doing a complex-to-real transform destroys the input signal.
func PlanDftR2C1d(in []float64, out []complex128, flag Flag) *Plan {
	// TODO: check that in and out have the appropriate dimensions
	fftw_in := (*C.double)((unsafe.Pointer)(&in[0]))
	fftw_out := (*C.fftw_complex)((unsafe.Pointer)(&out[0]))
	p := C.fftw_plan_dft_r2c_1d((C.int)(len(in)), fftw_in, fftw_out, C.uint(flag))
	return newPlan(p)
}

// Note: Executing this plan will destroy the data contained by in
func PlanDftC2R1d(in []complex128, out []float64, flag Flag) *Plan {
	// TODO: check that in and out have the appropriate dimensions
	fftw_in := (*C.fftw_complex)((unsafe.Pointer)(&in[0]))
	fftw_out := (*C.double)((unsafe.Pointer)(&out[0]))
	p := C.fftw_plan_dft_c2r_1d((C.int)(len(out)), fftw_in, fftw_out, C.uint(flag))
	return newPlan(p)
}
