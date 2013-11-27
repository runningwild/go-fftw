package fftw

// #include <fftw3.h>
import "C"

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func allocBytes(n int) unsafe.Pointer {
	// Try to allocate memory.
	buffer, err := C.fftw_malloc(C.size_t(n))
	if err != nil {
		// If malloc failed, trigger garbage collector and try again.
		runtime.GC()
		buffer, err = C.fftw_malloc(C.size_t(n))
		if err != nil {
			// If it still failed, then panic.
			panic(fmt.Errorf("could not malloc %d bytes: %s", n, err.Error()))
		}
	}
	return buffer
}

func allocCmplx(n int) []complex128 {
	buffer := allocBytes(16 * n)
	// Create a slice header for the memory.
	var elems []complex128
	header := (*reflect.SliceHeader)(unsafe.Pointer(&elems))
	*header = reflect.SliceHeader{uintptr(buffer), n, n}

	// Initialize to zero.
	for i := range elems {
		elems[i] = 0
	}
	return elems
}
