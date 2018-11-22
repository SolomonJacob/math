package math

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		arg1 int
		arg2 int
		want int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{3, 4, 7},
		{7, 8, 15},
		{0, 0, 0},
		{-3, 4, 1},
		{12, -6, 6},
		{15, -15, 0},
		{0, -1, -1},
		{-32, 0, -32},
	}

	for _, c := range cases {
		got := Add(c.arg1, c.arg2)
		want := c.want
		if got != want {
			t.Errorf("Add(%d, %d) = %d, want %d", c.arg1, c.arg2, got, c.want)
		}
	}

}

func TestSub(t *testing.T) {
	cases := []struct {
		arg1 int
		arg2 int
		want int
	}{

		{3, 2, 1},
		{9, 5, 4},
		{-3, -3, 0},
		{-3, 3, -6},
		{1, 0, 1},
		{0, -3, 3},
	}

	for _, c := range cases {
		got := Sub(c.arg1, c.arg2)
		want := c.want
		if got != want {
			t.Errorf("Sub(%d, %d) = %d, want %d", c.arg1, c.arg2, got, c.want)
		}
	}

}

func TestMult(t *testing.T) {
	cases := []struct {
		arg1 int
		arg2 int
		want int
	}{
		{3, 2, 6},
		{6, 1, 6},
		{-6, 2, -12},
		{-3, -3, 9},
		{12, 12, 144},
	}

	for _, c := range cases {
		got := Mult(c.arg1, c.arg2)
		want := c.want
		if got != want {
			t.Errorf("Mult(%d, %d) = %d, want %d", c.arg1, c.arg2, got, c.want)
		}
	}

}

func TestDivide(t *testing.T) {
	cases := []struct {
		arg1 int
		arg2 int
		want int
	}{
		{6, 2, 3},
		{10, 5, 2},
		{12, -6, -2},
		{-12, 2, -6},
	}

	for _, c := range cases {
		got := Divide(c.arg1, c.arg2)
		want := c.want
		if got != want {
			t.Errorf("Divide(%d, %d) = %d, want %d", c.arg1, c.arg2, got, c.want)
		}
	}

}
