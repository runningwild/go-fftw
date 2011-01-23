include $(GOROOT)/src/Make.inc

TARG=fftw

CGOFILES=fftw.go

CGO_LDFLAGS=-lfftw3 -lm -L/usr/lib/

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
