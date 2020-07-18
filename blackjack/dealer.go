package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

//Dealer NPC
type Dealer struct {
	hands []Card
}

func (d Dealer) String() string {
	return fmt.Sprintf("Hands: %v", d.hands)
}

//Show ディーラーの手札を表示する
func (d *Dealer) Show(s *ebiten.Image) {}
