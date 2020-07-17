package main

import "fmt"

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

//Card カード一枚を指す型
type Card struct {
	suit Suit
	rank Rank
}

func (c Card) String() string {
	return fmt.Sprintf("Suit: %v Rank: %v", c.suit, c.rank)
}
