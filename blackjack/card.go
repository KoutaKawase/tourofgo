package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

//Suit エースやクローバーなどを表す
type Suit struct{ value string }

func (s Suit) String() string {
	return s.value
}

//Rank A,1,2...J,Q,Kを表す
type Rank struct{ value int }

func (r Rank) String() string {
	return fmt.Sprintf("%d", r.value)
}

var (
	spade   = Suit{"spade"}
	heart   = Suit{"club"}
	diamond = Suit{"diamond"}
	club    = Suit{"heart"}
)

var (
	ace    = Rank{1}
	deuce  = Rank{2}
	trey   = Rank{3}
	cater  = Rank{4}
	cinque = Rank{5}
	sice   = Rank{6}
	seven  = Rank{7}
	eight  = Rank{8}
	nine   = Rank{9}
	ten    = Rank{10}
	jack   = Rank{11}
	queen  = Rank{12}
	king   = Rank{13}
)

//Suits デッキ作成時に使うsuit達
var Suits []Suit = []Suit{spade, heart, diamond, club}

//Ranks デッキ作成時に使うRank達
var Ranks []Rank = []Rank{ace, deuce, trey, cater, cinque, sice, seven, eight, nine, ten, jack, queen, king}

//CardWidth カード一つ分の幅
var CardWidth int = 121

//CardHeight カード一つ分の高さ
var CardHeight int = 171

//MarginRight カード間の隙間
var MarginRight int = 40

//Card カード一枚を指す型
type Card struct {
	suit    Suit
	rank    Rank
	surface *ebiten.Image
	back    *ebiten.Image
}

func (c Card) String() string {
	return fmt.Sprintf("Suit: %v, Rank: %v", c.suit, c.rank)
}

//Show カード一枚を画面に表示する役割
func (c *Card) Show(x, y float64, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(c.surface, op)
}

//ShowBack カードの背面を表示する
func (c *Card) ShowBack(x, y float64, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(c.back, op)
}
