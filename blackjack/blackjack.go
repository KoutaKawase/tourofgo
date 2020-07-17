package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth  = 740
	windowHeight = 580
)

var (
	boardColor color.RGBA = color.RGBA{30, 156, 86, 255}
)

//Game ebiten.Gameインターフェースを実装
type Game struct {
	deck Deck
}

//Update ゲームループ的な
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

//Draw 全てのフレームで呼ばれる描画関数
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(boardColor)
}

//Layout ゲーム画面サイズ
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Blackjack")

	deck := CreateDeck()
	game := Game{deck}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
