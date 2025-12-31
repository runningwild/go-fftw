package fftw32

import "testing"

func expectPanic(t *testing.T, name string, panicFn func()) {
	t.Helper()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: expect panic", name)
		}
	}()

	panicFn()
}

func TestNewPlanGuards(t *testing.T) {
	t.Parallel()

	var nilArray *Array

	expectPanic(t, "nil input", func() {
		NewPlan(nilArray, NewArray(1), Forward, Estimate)
	})

	expectPanic(t, "nil output", func() {
		NewPlan(NewArray(1), nilArray, Forward, Estimate)
	})

	expectPanic(t, "empty input", func() {
		NewPlan(NewArray(0), NewArray(0), Forward, Estimate)
	})

	expectPanic(t, "size <= 0", func() {
		NewPlanForSize(0, Forward, Estimate)
	})
}

func TestNewPlanForSize(t *testing.T) {
	t.Parallel()

	p, in, out := NewPlanForSize(8, Forward, Estimate)
	defer p.Destroy()

	if got := p.String(); got == "" {
		t.Fatalf("expected non-empty plan string")
	}

	in.Elems[0] = 1
	for i := 1; i < len(in.Elems); i++ {
		in.Elems[i] = 0
	}

	p.Execute()

	for i := range out.Elems {
		testAlmostEqual(t, real(out.Elems[i]), 1.0)
		testAlmostEqual(t, imag(out.Elems[i]), 0.0)
	}
}

func TestNewPlan2Guards(t *testing.T) {
	t.Parallel()

	var nilArray *Array2

	expectPanic(t, "nil input", func() {
		NewPlan2(nilArray, NewArray2(1, 1), Forward, Estimate)
	})

	expectPanic(t, "nil output", func() {
		NewPlan2(NewArray2(1, 1), nilArray, Forward, Estimate)
	})

	expectPanic(t, "empty input", func() {
		NewPlan2(NewArray2(0, 1), NewArray2(0, 1), Forward, Estimate)
	})
}

func TestNewPlan3Guards(t *testing.T) {
	t.Parallel()

	var nilArray *Array3

	expectPanic(t, "nil input", func() {
		NewPlan3(nilArray, NewArray3(1, 1, 1), Forward, Estimate)
	})

	expectPanic(t, "nil output", func() {
		NewPlan3(NewArray3(1, 1, 1), nilArray, Forward, Estimate)
	})

	expectPanic(t, "empty input", func() {
		NewPlan3(NewArray3(1, 0, 1), NewArray3(1, 0, 1), Forward, Estimate)
	})
}
