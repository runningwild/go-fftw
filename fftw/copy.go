package fftw

import "fmt"

func CopySlice2(dst *Array2, src [][]complex128) error {
	m0, m1, err := dims2(src)
	if err != nil {
		return err
	}
	n0, n1 := dst.Dims()
	if m0 != n0 || m1 != n1 {
		return fmt.Errorf("dimensions differ: dst (%d,%d), src (%d,%d)", n0, n1, m0, m1)
	}
	d := dst.Slice()
	for i, s := range src {
		copy(d[i], s)
	}
	return nil
}

func CopySlice3(dst *Array3, src [][][]complex128) error {
	m0, m1, m2, err := dims3(src)
	if err != nil {
		return err
	}
	n0, n1, n2 := dst.Dims()
	if m0 != n0 || m1 != n1 || m2 != n2 {
		return fmt.Errorf("dimensions differ: dst (%d,%d), src (%d,%d)", n0, n1, m0, m1)
	}
	d := dst.Slice()
	for i, si := range src {
		di := d[i]
		for j, sij := range si {
			copy(di[j], sij)
		}
	}
	return nil
}

func dims2(x [][]complex128) (int, int, error) {
	if len(x) == 0 {
		return 0, 0, nil
	}
	n0 := len(x)
	n1 := len(x[0])
	for _, xi := range x {
		if len(xi) != n1 {
			err := fmt.Errorf("jagged: found (%d,%d) then (,%d)", n0, n1, len(xi))
			return 0, 0, err
		}
	}
	return n0, n1, nil
}

func dims3(x [][][]complex128) (int, int, int, error) {
	if len(x) == 0 {
		return 0, 0, 0, nil
	}
	n0 := len(x)
	n1, n2, err := dims2(x[0])
	if err != nil {
		return 0, 0, 0, err
	}
	for _, xi := range x {
		if len(xi) != n1 {
			err := fmt.Errorf("jagged: found (%d,%d,%d) then (,%d,...)", n0, n1, n2, len(xi))
			return 0, 0, 0, err
		}
		for _, xij := range xi {
			if len(xij) != n2 {
				err := fmt.Errorf("jagged: found (%d,%d,%d) then (,,%d)", n0, n1, n2, len(xij))
				return 0, 0, 0, err
			}
		}
	}
	return n0, n1, n2, nil
}
