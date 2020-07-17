package main

import (
	"bytes"
	"fmt"
	"log"
	"path"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	assetsPath = "./assets"
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

func getCardAsset(count int) *ebiten.Image {
	fileName := fmt.Sprintf("card%d.png", count)
	imgPath := path.Join(assetsPath, fileName)

	img, _, err := ebitenutil.NewImageFromFile(imgPath, ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}

	return img
}

//CreateDeck 計52枚のトランプデッキを作成します。一枚一枚はCard構造体で表されます
func CreateDeck() Deck {
	var cards []Card
	//カードが現在何枚目を指しているか保持する用
	count := 1

	for _, suit := range Suits {
		for _, rank := range Ranks {
			img := getCardAsset(count)
			cards = append(cards, Card{suit, rank, img})
			count++
		}
	}
	return Deck{cards}
}
