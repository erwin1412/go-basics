package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("hello world")

	// arraySign([]int{2, 1})                    // 1
	// arraySign([]int{-2, 1})                   // -1
	// arraySign([]int{-1, -2, -3, -4, 3, 2, 1}) // 1

	fmt.Println(arraySign([]int{2, 1}))                    // Seharusnya output: 1
	fmt.Println(arraySign([]int{-2, 1}))                   // Seharusnya output: -1
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1})) // Seharusnya output: 1

	// isAnagram("anak", "kana")       // true
	// isAnagram("anak", "mana")       // false
	// isAnagram("anagram", "managra") // true

	fmt.Println(isAnagram("anak", "kana"))       // Seharusnya output: true
	fmt.Println(isAnagram("anak", "mana"))       // Seharusnya output: false
	fmt.Println(isAnagram("anagram", "managra")) // Seharusnya output: true

	// findTheDifference("abcd", "abcde") // 'e'
	// findTheDifference("abcd", "abced") // 'e'
	// findTheDifference("", "y")         // 'y'

	fmt.Println(string(findTheDifference("abcd", "abcde"))) // Seharusnya output: 'e'
	fmt.Println(string(findTheDifference("abcd", "abced"))) // Seharusnya output: 'e'
	fmt.Println(string(findTheDifference("", "y")))         // Seharusnya output: 'y'

	// canMakeArithmeticProgression([]int{1, 5, 3})    // true; 1, 3, 5 adalah baris aritmatik +2
	// canMakeArithmeticProgression([]int{5, 1, 9})    // true; 9, 5, 1 adalah baris aritmatik -4
	// canMakeArithmeticProgression([]int{1, 2, 4, 8}) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	fmt.Println(canMakeArithmeticProgression([]int{1, 5, 3}))
	fmt.Println(canMakeArithmeticProgression([]int{5, 1, 9}))
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4, 8}))

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// write code here
	productSign := 1
	for _, num := range nums {
		if num == 0 {
			return 0 // Jika ada elemen 0, langsung kembalikan 0
		}
		if num < 0 {
			productSign *= -1 // Ubah tanda jika elemen negatif
		}
	}
	return productSign
	// return 1 // if positive
	// return -1 // if negative
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// write code here
	if len(s) != len(t) {
		return false
	}

	// buat array untuk simpan ascii dari s dan t
	sBytes := []byte(s)
	tBytes := []byte(t)
	//disini mengurutkan ascii dari "anak" menjadi "aakn" dan "kana" menjadi "aakn" juga
	// kemudian dari "aakn" dan "aakn" akan di buat menjadi ASCII
	// kemudian di bandingkan nilai ASCII nya sama apa engga
	// sBytes setelah diurutkan: [97 97 107 110] (ASCII untuk 'a', 'a', 'k', 'n')
	// tBytes setelah diurutkan: [97 97 107 110] (ASCII untuk 'a', 'a', 'k', 'n')
	sort.Slice(sBytes, func(i, j int) bool { return sBytes[i] < sBytes[j] })
	sort.Slice(tBytes, func(i, j int) bool { return tBytes[i] < tBytes[j] })
	return string(sBytes) == string(tBytes)

}

// findTheDifference("abcd", "abcde") // 'e'
// 	findTheDifference("abcd", "abced") // 'e'
// 	findTheDifference("", "y")         // 'y'

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	// write code here
	// buat array untuk simpan ascii dari s dan t
	sBytes := []byte(s)
	tBytes := []byte(t)
	sort.Slice(sBytes, func(i, j int) bool { return sBytes[i] < sBytes[j] })
	sort.Slice(tBytes, func(i, j int) bool { return tBytes[i] < tBytes[j] })

	for i := 0; i < len(sBytes); i++ {
		// if()
		if sBytes[i] != tBytes[i] {
			return tBytes[i]
		}
	}
	return tBytes[len(tBytes)-1]
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// Urutkan array
	sort.Ints(arr)

	// Hitung selisih antara dua elemen pertama
	diff := arr[1] - arr[0]

	// Periksa apakah semua elemen memiliki selisih yang sama
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	// write code here
	symbols := 4
	numbers := 13
	d.cards = make([]Card, 0, symbols*numbers)
	for s := 0; s < symbols; s++ {
		for n := 1; n <= numbers; n++ {
			d.cards = append(d.cards, Card{symbol: s, number: n})
		}
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	// write code here
	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	// write code here
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
