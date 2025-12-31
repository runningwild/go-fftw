# Repository Guidelines

## Project Structure & Module Organization

- `fftw/`: Go bindings for double-precision FFTW (`fftw3`), source and tests live side-by-side.
- `fftw32/`: Go bindings for single-precision FFTW (`fftw3f`), mirroring the `fftw/` API.
- Tests are colocated with sources as `*_test.go` files (for example `fftw/fftw_test.go`).
- `README.md` documents installation and usage examples.

## Build, Test, and Development Commands

- `go test ./...`: run all unit tests across both packages (requires FFTW shared libs installed).
- `go test ./fftw`: run only double-precision package tests.
- `go test ./fftw32`: run only single-precision package tests.
- `go test -run TestName ./fftw`: run a focused test by name.

## Coding Style & Naming Conventions

- Go formatting: use `gofmt` defaults (tabs for indentation, standard import grouping).
- Exported API is `PascalCase`, unexported helpers are `camelCase`.
- File names are lower snake case (for example `array_test.go`), and tests follow `TestXxx` naming.
- Keep APIs symmetric across `fftw/` and `fftw32/` where possible.

## Testing Guidelines

- Tests use Go’s `testing` package plus `github.com/smartystreets/goconvey/convey`.
- No explicit coverage targets are defined; focus on correctness for both real and complex transforms.
- Ensure FFTW libraries are discoverable at runtime (for example, `/usr/local/lib`).

## Dependencies & Configuration Tips

- FFTW must be built as a shared library; example:
  - `./configure --enable-shared && make && make install`
- cgo flags are defined in `fftw/ldflags.go` and `fftw32/ldflags.go`; adjust include/lib paths if FFTW is installed elsewhere.

## Commit & Pull Request Guidelines

- Commit messages are short, imperative, and specific (for example “Add N-dim FFT”).
- PRs should include: summary, rationale, FFTW version tested, and exact `go test` commands run.
- If a change touches both packages, update both `fftw/` and `fftw32/` in the same PR.
