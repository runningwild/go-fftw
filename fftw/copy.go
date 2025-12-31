package fftw

import (
	"errors"
	"fmt"
)

var (
	ErrDimensionsMismatch = errors.New("dimensions mismatch")
	ErrJaggedArray        = errors.New("jagged array")
)

func CopySlice2(dst *Array2, src [][]complex128) error {
	srcDim0, srcDim1, err := dims2(src)
	if err != nil {
		return err
	}

	dstDim0, dstDim1 := dst.Dims()
	if srcDim0 != dstDim0 || srcDim1 != dstDim1 {
		return fmt.Errorf("%w: dst (%d,%d), src (%d,%d)", ErrDimensionsMismatch, dstDim0, dstDim1, srcDim0, srcDim1)
	}

	d := dst.Slice()
	for i, s := range src {
		copy(d[i], s)
	}

	return nil
}

func CopySlice3(dst *Array3, src [][][]complex128) error {
	srcDim0, srcDim1, srcDim2, err := dims3(src)
	if err != nil {
		return err
	}

	dstDim0, dstDim1, dstDim2 := dst.Dims()
	if srcDim0 != dstDim0 || srcDim1 != dstDim1 || srcDim2 != dstDim2 {
		return fmt.Errorf("%w: dst (%d,%d,%d), src (%d,%d,%d)",
			ErrDimensionsMismatch, dstDim0, dstDim1, dstDim2, srcDim0, srcDim1, srcDim2)
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

	dim0 := len(x)

	dim1 := len(x[0])
	for _, xi := range x {
		if len(xi) != dim1 {
			return 0, 0, fmt.Errorf("%w: found (%d,%d) then (,%d)", ErrJaggedArray, dim0, dim1, len(xi))
		}
	}

	return dim0, dim1, nil
}

func dims3(x [][][]complex128) (int, int, int, error) {
	if len(x) == 0 {
		return 0, 0, 0, nil
	}

	dim0 := len(x)

	dim1, dim2, err := dims2(x[0])
	if err != nil {
		return 0, 0, 0, err
	}

	for _, xi := range x {
		if len(xi) != dim1 {
			return 0, 0, 0, fmt.Errorf("%w: found (%d,%d,%d) then (,%d,...)", ErrJaggedArray, dim0, dim1, dim2, len(xi))
		}

		for _, xij := range xi {
			if len(xij) != dim2 {
				return 0, 0, 0, fmt.Errorf("%w: found (%d,%d,%d) then (,,%d)", ErrJaggedArray, dim0, dim1, dim2, len(xij))
			}
		}
	}

	return dim0, dim1, dim2, nil
}
