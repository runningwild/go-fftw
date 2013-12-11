package fftw

// #include <fftw3.h>
import "C"

import (
	"runtime"
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

func MakePlan1(in, out *Array, dir Direction, flag Flag) *Plan {
	// TODO: check that len(in) == len(out)
	n := in.Len()
	var (
		n_    = C.int(n)
		in_   = (*C.fftw_complex)(in.Ptr())
		out_  = (*C.fftw_complex)(out.Ptr())
		dir_  = C.int(dir)
		flag_ = C.uint(flag)
	)
	p := C.fftw_plan_dft_1d(n_, in_, out_, dir_, flag_)
	return newPlan(p)
}

func MakePlan2(in, out *Array2, dir Direction, flag Flag) *Plan {
	// TODO: check that in and out have the same dimensions
	n0, n1 := in.Dims()
	var (
		n0_   = C.int(n0)
		n1_   = C.int(n1)
		in_   = (*C.fftw_complex)(in.Ptr())
		out_  = (*C.fftw_complex)(out.Ptr())
		dir_  = C.int(dir)
		flag_ = C.uint(flag)
	)
	p := C.fftw_plan_dft_2d(n0_, n1_, in_, out_, dir_, flag_)
	return newPlan(p)
}

func MakePlan3(in, out *Array3, dir Direction, flag Flag) *Plan {
	// TODO: check that in and out have the same dimensions
	n0, n1, n2 := in.Dims()
	var (
		n0_   = C.int(n0)
		n1_   = C.int(n1)
		n2_   = C.int(n2)
		in_   = (*C.fftw_complex)(in.Ptr())
		out_  = (*C.fftw_complex)(out.Ptr())
		dir_  = C.int(dir)
		flag_ = C.uint(flag)
	)
	p := C.fftw_plan_dft_3d(n0_, n1_, n2_, in_, out_, dir_, flag_)
	return newPlan(p)
}
