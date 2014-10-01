package main

import (
	"flag"
	"fmt"
	"time"
	"math/rand"
)

func getDeckSize(s string) int  {
	switch s {
		case "Rider", "shadow" :
			return 78
		case "Wildwood" :
			return 78
		}
	return 0
}

func main() {

	deckPtr := flag.String("deck", "Rider", "a string")
	spreadPtr := flag.String("spread", "CelticCross", "a string")
	reversePtr := flag.Bool("reverse", true, "a bool")
	//shufflePtr := flag.Bool("shuffle", true, "a bool")

	flag.Parse()

	fmt.Println()
	fmt.Println("Deck: ", *deckPtr)
	fmt.Println("Spread: ", *spreadPtr)
	fmt.Println("Reverse: ", *reversePtr)
//	fmt.Println("Shuffle: ", *shufflePtr)

	tarotDeck := make(map[int]string)
	if *deckPtr == "shadow" || *deckPtr == "Rider" {
		tarotDeck = map[int]string{

			1:    "The Fool ",
			100:  "The Fool, Reversed ",
			2:    "The Magician ",
			200:  "The Magician, Reversed ",
			3:    "The High Priestess ",
			300:  "The High Priestess, Reversed ",
			4:    "The Empress ",
			400:  "The Empress, Reversed ",
			5:    "The Emperor ",
			500:  "The Emperor, Reversed ",
			6:    "The Hierophant ",
			600:  "The Hierophant, Reversed ",
			7:    "The Lovers ",
			700:  "The Lovers, Reversed ",
			8:    "The Chariot ",
			800:  "The Chariot, Reversed ",
			9:    "Strength ",
			900:  "Strength, Reversed ",
			10:   "The Hermit ",
			1000: "The Hermit, Reversed ",
			11:   "Wheel of Fortune ",
			1100: "Wheel of Fortune, Reversed ",
			12:   "Justice ",
			1200: "Justice, Reversed ",
			13:   "The Hanged Man ",
			1300: "The Hanged Man, Reversed ",
			14:   "Death ",
			1400: "Death, Reversed ",
			15:   "Temperance ",
			1500: "Temperance, Reversed ",
			16:   "The Devil ",
			1600: "The Devil, Reversed ",
			17:   "The Tower ",
			1700: "The Tower, Reversed ",
			18:   "The Star ",
			1800: "The Star, Reversed ",
			19:   "The Moon ",
			1900: "The Moon, Reversed ",
			20:   "The Sun ",
			2000: "The Sun, Reversed ",
			21:   "Judgment ",
			2100: "Judgement, Reversed ",
			22:   "The World ",
			2200: "The World, Reversed ",

			23:   "Ace of Wands ",
			2300: "Ace of Wands, Reversed ",
			24:   "Two of Wands ",
			2400: "Two of Wands, Reversed ",
			25:   "Three of Wands ",
			2500: "Three of Wands, Reversed ",
			26:   "Four of Wands ",
			2600: "Four of Wands, Reversed ",
			27:   "Five of Wands ",
			2700: "Five of Wands, Reversed ",
			28:   "Six of Wands ",
			2800: "Six of Wands, Reversed ",
			29:   "Seven of Wands ",
			2900: "Seven of Wands, Reversed ",
			30:   "Eight of Wands ",
			3000: "Eight of Wands, Reversed ",
			31:   "Nine of Wands ",
			3100: "Nine of Wands, Reversed ",
			32:   "Ten of Wands ",
			3200: "Ten of Wands, Reversed ",
			33:   "Page of Wands ",
			3300: "Page of Wands, Reversed ",
			34:   "Knight of Wands ",
			3400: "Knight of Wands, Reversed ",
			35:   "Queen of Wands ",
			3500: "Queen of Wands, Reversed ",
			36:   "King of Wands ",
			3600: "King of Wands, Reversed ",

			37:   "Ace of Pentacles ",
			3700: "Ace of Pentacles, Reversed ",
			38:   "Two of Pentacles ",
			3800: "Two of Pentacles, Reversed ",
			39:   "Three of Pentacles ",
			3900: "Three of Pentacles, Reversed ",
			40:   "Four of Pentacles ",
			4000: "Four of Pentacles, Reversed ",
			41:   "Five of Pentacles ",
			4100: "Five of Pentacles, Reversed ",
			42:   "Six of Pentacles ",
			4200: "Six of Pentacles, Reversed ",
			43:   "Seven of Pentacles ",
			4300: "Seven of Pentacles, Reversed ",
			44:   "Eight of Pentacles ",
			4400: "Eight of Pentacles, Reversed ",
			45:   "Nine of Pentacles ",
			4500: "Nine of Pentacles, Reversed ",
			46:   "Ten of Pentacles ",
			4600: "Ten of Pentacles, Reversed ",
			47:   "Page of Pentacles ",
			4700: "Page of Pentacles, Reversed ",
			48:   "Knight of Pentacles ",
			4800: "Knight of Pentacles, Reversed ",
			49:   "Queen of Pentacles ",
			4900: "Queen of Pentacles, Reversed ",
			50:   "King of Pentacles ",
			5000: "King of Pentacles, Reversed ",

			51:   "Ace of Cups ",
			5100: "Ace of Cups, Reversed ",
			52:   "Two of Cups ",
			5200: "Two of Cups, Reversed ",
			53:   "Three of Cups ",
			5300: "Three of Cups, Reversed ",
			54:   "Four of Cups ",
			5400: "Four of Cups, Reversed ",
			55:   "Five of Cups ",
			5500: "Five of Cups, Reversed ",
			56:   "Six of Cups ",
			5600: "Six of Cups, Reversed ",
			57:   "Seven of Cups ",
			5700: "Seven of Cups, Reversed ",
			58:   "Eight of Cups ",
			5800: "Eight of Cups, Reversed ",
			59:   "Nine of Cups ",
			5900: "Nine of Cups, Reversed ",
			60:   "Ten of Cups ",
			6000: "Ten of Cups, Reversed ",
			61:   "Page of Cups ",
			6100: "Page Cups, Reversed ",
			62:   "Knight of Cups ",
			6200: "Knight of Cups, Reversed ",
			63:   "Queen of Cups ",
			6300: "Queen of Cups, Reversed ",
			64:   "King of Cups ",
			6400: "King of Cups, Reversed ",

			65:   "Ace of Swords ",
			6500: "Ace of Swords, Reversed ",
			66:   "Two of Swords ",
			6600: "Two of Swords, Reversed ",
			67:   "Three of Swords ",
			6700: "Three of Swords, Reversed ",
			68:   "Four of Swords ",
			6800: "Four of Swords, Reversed ",
			69:   "Five of Swords ",
			6900: "Five of Swords, Reversed ",
			70:   "Six of Swords ",
			7000: "Six of Swords, Reversed ",
			71:   "Seven of Swords ",
			7100: "Seven of Swords, Reversed ",
			72:   "Eight of Swords ",
			7200: "Eight of Swords, Reversed ",
			73:   "Nine of Swords ",
			7300: "Nine of Swords, Reversed ",
			74:   "Ten of Swords ",
			7400: "Ten of Swords, Reversed ",
			75:   "Page of Swords ",
			7500: "Page of Swords, Reversed ",
			76:   "Knight of Swords ",
			7600: "Knight of Swords, Reversed ",
			77:   "Queen of Swords ",
			7700: "Queen of Swords, Reversed ",
			78:   "King of Swords ",
			7800: "King of Swords, Reversed ",
}
		} else {



	if *deckPtr == "Wildwood" {

		tarotDeck = map[int]string{
			1:    "The Wanderer ",
			100:  "The Wanderer, Reversed ",
			2:    "The Shaman ",
			200:  "The Shaman, Reversed ",
			3:    "The Seer ",
			300:  "The Seer, Reversed ",
			4:    "The Green Woman ",
			400:  "The Green Woman, Reversed ",
			5:    "The Green Man ",
			500:  "The Green Man, Reversed ",
			6:    "The Ancestor ",
			600:  "The Ancestor, Reversed ",
			7:    "The Forest Lovers ",
			700:  "The Forest Lovers, Reversed ",
			8:    "The Archer ",
			800:  "The Archer, Reversed ",
			9:    "The Stag ",
			900:  "The Stag, Reversed ",
			10:   "The Hooded Man ",
			1000: "The Hooded Man, Reversed ",
			11:   "The Wheel ",
			1100: "The Wheel, Reversed ",
			12:   "The Woodward ",
			1200: "The Woodward, Reversed ",
			13:   "The Mirror ",
			1300: "The Mirror, Reversed ",
			14:   "The Journey ",
			1400: "The Journey, Reversed ",
			15:   "Balance ",
			1500: "Balance, Reversed ",
			16:   "The Guardian ",
			1600: "The Guardian, Reversed ",
			17:   "The Blasted Oak ",
			1700: "The Blasted Oak, Reversed ",
			18:   "The Pole Star ",
			1800: "The Pole Star, Reversed ",
			19:   "The Moon on Water",
			1900: "The Moon on Water, Reversed ",
			20:   "The Sun of Life",
			2000: "The Sun of Life, Reversed ",
			21:   "The Great Bear ",
			2100: "The Great Bear, Reversed ",
			22:   "The World Tree",
			2200: "The World Tree, Reversed ",

			23:   "Ace of Arrows ",
			2300: "Ace of Arrows, Reversed ",
			24:   "Two of Arrows ",
			2400: "Two of Arrows, Reversed ",
			25:   "Three of Arrows ",
			2500: "Three of Arrows, Reversed ",
			26:   "Four of Arrows ",
			2600: "Four of Arrows, Reversed ",
			27:   "Five of Arrows ",
			2700: "Five of Arrows, Reversed ",
			28:   "Six of Arrows ",
			2800: "Six of Arrows, Reversed ",
			29:   "Seven of Arrows ",
			2900: "Seven of Arrows, Reversed ",
			30:   "Eight of Arrows ",
			3000: "Eight of Arrows, Reversed ",
			31:   "Nine of Arrows ",
			3100: "Nine of Arrows, Reversed ",
			32:   "Ten of Arrows ",
			3200: "Ten of Arrows, Reversed ",
			33:   "Page of Arrows ",
			3300: "Page of Arrows, Reversed ",
			34:   "Knight of Arrows ",
			3400: "Knight of Arrows, Reversed ",
			35:   "Queen of Arrows ",
			3500: "Queen of Arrows, Reversed ",
			36:   "King of Arrows ",
			3600: "King of Arrows, Reversed ",

			37:   "Ace of Bows ",
			3700: "Ace of Bows, Reversed ",
			38:   "Two of Bows ",
			3800: "Two of Bows, Reversed ",
			39:   "Three of Bows ",
			3900: "Three of Bows, Reversed ",
			40:   "Four of Bows ",
			4000: "Four of Bows, Reversed ",
			41:   "Five of Bows ",
			4100: "Five of Bows, Reversed ",
			42:   "Six of Bows ",
			4200: "Six of Bows, Reversed ",
			43:   "Seven of Bows ",
			4300: "Seven of Bows, Reversed ",
			44:   "Eight of Bows ",
			4400: "Eight of Bows, Reversed ",
			45:   "Nine of Bows ",
			4500: "Nine of Bows, Reversed ",
			46:   "Ten of Bows ",
			4600: "Ten of Bows, Reversed ",
			47:   "Page of Bows ",
			4700: "Page of Bows, Reversed ",
			48:   "Knight of Bows ",
			4800: "Knight of Bows, Reversed ",
			49:   "Queen of Bows ",
			4900: "Queen of Bows, Reversed ",
			50:   "King of Bows ",
			5000: "King of Bows, Reversed ",

			51:   "Ace of Stones ",
			5100: "Ace of Stones, Reversed ",
			52:   "Two of Stones ",
			5200: "Two of Stones, Reversed ",
			53:   "Three of Stones ",
			5300: "Three of Stones, Reversed ",
			54:   "Four of Stones ",
			5400: "Four of Stones, Reversed ",
			55:   "Five of Stones ",
			5500: "Five of Stones, Reversed ",
			56:   "Six of Stones ",
			5600: "Six of Stones, Reversed ",
			57:   "Seven of Stones ",
			5700: "Seven of Stones, Reversed ",
			58:   "Eight of Stones ",
			5800: "Eight of Stones, Reversed ",
			59:   "Nine of Stones ",
			5900: "Nine of Stones, Reversed ",
			60:   "Ten of Stones ",
			6000: "Ten of Stones, Reversed ",
			61:   "Page of Stones ",
			6100: "Page Stones, Reversed ",
			62:   "Knight of Stones ",
			6200: "Knight of Stones, Reversed ",
			63:   "Queen of Stones ",
			6300: "Queen of Stones, Reversed ",
			64:   "King of Stones ",
			6400: "King of Stones, Reversed ",

			65:   "Ace of Vessels ",
			6500: "Ace of Vessels, Reversed ",
			66:   "Two of Vessels ",
			6600: "Two of Vessels, Reversed ",
			67:   "Three of Vessels ",
			6700: "Three of Vessels, Reversed ",
			68:   "Four of Vessels ",
			6800: "Four of Vessels, Reversed ",
			69:   "Five of Vessels ",
			6900: "Five of Vessels, Reversed ",
			70:   "Six of Vessels ",
			7000: "Six of Vessels, Reversed ",
			71:   "Seven of Vessels ",
			7100: "Seven of Vessels, Reversed ",
			72:   "Eight of Vessels ",
			7200: "Eight of Vessels, Reversed ",
			73:   "Nine of Vessels ",
			7300: "Nine of Vessels, Reversed ",
			74:   "Ten of Vessels ",
			7400: "Ten of Vessels, Reversed ",
			75:   "Page of Vessels ",
			7500: "Page of Vessels, Reversed ",
			76:   "Knight of Vessels ",
			7600: "Knight of Vessels, Reversed ",
			77:   "Queen of Vessels ",
			7700: "Queen of Vessels, Reversed ",
			78:   "King of Vessels ",
			7800: "King of Vessels, Reversed ",
		}

	}

}
	
	var deckSize = getDeckSize(*deckPtr)
//	fmt.Println("getDeckSize", deckSize)
	fmt.Println()



// This following code gets and array of 0-decksize random
//integers, adds 1 to each to change the range to
// 1-decksize, then for each int, randomly multiplies by
// 100 or not. The int is then used to index as a
// key into the appropriate map

	var list []int
	rand.Seed(time.Now().UTC().UnixNano())
	list = rand.Perm(deckSize)

	var i = 0
	for  i = 0; i < deckSize ; i++ {
		list[i] = list[i] + 1
		if *reversePtr {

			if (rand.Intn(3 - 1)) == 1 {
				list[i] = list[i] * 100
			}

			//add some deck cut logic and some keep ptr
			//to dealt point
			//so no new deal until cards gone.  Not
			// needed for tarot, really

		}

	}

	//  Logic for spread count needed and way to make persistant
	// either in big prread next spread loop in mem or write file/index
	// then would have to skip shuffle, etc.or ask new deck

	

	var spreadCount = 10
	
	switch *spreadPtr {
		case "FiveCard" :
			spreadCount = 5
		case "Relationship" :
			spreadCount = 10
		case "CelticCross" :
			spreadCount = 10
		case "Daily" :
			spreadCount = 1
		case "Zodiac" :
			spreadCount = deckSize
		case "ellipse" :
			spreadCount = 7
		case "Mirror" :
			spreadCount = 8
		case "mandala" :
			spreadCount = 8
		}




	spreadMeaning := make(map[int]string)

	if *spreadPtr == "FiveCard"  {
		spreadMeaning = map[int]string{
			1:   "Past influence - how did I get here?",
			2:   "Recent past - coulda, woulda, shoulda...",
			3:   "Present situation - what are today's options?",
			4:   "Near future - likely outcomes",
			5:   "Future - changing outcomes",
		}
	} 

	if *spreadPtr == "Daily" {
		spreadMeaning = map[int]string {
			1:    "The querant's path for today",
		}
	}

	if *spreadPtr == "Mirror" {
		spreadMeaning = map[int]string {
			1:   "The querant and the query",
			2:   "How querant views the other",
			3:   "How the parties view themselves",
			4:   "How the other fulfills the querant",
			5:   "How the other views the querant",
			6:   "Negative flows in the relationship",
			7:   "Positive flows in the relationship",
			8:   "Synergistic outcome",
		}
	}

	if  *spreadPtr == "Relationship" {

		spreadMeaning = map[int]string {
			1:   "The querant and the query",
			2:   "How querant views the other",
			3:   "How the parties view themselves",
			4:   "How the other fulfills the querant",
			5:   "How the other views the querant",
			6:   "Negative flows in the relationship",
			7:   "Positive flows in the relationship",
			8:   "Synergistic outcome",
			9:   "Desired outcome querant",
			10:  "Desired outcome other",

		}
	}
	
	for  i := 0; i < spreadCount; i++ {
		var j int
		j = list[i]

		fmt.Print(i + 1)
		fmt.Print(": ")
		fmt.Println(spreadMeaning[i+1])
		fmt.Println(tarotDeck[j])
		fmt.Println()

	}
}
