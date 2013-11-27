package fftw

func dims2(x [][]complex128) (n0, n1 int) {
	n0 = len(x)
	if n0 == 0 {
		return
	}
	n1 = len(x[0])
	return
}

func dims3(x [][][]complex128) (n0, n1, n2 int) {
	n0 = len(x)
	if n0 == 0 {
		return
	}
	n1 = len(x[0])
	if n1 == 0 {
		return
	}
	n2 = len(x[0][0])
	return
}
