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
func (d *Dealer) Show(s *ebiten.Image) {
	for _, card := range d.hands {
		card.Show(DealerX, DealerY, s)
	}
}
