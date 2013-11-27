package fftw

// #include <fftw3.h>
import "C"

import (
	"runtime"
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

func plan1D(in, out *Array1D, dir Direction, flag Flag) *plan {
	// TODO: check that len(in) == len(out)
	n := len(in.Elems)
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

func plan2D(in, out *Array2D, dir Direction, flag Flag) *plan {
	// TODO: check that in and out have the same dimensions
	n0, n1 := dims2(in.Elems)
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

func plan3D(in, out *Array3D, dir Direction, flag Flag) *plan {
	// TODO: check that in and out have the same dimensions
	n0, n1, n2 := dims3(in.Elems)
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
