package main

import (
	"bytes"
	"fmt"
)

//Deck カード５２枚を保持する型
type Deck struct {
	cards []Card
}

func (d Deck) String() string {
	var buffer bytes.Buffer
	for _, card := range d.cards {
		buffer.WriteString(fmt.Sprintf("%v\n", card))
	}
	return buffer.String()
}

//CreateDeck 計52枚のトランプデッキを作成します。一枚一枚はCard構造体で表されます
func CreateDeck() Deck {
	var cards []Card

	for _, suit := range Suits {
		for _, rank := range Ranks {
			cards = append(cards, Card{suit, rank})
		}
	}
	return Deck{cards}
}
