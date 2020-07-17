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
