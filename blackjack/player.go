package main

import "fmt"

//Player あなた自身が操作するもの
type Player struct {
	hands []Card
}

func (p Player) String() string {
	return fmt.Sprintf("Hands: %v", p.hands)
}
