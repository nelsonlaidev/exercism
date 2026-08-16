package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	n1 := ParseCard(card1)
	n2 := ParseCard(card2)
	d := ParseCard(dealerCard)
	mySum := n1 + n2

	switch {
	case n1 == 11 && n2 == 11:
		return "P"
	case mySum == 21:
		if d != 11 && d != 10 {
			return "W"
		} else {
			return "S"
		}
	case mySum <= 20 && mySum >= 17:
		return "S"
	case mySum <= 16 && mySum >= 12:
		if d >= 7 {
			return "H"
		}
		return "S"
	case mySum <= 11:
		return "H"
	default:
		return "H"
	}
}
