package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

const (
	//DealerY ディーラーカードのデフォルトY座標 全てのディーラー手札に共通
	DealerY = 5
	//DealerX ディーラーカードのデフォルトX座標
	DealerX = 500
)

//Dealer NPC
type Dealer struct {
	hands []Card
}

func (d Dealer) String() string {
	return fmt.Sprintf("Hands: %v", d.hands)
}

//Show ディーラーの手札を表示する
func (d *Dealer) Show(s *ebiten.Image, isAllDone bool) {
	for i, card := range d.hands {
		x := CalcDealerX(i)
		//一枚だけ伏せて表示したいのでiが0の時だけ背面表示している
		if !isAllDone && i == 0 {
			card.ShowBack(x, DealerY, s)
		} else {
			card.Show(x, DealerY, s)
		}
	}
}

//CalcDealerX ディーラーのX座標を計算し返す
func CalcDealerX(index int) float64 {
	x := DealerX - (CardWidth+MarginRight)*index
	return float64(x)
}
