package fftw

var DefaultFlag = Estimate

func FFT1D(x *Array1D, dir Direction, flag Flag) *Array1D {
	y := NewArray1D(x.Len())
	plan1D(x, y, dir, flag).Execute()
	return y
}

func FFT2D(x *Array2D, dir Direction, flag Flag) *Array2D {
	n0, n1 := x.Dims()
	y := NewArray2D(n0, n1)
	plan2D(x, y, dir, flag).Execute()
	return y
}

func FFT3D(x *Array3D, dir Direction, flag Flag) *Array3D {
	n0, n1, n2 := x.Dims()
	y := NewArray3D(n0, n1, n2)
	plan3D(x, y, dir, flag).Execute()
	return y
}
