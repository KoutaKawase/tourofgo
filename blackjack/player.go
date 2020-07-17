package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

const (
	x1 = 100
	y1 = 300
	x2 = 240
	y2 = 400
)

//Player あなた自身が操作するもの
type Player struct {
	hands []Card
}

func (p Player) String() string {
	return fmt.Sprintf("Hands: %v", p.hands)
}

//Show 画面に画像としてプレイヤー手札を表示する関数
func (p *Player) Show(s *ebiten.Image) {
	for _, card := range p.hands {
		card.Show(x1, y1, s)
	}
}
