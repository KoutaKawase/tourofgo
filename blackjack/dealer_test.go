package main

import "testing"

func Test正しいディーラーの座標が返ってきているか(t *testing.T) {
	// 	１つめはdefaultX(500)
	// 2つめはdefaultX(500) - (cardwidth(121) + marginRight(40)) で 339
	// 3つめは500 - (121 + 40) * 2 で 178
	// 4つめは500 - (121 + 40) * 3 で 17
	cases := []struct {
		index int
		wantX float64
	}{
		{0, 500},
		{1, 339},
		{2, 178},
		{3, 17},
	}

	for _, c := range cases {
		gotX := CalcDealerX(c.index)

		if gotX != c.wantX {
			t.Errorf("GetDealerPoint(%d) ==  %f, but actual Y = %f", c.index, c.wantX, gotX)
		}
	}
}
