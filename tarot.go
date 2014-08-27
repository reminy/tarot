package main

import (
	"fmt"
	"time"
	"os"
	"math/rand"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]
 	fmt.Println(argsWithProg)	
 	fmt.Println(argsWithoutProg)	
 	fmt.Println(arg)	


	riderDeck := make(map[int]string)

	riderDeck = map[int]string{
	1: "The Fool ",
	100: "The Fool, Reversed ",
	2: "The Magician ",
	200: "The Magician, Reversed ",
	3: "The High Priestess ",
	300: "The High Priestess, Reversed ",
	4: "The Empress ",
	400: "The Empress, Reversed ",
	5: "The Emperor ",
	500: "The Emperor, Reversed ",
	6: "The Hierophant ",
	600: "The Hierophant, Reversed ",
	7: "The Lovers ",
	700: "The Lovers, Reversed ",
	8: "The Chariot ",
	800: "The Chariot, Reversed ",
	9: "Strength ",
	900: "Strength, Reversed ",
	10: "The Hermit ",
	1000: "The Hermit, Reversed ",
	11: "Wheel of Fortune ",
	1100: "Wheel of Fortune, Reversed ",
	12: "Justice ",
	1200: "Justice, Reversed ",
	13: "The Hanged Man ",
	1300: "The Hanged Man, Reversed ",
	14: "Death ",
	1400: "Death, Reversed ",
	15: "Temperance ",
	1500: "Temperance, Reversed ",
	16: "The Devil ",
	1600: "The Devil, Reversed ",
	17: "The Tower ",
	1700: "The Tower, Reversed ",
	18: "The Star ",
	1800: "The Star, Reversed ",
	19: "The Moon ",
	1900: "The Moon, Reversed ",
	20: "The Sun ",
	2000: "The Sun, Reversed ",
	21: "Judgment ",
	2100: "Judgement, Reversed ",
	22: "The World ",
	2200: "The World, Reversed ",
	23: "Ace of Wands ",
	2300: "Ace of Wands, Reversed ",
	24: "Two of Wands ",
	2400: "Two of Wands, Reversed ",
	25: "Three of Wands ",
	2500: "Three of Wands, Reversed ",
	26: "Four of Wands ",
	2600: "Four of Wands, Reversed ",
	27: "Five of Wands ",
	2700: "Five of Wands, Reversed ",
	28: "Six of Wands ",
	2800: "Six of Wands, Reversed ",
	29: "Seven of Wands ",
	2900: "Seven of Wands, Reversed ",
	30: "Eight of Wands ",
	3000: "Eight of Wands, Reversed ",
	31: "Nine of Wands ",
	3100: "Nine of Wands, Reversed ",
	32: "Ten of Wands ",
	3200: "Ten of Wands, Reversed ",
	33: "Page of Wands ",
	3300: "Page of Wands, Reversed ",
	34: "Knight of Wands ",
	3400: "Knight of Wands, Reversed ",
	35: "Queen of Wands ",
	3500: "Queen of Wands, Reversed ",
	36: "King of Wands ",
	3600: "King of Wands, Reversed ",
	37: "Ace of Pentacles ",
	3700: "Ace of Pentacles, Reversed ",
	38: "Two of Pentacles ",
	3800: "Two of Pentacles, Reversed ",
	39: "Three of Pentacles ",
	3900: "Three of Pentacles, Reversed ",
	40: "Four of Pentacles ",
	4000: "Four of Pentacles, Reversed ",
	41: "Five of Pentacles ",
	4100: "Five of Pentacles, Reversed ",
	42: "Six of Pentacles ",
	4200: "Six of Pentacles, Reversed ",
	43: "Seven of Pentacles ",
	4300: "Seven of Pentacles, Reversed ",
	44: "Eight of Pentacles ",
	4400: "Eight of Pentacles, Reversed ",
	45: "Nine of Pentacles ",
	4500: "Nine of Pentacles, Reversed ",
	46: "Ten of Pentacles ",
	4600: "Ten of Pentacles, Reversed ",
	47: "Page of Pentacles ",
	4700: "Page of Pentacles, Reversed ",
	48: "Knight of Pentacles ",
	4800: "Knight of Pentacles, Reversed ",
	49: "Queen of Pentacles ",
	4900: "Queen of Pentacles, Reversed ",
	50: "King of Pentacles ",
	5000: "King of Pentacles, Reversed ",
	51: "Ace of Cups ",
	5100: "Ace of Cups, Reversed ",
	52: "Two of Cups ",
	5200: "Two of Cups, Reversed ",
	53: "Three of Cups ",
	5300: "Three of Cups, Reversed ",
	54: "Four of Cups ",
	5400: "Four of Cups, Reversed ",
	55: "Five of Cups ",
	5500: "Five of Cups, Reversed ",
	56: "Six of Cups ",
	5600: "Six of Cups, Reversed ",
	57: "Seven of Cups ",
	5700: "Seven of Cups, Reversed ",
	58: "Eight of Cups ",
	5800: "Eight of Cups, Reversed ",
	59: "Nine of Cups ",
	5900: "Nine of Cups, Reversed ",
	60: "Ten of Cups ",
	6000: "Ten of Cups, Reversed ",
	61: "Page of Cups ",
	6100: "Page Cups, Reversed ",
	62: "Knight of Cups ",
	6200: "Knight of Cups, Reversed ",
	63: "Queen of Cups ",
	6300: "Queen of Cups, Reversed ",
	64: "King of Cups ",
	6400: "King of Cups, Reversed ",
	65: "Ace of Swords ",
	6500: "Ace of Swords, Reversed ",
	66: "Two of Swords ",
	6600: "Two of Swords, Reversed ",
	67: "Three of Swords ",
	6700: "Three of Swords, Reversed ",
	68: "Four of Swords ",
	6800: "Four of Swords, Reversed ",
	69: "Five of Swords ",
	6900: "Five of Swords, Reversed ",
	70: "Six of Swords ",
	7000: "Six of Swords, Reversed ",
	71: "Seven of Swords ",
	7100: "Seven of Swords, Reversed ",
	72: "Eight of Swords ",
	7200: "Eight of Swords, Reversed ",
	73: "Nine of Swords ",
	7300: "Nine of Swords, Reversed ",
	74: "Ten of Swords ",
	7400: "Ten of Swords, Reversed ",
	75: "Page of Swords ",
	7500: "Page of Swords, Reversed ",
	76: "Knight of Swords ",
	7600: "Knight of Swords, Reversed ",
	77: "Queen of Swords ",
	7700: "Queen of Swords, Reversed ",
	78: "King of Swords ",
	7800: "King of Swords, Reversed ",

	}


	var i int = 78
	var list []int
	rand.Seed(time.Now().UTC().UnixNano())
	list = rand.Perm(i)


	for i := 0; i < 78; i++ {
		list[i] = list[i] + 1
	if arg=="reverse" {

		if 	(rand.Intn(3 - 1)) == 1 {
			list[i] = list[i]  * 100
		}
}

	}


//	for i:=0 ; i < 78 ;  i++ {
//		fmt.Println(riderDeck[i+1])
//	}
	for i:=0 ; i < 78 ; i++  {
		var j int
	        j = list[i]

		fmt.Println(j)
		fmt.Println(riderDeck[j])
	}

}
