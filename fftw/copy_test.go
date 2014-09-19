package fftw

import "testing"

func TestCopySlice2(t *testing.T) {
	cases := []struct {
		M, N int
		Lens []int
		Err  bool
	}{
		// Test size mismatch.
		{4, 4, []int{4, 4, 4, 4}, false},
		{3, 3, []int{4, 4, 4, 4}, true},
		{3, 4, []int{4, 4, 4, 4}, true},
		{4, 3, []int{4, 4, 4, 4}, true},
		{2, 4, []int{4, 4}, false},
		{4, 2, []int{2, 2, 2, 2}, false},
		{2, 4, []int{2, 2, 2, 2}, true},
		{4, 2, []int{4, 4}, true},
		// Test valid or invalid dimensions.
		{4, 4, []int{4, 4, 4, 4}, false},
		{3, 4, []int{4, 4, 3}, true},
		{3, 4, []int{3, 4, 4}, true},
		{0, 0, []int{}, false},
		{1, 0, []int{0}, false},
		{1, 1, []int{1}, false},
		{3, 0, []int{0, 0, 0}, false},
		{2, 1, []int{1, 0}, true},
		{2, 1, []int{0, 1}, true},
		{2, 1, []int{0, 1}, true},
	}

	for _, test := range cases {
		x := make([][]complex128, len(test.Lens))
		for i := range test.Lens {
			x[i] = make([]complex128, test.Lens[i])
			for j := range x[i] {
				x[i][j] = complex(float64((i+1)*j), 0)
			}
		}
		arr := NewArray2(test.M, test.N)
		err := CopySlice2(arr, x)
		if test.Err && err != nil {
			continue
		}
		if test.Err {
			t.Errorf("expect error: %+v", test)
			continue
		}
		if err != nil {
			t.Errorf("error: %v", err)
			continue
		}
		for i := 0; i < test.M; i++ {
			for j := 0; j < test.N; j++ {
				want := complex(float64((i+1)*j), 0)
				got := arr.At(i, j)
				if got != want {
					t.Errorf("at %d, %d: want %v, got %v", i, j, got, want)
				}
			}
		}
	}
}

func TestCopySlice3(t *testing.T) {
	cases := []struct {
		M, N, P int
		Lens    [][]int
		Err     bool
	}{
		// Test size mismatch.
		{
			4, 4, 4,
			[][]int{
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
			},
			false,
		},
		{
			3, 3, 3,
			[][]int{
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
			},
			true,
		},
		{
			3, 4, 4,
			[][]int{
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
			},
			true,
		},
		{
			4, 3, 4,
			[][]int{
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
			},
			true,
		},
		{
			4, 4, 3,
			[][]int{
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
				[]int{4, 4, 4, 4},
			},
			true,
		},
		{
			2, 3, 4,
			[][]int{[]int{4, 4, 4}, []int{4, 4, 4}},
			false,
		},
		{
			2, 4, 3,
			[][]int{[]int{4, 4, 4}, []int{4, 4, 4}},
			true,
		},
		{
			3, 2, 4,
			[][]int{[]int{4, 4, 4}, []int{4, 4, 4}},
			true,
		},
		{
			3, 4, 2,
			[][]int{[]int{4, 4, 4}, []int{4, 4, 4}},
			true,
		},
		{
			4, 2, 3,
			[][]int{[]int{4, 4, 4}, []int{4, 4, 4}},
			true,
		},
		{
			4, 3, 2,
			[][]int{[]int{4, 4, 4}, []int{4, 4, 4}},
			true,
		},
		// Singleton dimensions.
		{
			1, 3, 4,
			[][]int{[]int{4, 4, 4}},
			false,
		},
		{
			2, 1, 4,
			[][]int{[]int{4}, []int{4}},
			false,
		},
		{
			2, 3, 1,
			[][]int{[]int{1, 1, 1}, []int{1, 1, 1}},
			false,
		},
		// Test valid or invalid dimensions.
		{
			2, 3, 4,
			[][]int{[]int{3, 4, 4}, []int{4, 4, 4}},
			true,
		},
		{
			2, 3, 4,
			[][]int{[]int{4, 4, 3}, []int{4, 4, 4}},
			true,
		},
		{
			2, 3, 4,
			[][]int{[]int{4, 4, 4}, []int{3, 4, 4}},
			true,
		},
		{
			2, 3, 4,
			[][]int{[]int{4, 4, 4}, []int{4, 4, 3}},
			true,
		},
		{
			2, 3, 4,
			[][]int{[]int{4, 4, 4}, []int{4, 4}},
			true,
		},
		{
			2, 3, 4,
			[][]int{[]int{4, 4, 4}, []int{4, 4, 4, 4}},
			true,
		},
	}

	for _, test := range cases {
		x := make([][][]complex128, len(test.Lens))
		for i := range test.Lens {
			x[i] = make([][]complex128, len(test.Lens[i]))
			for j := range x[i] {
				x[i][j] = make([]complex128, test.Lens[i][j])
				for k := range x[i][j] {
					x[i][j][k] = complex(float64((((i+1)*j)+1)*k), 0)
				}
			}
		}
		arr := NewArray3(test.M, test.N, test.P)
		err := CopySlice3(arr, x)
		if test.Err && err != nil {
			continue
		}
		if test.Err {
			t.Errorf("expect error: %+v", test)
			continue
		}
		if err != nil {
			t.Errorf("error: %v", err)
			continue
		}
		for i := 0; i < test.M; i++ {
			for j := 0; j < test.N; j++ {
				for k := 0; k < test.P; k++ {
					want := complex(float64(((i+1)*j+1)*k), 0)
					got := arr.At(i, j, k)
					if got != want {
						t.Errorf("at %d, %d, %d: want %v, got %v", i, j, k, got, want)
					}
				}
			}
		}
	}
}
