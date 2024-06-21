package main

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := 0
	switch card {
	case "ace":
		cards = 11
	case "two":
		cards = 2
	case "three":
		cards = 3
	case "four":
		cards = 4
	case "five":
		cards = 5
	case "six":
		cards = 6
	case "seven":
		cards = 7
	case "eight":
		cards = 8
	case "nine":
		cards = 9
	case "ten", "jack", "queen", "king":
		cards = 10

	case "other":
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
		if dealerAll >= 10 {
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
