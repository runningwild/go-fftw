Go bindings for FFTW v3.2.2
Maintained by Jonathan Wills: runningwild@gmail.com
Feel free to email me patches, suggestions or bugs.

FFTW homepage: http://www.fftw.org/
Documentation for the latest version: http://www.fftw.org/fftw3_doc/

These bindings are incomplete, but should include enough functionality that you can do whatever transforms you need (perhaps not as easily as you would like, for now).  The function definitions do not mirror exactly what is written in the docs.  For example, passing arrays does not require passing the size of the arrays.

Usage:
Here is an example of doing a simple DFT with these bindings

    data := fftw.NewArray(64) // Similar to calling make([]complex128, 64)
    forward := fftw.NewPlan(data, data, fftw.Forward, fftw.Estimate)
    backward := fftw.NewPlan(data, data, fftw.Backward, fftw.Estimate)
    defer forward.Destroy()  // Make sure to clean things up, since there are allocations happening
    defer backward.Destroy() // in c-land somewhere.
    // ... fill in data with something interesting
    forward.Execute() // Transforms data, in place, to frequency domain
    // ... do something interesting with data
    backward.Execute() // Returns data, in place, to time domain

Installation:
When installing fftw you must compile it as a shared library:

    ./configure --enable-shared
    make
    make install

Once installed properly, these bindings can be installed like so:

    go get github.com/runningwild/go-fftw

