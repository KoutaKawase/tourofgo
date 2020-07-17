package main

import "fmt"

//Dealer NPC
type Dealer struct {
	hands []Card
}

func (d Dealer) String() string {
	return fmt.Sprintf("Hands: %v", d.hands)
}
