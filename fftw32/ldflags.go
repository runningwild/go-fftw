package fftw32

// #cgo CFLAGS: -I/usr/local/include
// #cgo darwin CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/usr/local/lib -lfftw3f -lm
// #cgo darwin LDFLAGS: -L/opt/homebrew/lib
import "C"
