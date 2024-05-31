package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "ten", "jack", "queen", "king":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == card1 {
		return "P"
	}
	var sum int = ParseCard(card1) + ParseCard(card2)
	var dealer int = ParseCard(dealerCard)
	if sum == 21 {
		if dealerCard == "ace" || dealer == 10 {
			return "S"
		}
		return "W"
	} else if sum >= 17 {
		return "S"
	} else if sum >= 12 {
		if dealer <= 6 {
			return "S"
		}
		return "H"
	} else if sum <= 11 {
		return "H"
	}
	return "S"
}
