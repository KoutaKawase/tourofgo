package main

//TODO: 勝者が存在するかテストかく
import "testing"

func Test横一列が何らかのマークで埋まってたらそのマークを勝者となる(t *testing.T) {
	current := Players["circle"]
	board := [][]string{
		[]string{"Y1", " _", " _", " _"},
		[]string{"Y2", " 0", " 0", " 0"},
		[]string{"Y3", " _", " _", " _"},
		[]string{" ", " X1", "X2", "X3"},
	}
	board2 := [][]string{
		[]string{"Y1", " _", " _", " _"},
		[]string{"Y2", " _", " _", " _"},
		[]string{"Y3", " 0", " 0", " 0"},
		[]string{" ", " X1", "X2", "X3"},
	}

	cases := []struct {
		board   [][]string
		current string
		want    string
	}{
		{board, current, current},
		{board2, current, current},
	}

	for _, c := range cases {
		got := GetWinner(c.board, c.current)
		if got != c.want {
			t.Errorf("GetWinner(%q, %q) == %q, want %q", c.board, c.current, got, c.want)
		}
	}
}

func Testどれか縦一列がcurrentで埋まってたら勝者としてcurrentを返す(t *testing.T) {
	current := Players["circle"]
	board := [][]string{
		[]string{"Y1", " _", " 0", " _"},
		[]string{"Y2", " _", " 0", " _"},
		[]string{"Y3", " _", " 0", " _"},
		[]string{" ", " X1", "X2", "X3"},
	}
	board2 := [][]string{
		[]string{"Y1", " _", " _", " 0"},
		[]string{"Y2", " _", " _", " 0"},
		[]string{"Y3", " _", " _", " 0"},
		[]string{" ", " X1", "X2", "X3"},
	}

	cases := []struct {
		board   [][]string
		current string
		want    string
	}{
		{board, current, current},
		{board2, current, current},
	}

	for _, c := range cases {
		got := GetWinner(c.board, c.current)
		if got != c.want {
			t.Errorf("GetWinner(%q, %q) == %q, want %q", c.board, c.current, got, c.want)
		}
	}
}

func Test斜めに揃っていたらwinnerとしてcurrentを返す(t *testing.T) {
	current := Players["circle"]
	board := [][]string{
		[]string{"Y1", " 0", " _", " _"},
		[]string{"Y2", " _", " 0", " _"},
		[]string{"Y3", " _", " _", " 0"},
		[]string{" ", " X1", "X2", "X3"},
	}
	board2 := [][]string{
		[]string{"Y1", " _", " _", " 0"},
		[]string{"Y2", " _", " 0", " _"},
		[]string{"Y3", " 0", " _", " _"},
		[]string{" ", " X1", "X2", "X3"},
	}

	cases := []struct {
		board   [][]string
		current string
		want    string
	}{
		{board, current, current},
		{board2, current, current},
	}

	for _, c := range cases {
		got := GetWinner(c.board, c.current)
		if got != c.want {
			t.Errorf("GetWinner(%q, %q) == %q, want %q", c.board, c.current, got, c.want)
		}
	}
}
