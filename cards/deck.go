package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, value := range cardValues {
		for _, suit := range cardsSuits {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(fmt.Sprintf("Card %d: %s", i+1, card))
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)	
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) 
	if err != nil {
		log.Fatal(err)	
	}
	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle ()  {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
