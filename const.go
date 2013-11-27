package fftw

// #include <fftw3.h>
import "C"

type Direction int

const (
	Forward  = Direction(C.FFTW_FORWARD)
	Backward = Direction(C.FFTW_BACKWARD)
)

type Flag uint

const (
	Estimate = Flag(C.FFTW_ESTIMATE)
	Measure  = Flag(C.FFTW_MEASURE)
)
