package main

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := 0
	switch {
	case card == "ace":
		cards = 11
	case card == "two":
		cards = 2
	case card == "three":
		cards = 3
	case card == "four":
		cards = 4
	case card == "five":
		cards = 5
	case card == "six":
		cards = 6
	case card == "seven":
		cards = 7
	case card == "eight":
		cards = 8
	case card == "nine":
		cards = 9
	case card == "ten":
		cards = 10
	case card == "jack":
		cards = 10
	case card == "queen":
		cards = 10
	case card == "king":
		cards = 10
	case card == "other":
		cards = 0
	}
	return cards
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerAll := ParseCard(card1) + ParseCard(card2)
	dealerAll := ParseCard(dealerCard)
	var option string
	switch {
	case card1 == "ace" && card2 == "ace":
		option = "P"
	case playerAll == 21:
		if dealerCard == "ace" || dealerAll == 10 {
			option = "W"
		} else {
			option = "S"
		}
	case playerAll >= 17 && playerAll <= 20:
		option = "S"
	case playerAll >= 12 && playerAll <= 16:
		if dealerAll >= 7 {
			option = "H"
		} else {
			option = "S"
		}
	case playerAll <= 11:
		option = "H"

	}
	return option
}
