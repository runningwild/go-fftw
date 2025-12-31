# Go FFTW Bindings

Go bindings for FFTW (Fastest Fourier Transform in the West).

- FFTW homepage: http://www.fftw.org/
- FFTW docs: http://www.fftw.org/fftw3_doc/

## Requirements

- FFTW built as a shared library (`--enable-shared`).
- cgo enabled (this package uses FFTW via cgo).

### Install FFTW

```bash
./configure --enable-shared
make
make install
```

On macOS (Homebrew), FFTW is typically installed under `/opt/homebrew`.

## Install

```bash
go get github.com/meko-christian/go-fftw
```

## Packages

- `fftw`: double-precision (`fftw3`) bindings.
- `fftw32`: single-precision (`fftw3f`) bindings.

## Usage

```go
data := fftw.NewArray(64) // Similar to make([]complex128, 64)
forward := fftw.NewPlan(data, data, fftw.Forward, fftw.Estimate)
backward := fftw.NewPlan(data, data, fftw.Backward, fftw.Estimate)
defer forward.Destroy()  // Free FFTW plan resources
defer backward.Destroy() // when no longer needed.

// ... fill in data with something interesting
forward.Execute() // Frequency domain
// ... do something interesting with data
backward.Execute() // Back to time domain (scaled by len(data))
```

## Notes

- These bindings do not mirror FFTW’s C API exactly. For example, array sizes are inferred.
- `fftw.Measure` may overwrite input buffers during planning. Wrapper helpers use `fftw.Estimate`.

## License

See `LICENSE`.
