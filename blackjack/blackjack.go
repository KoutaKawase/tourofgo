package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth  = 640
	windowHeight = 480
)

var (
	boardColor color.RGBA = color.RGBA{30, 156, 86, 255}
)

//Game ebiten.Gameインターフェースを実装
type Game struct{}

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
	ebiten.SetWindowTitle("Hello from ebiten!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
