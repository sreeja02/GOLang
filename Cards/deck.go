package main

import ( "fmt"
"io/ioutil"
 "strings"
 "os"
 "math/rand"
 "time"
)

type deck []string // This is of type 'deck'. It is a slice of strings

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades", "Diamonds", "Hearts","Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {// i and j were not used hence underscore
			cards = append(cards, value + "of" + suit)
		}
	}
	return cards
}

func (d deck) print() { //Any variable of type deck and access print
	for i, card := range d {
		fmt.Println(i,card) //The variable is reference to the cards from main.go. We can write func(cards deck). Similar to this in JS.
	}
}


func deal(d deck, handSize int) (deck,deck){
	return d[:handSize] , d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err!=nil {
		fmt.Println("Error",err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	//seed value should be changed to generate random numbers
	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}