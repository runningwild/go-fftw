/*
Package fftw is a cgo wrapper around the Fastest Fourier Transform in the West.

http://www.fftw.org/

Provides simple functions to compute a transform without destroying existing data.
	x := fftw.NewArray(100)
	// ...
	xhat := fftw.FFT(x)
	x = fftw.IFFT(xhat)
Beware that scaling is the same as in FFTW, so that computing forward and then inverse transforms scales the original input by the length of the sequence.

Use fftw.XxxTo() to do in-place operations
	fftw.FFTTo(x, x)
	fftw.IFFTTo(x, x)
or re-use pre-allocated arrays:
	fftw.FFTTo(xhat, x)
	fftw.IFFTTo(x, xhat)

It is also possible to use fftw.Plan explicitly:
	p := fftw.NewPlan(x, x, fftw.Forward, fftw.Estimate)
	defer p.Destroy()
	p.Execute()
Beware that when using fftw.Measure instead of fftw.Estimate, the contents of the array may be overwritten by fftw.NewPlan().
For this reason, all functions similar to fftw.FFT() and fftw.FFTTo() use fftw.Estimate.

It's possible to use FFTW with memory not allocated by fftw.NewArrayX():
	x := make([]complex128, 100)
	xhat := fftw.FFT(&fftw.Array{x})
*/
package fftw
