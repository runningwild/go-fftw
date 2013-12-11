/*
Cgo wrapper for the Fastest Fourier Transform in the West.

http://www.fftw.org/

Provides simple functions to compute a transform without destroying existing data.
	x := fftw.NewArray(100)
	// ...
	xhat := fftw.FFT(x)
	x = fftw.IFFT(xhat)
Beware: Scaling is the same as in FFTW, so that computing forward and then inverse transforms scales the original input by the length of the sequence.

Use a Plan explicitly to recycle memory and to do in-place transforms.
	x := fftw.NewArray(100)
	// ...
	fftw.MakePlan1(x, x, fftw.Forward, fftw.Estimate).Execute()

It's possible to use FFTW with memory not allocated by fftw.NewArrayX().
	x := make([]complex128, 100)
	// ...

	xhat := fftw.FFT(&fftw.Array{x})

	// or in-place
	arr := &fftw.Array{x}
	fftw.MakePlan1(arr, arr, fftw.Forward, fftw.Estimate).Execute()
*/
package fftw
