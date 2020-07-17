package main

import (
	"testing"
)

func Testデッキが正しく作られているか(t *testing.T) {
	expectedLength := 52
	expectedFirst := Card{Suit{"spade"}, Rank{1}, nil}
	expectedLast := Card{Suit{"heart"}, Rank{13}, nil}

	deck := CreateDeck()
	last := expectedLength - 1

	if len(deck.cards) != expectedLength {
		t.Errorf("expected length == %d, actual %d", expectedLength, len(deck.cards))
	}

	if deck.cards[0].rank != expectedFirst.rank || deck.cards[0].suit != expectedFirst.suit {
		t.Errorf("expected first card == %v, actual %v", expectedFirst, deck.cards[0])
	}

	if deck.cards[last].rank != expectedLast.rank || deck.cards[last].suit != expectedLast.suit {
		t.Errorf("expected last card == %v, actual %v", expectedLast, deck.cards[last])
	}
}

func Test二枚のカードが返ってきているか(t *testing.T) {
	d := CreateDeck()
	//ようは一番最後から２つ分と等しくなっててほしい
	last := d.cards[len(d.cards)-1]
	last2 := d.cards[len(d.cards)-2]
	cards := d.Distribute()

	if cards[0] != last || cards[1] != last2 {
		t.Errorf("expected %v and %v, actual %v and %v", last, last2, cards[0], cards[1])
	}
}
