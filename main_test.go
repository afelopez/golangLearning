package test

import "testing"

// func TestSum(t *testing.T) {
// 	total := Sum(5, 5)
// 	if total != 10 {
// 		t.Error("Suma was incorrect, got ", total, "expected ", 10)
// 	}
// }

func TestSum(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc    string
		a, b, n int
	}{
		{
			desc: "Tetsing 1+2 = 3",
			a:    1, b: 2, n: 3,
		},
		{
			desc: "Tetsing 2+2 = 4",
			a:    2, b: 2, n: 4,
		},
		{
			desc: "Tetsing 25+26 = 51",
			a:    25, b: 26, n: 51,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			total := Sum(tC.a, tC.b)
			if total != tC.n {
				t.Error("Sum was incorrect, got", total, "expected", tC.n)
			}
		})
	}
}
