/*
Cgo wrapper for the Fastest Fourier Transform in the West.

http://www.fftw.org/

Provides simple functions to compute a transform without destroying existing data.
	x := fftw.NewArray(100)
	// ...
	xhat := fftw.FFT(x)
	x = fftw.IFFT(xhat)
Beware: Scaling is the same as in FFTW, so that computing forward and then inverse transforms scales the original input by the length of the sequence.

Memory is allocated using fftw_malloc and automatically freed by the garbage collector using fftw_free.

Use a Plan explicity to recycle memory and to do in-place transforms.
	x := fftw.NewArray(100)
	// ...
	MakePlan1(x, x, fftw.Forward, fftw.Estimate).Execute()
*/
package fftw
