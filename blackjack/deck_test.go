package main

import "testing"

func Testデッキが正しく作られているか(t *testing.T) {
	expectedLength := 52
	expectedFirst := Card{Suit{"spade"}, Rank{1}, nil}
	expectedLast := Card{Suit{"heart"}, Rank{13}, nil}

	deck := CreateDeck()
	last := expectedLength - 1

	if len(deck.cards) != expectedLength {
		t.Errorf("expected length == %d, actual %d", expectedLength, len(deck.cards))
	}

	if deck.cards[0] != expectedFirst {
		t.Errorf("expected first card == %v, actual %v", expectedFirst, deck.cards[0])
	}

	if deck.cards[last] != expectedLast {
		t.Errorf("expected last card == %v, actual %v", expectedLast, deck.cards[last])
	}
}
