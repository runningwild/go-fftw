package fftw32

// #include <fftw3.h>
import "C"

import (
	"runtime"
	"sync"
	"unsafe"
)

// According to fftw's doc on multithreading, creation and destruction of plans should be single-
// threaded, so this will serve to synchronize that stuff, and hopefull multi-threaded is ok as long
// as it's all synchronous.
//
//nolint:gochecknoglobals
var createDestroyMu sync.Mutex

type Plan struct {
	fftwP C.fftwf_plan
	pin   runtime.Pinner
}

// NewPlanForSize allocates input/output arrays of length n and returns a plan for them.
//
// This is a convenience helper for callers who want to create and reuse a plan based
// solely on the transform size.
func NewPlanForSize(n int, dir Direction, flag Flag) (p *Plan, in *Array, out *Array) {
	if n <= 0 {
		panic("fftw32: n must be > 0")
	}
	in = NewArray(n)
	out = NewArray(n)
	p = NewPlan(in, out, dir, flag)
	return p, in, out
}

func NewPlan(in, out *Array, dir Direction, flag Flag) *Plan {
	if in == nil || out == nil {
		panic("fftw32: input and output must be non-nil")
	}
	if in.Len() == 0 {
		panic("fftw32: input and output must be non-empty")
	}
	if in.Len() != out.Len() {
		panic("fftw32: input and output lengths must match")
	}
	plan := &Plan{fftwP: nil, pin: runtime.Pinner{}}
	plan.pin.Pin(in.ptr())
	plan.pin.Pin(out.ptr())
	n := in.Len()
	var (
		numElems = C.int(n)
		inPtr    = (*C.fftwf_complex)(unsafe.Pointer(in.ptr()))
		outPtr   = (*C.fftwf_complex)(unsafe.Pointer(out.ptr()))
		dir_     = C.int(dir)
		flag_    = C.uint(flag)
	)
	createDestroyMu.Lock()
	plan.fftwP = C.fftwf_plan_dft_1d(numElems, inPtr, outPtr, dir_, flag_)
	createDestroyMu.Unlock()
	runtime.SetFinalizer(plan, planFinalizer)
	return plan
}

func NewPlan2(in, out *Array2, dir Direction, flag Flag) *Plan {
	if in == nil || out == nil {
		panic("fftw32: input and output must be non-nil")
	}
	in0, in1 := in.Dims()
	out0, out1 := out.Dims()
	if in0 <= 0 || in1 <= 0 {
		panic("fftw32: input and output must be non-empty")
	}
	if in0 != out0 || in1 != out1 {
		panic("fftw32: input and output dimensions must match")
	}
	plan := &Plan{fftwP: nil, pin: runtime.Pinner{}}
	plan.pin.Pin(in.ptr())
	plan.pin.Pin(out.ptr())
	var (
		dim0   = C.int(in0)
		dim1   = C.int(in1)
		inPtr  = (*C.fftwf_complex)(unsafe.Pointer(in.ptr()))
		outPtr = (*C.fftwf_complex)(unsafe.Pointer(out.ptr()))
		dir_   = C.int(dir)
		flag_  = C.uint(flag)
	)
	createDestroyMu.Lock()
	plan.fftwP = C.fftwf_plan_dft_2d(dim0, dim1, inPtr, outPtr, dir_, flag_)
	createDestroyMu.Unlock()
	runtime.SetFinalizer(plan, planFinalizer)
	return plan
}

func NewPlan3(in, out *Array3, dir Direction, flag Flag) *Plan {
	if in == nil || out == nil {
		panic("fftw32: input and output must be non-nil")
	}
	in0, in1, in2 := in.Dims()
	out0, out1, out2 := out.Dims()
	if in0 <= 0 || in1 <= 0 || in2 <= 0 {
		panic("fftw32: input and output must be non-empty")
	}
	if in0 != out0 || in1 != out1 || in2 != out2 {
		panic("fftw32: input and output dimensions must match")
	}
	plan := &Plan{fftwP: nil, pin: runtime.Pinner{}}
	plan.pin.Pin(in.ptr())
	plan.pin.Pin(out.ptr())
	var (
		dim0   = C.int(in0)
		dim1   = C.int(in1)
		dim2   = C.int(in2)
		inPtr  = (*C.fftwf_complex)(unsafe.Pointer(in.ptr()))
		outPtr = (*C.fftwf_complex)(unsafe.Pointer(out.ptr()))
		dir_   = C.int(dir)
		flag_  = C.uint(flag)
	)
	createDestroyMu.Lock()
	plan.fftwP = C.fftwf_plan_dft_3d(dim0, dim1, dim2, inPtr, outPtr, dir_, flag_)
	createDestroyMu.Unlock()
	runtime.SetFinalizer(plan, planFinalizer)
	return plan
}

func (p *Plan) Execute() *Plan {
	C.fftwf_execute(p.fftwP)
	return p
}

// String returns FFTW's textual description of the plan.
//
// This uses fftwf_sprint_plan under the hood.
func (p *Plan) String() string {
	if p == nil {
		return "<nil>"
	}
	if p.fftwP == nil {
		return "<destroyed fftw32 plan>"
	}
	cstr := C.fftwf_sprint_plan(p.fftwP)
	if cstr == nil {
		return ""
	}
	defer C.fftwf_free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func (p *Plan) Destroy() {
	createDestroyMu.Lock()
	if p.fftwP != nil {
		C.fftwf_destroy_plan(p.fftwP)
	}
	p.fftwP = nil
	createDestroyMu.Unlock()
	p.pin.Unpin()
}

func planFinalizer(p *Plan) {
	p.Destroy()
}
