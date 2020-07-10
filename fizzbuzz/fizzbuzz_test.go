package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{15, "FizzBuzz"},
		{3, "Fizz"},
		{5, "Buzz"},
		{1, "1"},
		{4, "4"},
	}
	for _, c := range cases {
		got := FizzBuzz(c.in)
		if got != c.want {
			t.Errorf("FizzBuzz(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
