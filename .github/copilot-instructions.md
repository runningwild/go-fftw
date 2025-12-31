# Copilot instructions for go-fftw

## Big picture

- This repo provides Go (cgo) bindings to FFTW 3:
  - `fftw/`: double-precision complex FFTs (`complex128`) linking `-lfftw3`.
  - `fftw32/`: single-precision complex FFTs (`complex64`) linking `-lfftw3f`.
- Both packages intentionally mirror each other’s API shape; if you change one, usually change the other.

## Local build / test workflow

- This repo has no `go.mod` in its root; run in GOPATH mode or from a parent module.
- First-time test setup (module mode): `go get github.com/smartystreets/goconvey/convey`
- Run all tests: `go test ./...`
- Run only one package: `go test ./fftw` or `go test ./fftw32`
- Focus a single test: `go test -run TestName ./fftw`
- Tests depend on `github.com/smartystreets/goconvey/convey`.

## FFTW dependency + CGO linking

- FFTW must be installed as a shared library (see README for `./configure --enable-shared`).
- You need FFTW headers available at build time (`fftw3.h`); on Linux this is commonly provided by a `libfftw3-dev`-style package.
- Default cgo flags are hard-coded to `/usr/local`:
  - `fftw/ldflags.go` (`-I/usr/local/include`, `-L/usr/local/lib -lfftw3 -lm`)
  - `fftw32/ldflags.go` (`-I/usr/local/include`, `-L/usr/local/lib -lfftw3f -lm`)
- If your FFTW is elsewhere (common on Linux), prefer setting `CGO_CFLAGS` / `CGO_LDFLAGS` when building/testing, or update the two `ldflags.go` files.

## API / codebase conventions

- High-level helpers (`FFT`, `IFFT`, `FFT2`, `FFT3`, `FFTN`, and `XxxTo` variants) allocate destination arrays and use `Estimate` planning by default.
- Advanced usage is via `Plan`:
  - Create with `NewPlan*`, execute with `Execute()`, and always free native resources with `Destroy()` (often via `defer`).
  - Plan creation and destruction are serialized via a package-level mutex (`createDestroyMu`) per FFTW threading guidance.
- Flags:
  - `Estimate` is safe for preserving input.
  - `Measure` may overwrite input during plan creation; avoid using it in helper APIs unless you explicitly document the mutation.

## Data types and memory

- Arrays are thin wrappers around Go slices:
  - `fftw.Array{Elems: []complex128}` and `fftw32.Array{Elems: []complex64}` (you can also use the single-field literal form `&fftw.Array{x}`).
  - 2D/3D arrays store a flat `Elems` slice plus dimension metadata; see `Slice()` in `fftw/array.go` for row-major views.
- Avoid zero-length arrays when calling into FFTW: the internal `ptr()` helpers take `&Elems[0]`.
