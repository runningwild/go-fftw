package fftw

// FFT computes the Fourier transform of src.
// It allocates memory in which to return the result.
func FFT(src *Array) *Array {
	dst := NewArray(src.Len())
	fftDir(dst, src, Forward)
	return dst
}

// IFFT computes the inverse Fourier transform of src.
// It allocates memory in which to return the result.
func IFFT(src *Array) *Array {
	dst := NewArray(src.Len())
	fftDir(dst, src, Backward)
	return dst
}

// FFTTo computes the Fourier transform of src
// and returns the result in dst.
func FFTTo(dst, src *Array) { fftDir(dst, src, Forward) }

// IFFTTo computes the inverse Fourier transform of src
// and returns the result in dst.
func IFFTTo(dst, src *Array) { fftDir(dst, src, Backward) }

func fftDir(dst, src *Array, dir Direction) {
	p := NewPlan(src, dst, dir, Estimate)
	defer p.Destroy()
	p.Execute()
}

// FFT2 computes the Fourier transform of src.
// It allocates memory in which to return the result.
func FFT2(src *Array2) *Array2 {
	dst := NewArray2(src.Dims())
	fft2Dir(dst, src, Forward)
	return dst
}

// IFFT2 computes the inverse Fourier transform of src.
// It allocates memory in which to return the result.
func IFFT2(src *Array2) *Array2 {
	dst := NewArray2(src.Dims())
	fft2Dir(dst, src, Backward)
	return dst
}

// FFT2To computes the Fourier transform of src
// and returns the result in dst.
func FFT2To(dst, src *Array2) { fft2Dir(dst, src, Forward) }

// IFFT2To computes the inverse Fourier transform of src
// and returns the result in dst.
func IFFT2To(dst, src *Array2) { fft2Dir(dst, src, Backward) }

func fft2Dir(dst, src *Array2, dir Direction) {
	p := NewPlan2(src, dst, dir, Estimate)
	defer p.Destroy()
	p.Execute()
}

// FFT3 computes the Fourier transform of src.
// It allocates memory in which to return the result.
func FFT3(src *Array3) *Array3 {
	dst := NewArray3(src.Dims())
	fft3Dir(dst, src, Forward)
	return dst
}

// IFFT3 computes the inverse Fourier transform of src.
// It allocates memory in which to return the result.
func IFFT3(src *Array3) *Array3 {
	dst := NewArray3(src.Dims())
	fft3Dir(dst, src, Backward)
	return dst
}

// FFT3To computes the Fourier transform of src
// and returns the result in dst.
func FFT3To(dst, src *Array3) { fft3Dir(dst, src, Forward) }

// IFFT3To computes the inverse Fourier transform of src
// and returns the result in dst.
func IFFT3To(dst, src *Array3) { fft3Dir(dst, src, Backward) }

func fft3Dir(dst, src *Array3, dir Direction) {
	p := NewPlan3(src, dst, dir, Estimate)
	defer p.Destroy()
	p.Execute()
}

// FFTN computes the Fourier transform of src.
// It allocates memory in which to return the result.
func FFTN(src *ArrayN) *ArrayN {
	dst := NewArrayN(src.Dims())
	fftNDir(dst, src, Forward)
	return dst
}

// IFFTN computes the inverse Fourier transform of src.
// It allocates memory in which to return the result.
func IFFTN(src *ArrayN) *ArrayN {
	dst := NewArrayN(src.Dims())
	fftNDir(dst, src, Backward)
	return dst
}

// FFTNTo computes the Fourier transform of src
// and returns the result in dst.
func FFTNTo(dst, src *ArrayN) { fftNDir(dst, src, Forward) }

// IFFTNTo computes the inverse Fourier transform of src
// and returns the result in dst.
func IFFTNTo(dst, src *ArrayN) { fftNDir(dst, src, Backward) }

func fftNDir(dst, src *ArrayN, dir Direction) {
	p := NewPlanN(src, dst, dir, Estimate)
	defer p.Destroy()
	p.Execute()
}
