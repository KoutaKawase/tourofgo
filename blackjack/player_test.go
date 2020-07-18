package main

import "testing"

func Test期待するプレイヤー手札の座標が返ってくるか(t *testing.T) {
	// 	１つめはdefaultX(100)
	// ２つめはdefaultX(100) + cardwidth(121) + marginRight(40)で231になってほしい
	// 3つめならdefaultX(100) + (cardwidth(121) + marginRight(40)) * 2で362になって欲しい
	// 4つめなら100 + (121 + 40) * 3で493になって欲しい
	cases := []struct {
		index int
		wantX float64
	}{
		{0, 100},
		{1, 261},
		{2, 422},
		{3, 583},
	}

	for _, c := range cases {
		gotX := CalcPlayerX(c.index)

		if gotX != c.wantX {
			t.Errorf("GetPlayerPoint(%d) ==  %f, but actual Y = %f", c.index, c.wantX, gotX)
		}
	}
}
