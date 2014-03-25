package fftw

var DefaultFlag = Estimate

// Computes the DFT.
// Allocates memory in which to return the result.
func FFT(src *Array) *Array {
	dst := NewArray(src.Len())
	fftTo(dst, src, Forward, DefaultFlag)
	return dst
}

// Computes the inverse DFT.
// Allocates memory in which to return the result.
func IFFT(src *Array) *Array {
	dst := NewArray(src.Len())
	fftTo(dst, src, Backward, DefaultFlag)
	return dst
}

func fftTo(dst, src *Array, dir Direction, flag Flag) {
	NewPlan(src, dst, dir, flag).Execute().Destroy()
}

// 2D version of FFT.
func FFT2(src *Array2) *Array2 {
	return fft2(src, Forward)
}

// 2D version of IFFT.
func IFFT2(src *Array2) *Array2 {
	return fft2(src, Backward)
}

// Allocates memory.
func fft2(src *Array2, dir Direction) *Array2 {
	n0, n1 := src.Dims()
	dst := NewArray2(n0, n1)
	fft2To(dst, src, dir, DefaultFlag)
	return dst
}

func fft2To(dst, src *Array2, dir Direction, flag Flag) {
	NewPlan2(src, dst, dir, flag).Execute().Destroy()
}

// 3D version of FFT.
func FFT3(src *Array3) *Array3 {
	return fft3(src, Forward)
}

// 3D version of IFFT.
func IFFT3(src *Array3) *Array3 {
	return fft3(src, Backward)
}

// Allocates memory.
func fft3(src *Array3, dir Direction) *Array3 {
	n0, n1, n2 := src.Dims()
	dst := NewArray3(n0, n1, n2)
	fft3To(dst, src, dir, DefaultFlag)
	return dst
}

func fft3To(dst, src *Array3, dir Direction, flag Flag) {
	NewPlan3(src, dst, dir, flag).Execute().Destroy()
}
