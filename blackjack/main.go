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
	deck   Deck
	player Player
	dealer Dealer
}

//Update ゲームループ的な
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

//Draw 全てのフレームで呼ばれる描画関数
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(boardColor)

	g.player.Show(screen)
	g.dealer.Show(screen)
}

//Layout ゲーム画面サイズ
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) init() {
	g.deck.Shuffle()
	g.player = Player{g.deck.Distribute()}
	g.dealer = Dealer{g.deck.Distribute()}
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Blackjack")

	deck := CreateDeck()
	player := Player{}
	dealer := Dealer{}
	game := Game{deck, player, dealer}
	game.init()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
