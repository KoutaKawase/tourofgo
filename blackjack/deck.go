package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"path"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	assetsPath = "./assets/cards"
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

//Shuffle デッキの中のカードをシャッフルする
func (d *Deck) Shuffle() {
	deck := d.cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

//Distribute 最初に配るときに呼ばれる。呼ばれる度２枚のカードをポップして返す
func (d *Deck) Distribute() []Card {
	var cards []Card

	for i := 0; i < 2; i++ {
		last, popedDeck := d.cards[len(d.cards)-1], d.cards[:len(d.cards)-1]
		cards = append(cards, last)
		d.cards = popedDeck
	}

	return cards
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
