/*
package fftw32 is a cgo wrapper around the Fastest Fourier Transform in the West.

http://www.fftw.org/

Provides simple functions to compute a transform without destroying existing data.
	x := fftw.NewArray(100)
	// ...
	xhat := fftw.FFT(x)
	x = fftw.IFFT(xhat)
Beware: Scaling is the same as in FFTW, so that computing forward and then inverse transforms scales the original input by the length of the sequence.

Use a Plan explicitly to recycle memory and to do in-place transforms.
Always remember to destroy a plan.
	p := fftw.NewPlan(x, x, fftw.Forward, fftw.Estimate)
	defer p.Destroy()
	p.Execute()
Execute returns the plan to permit a chained call.
	fftw.NewPlan(x, x, fftw.Forward, fftw.Estimate).Execute().Destroy()

It's possible to use FFTW with memory not allocated by fftw.NewArrayX().
	x := make([]complex64, 100)
	// ...

	xhat := fftw.FFT(&fftw.Array{x})

	// or in-place
	arr := &fftw.Array{x}
	fftw.NewPlan(arr, arr, fftw.Forward, fftw.Estimate).Execute().Destroy()

*/
package fftw32
