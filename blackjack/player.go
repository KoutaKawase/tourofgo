package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

const (
	//DefaultY カード一枚目のY座標　Y座標は全ての手札に固定
	DefaultY = 400
	//DefaultX カード一枚目のX座標
	DefaultX = 100
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
	for i, card := range p.hands {
		x := CalcPlayerX(i)
		card.Show(x, DefaultY, s)
	}
}

//CalcPlayerX プレイヤー手札の座標を出す関数。indexはプレイヤー手札の何枚目か。何枚目かから座標を算出する
func CalcPlayerX(index int) float64 {
	x := DefaultX + (CardWidth+MarginRight)*index
	return float64(x)
}
