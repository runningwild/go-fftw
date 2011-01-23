package main

import "fftw"

func main() {
  var d fftw.Plan
  print(&d)
  x := fftw.Alloc(10)
  print(x)
}
