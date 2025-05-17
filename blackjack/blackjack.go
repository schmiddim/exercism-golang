package blackjack

var cards = []struct {
	card  string
	value int
}{
	{card: "ace", value: 11},
	{card: "two", value: 2},
	{card: "three", value: 3},
	{card: "four", value: 4},
	{card: "five", value: 5},
	{card: "six", value: 6},
	{card: "seven", value: 7},
	{card: "eight", value: 8},
	{card: "nine", value: 9},
	{card: "ten", value: 10},
	{card: "jack", value: 10},
	{card: "queen", value: 10},
	{card: "king", value: 10},
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

	for _, c := range cards {
		if card == c.card {
			return c.value
		}
	}
	return 0

}

func calculateHand(card1, card2 string) int {
	var v1, v2 int
	for _, c := range cards {
		if card1 == c.card {
			v1 = c.value
		}
		if card2 == c.card {
			v2 = c.value
		}
	}
	return v1 + v2

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	var response string
	hand := calculateHand(card1, card2)
	dealerValue := ParseCard(dealerCard)
	if card1 == card2 && card1 == "ace" {
		return "P"
	}
	switch {
	case hand == 22:
		response = "P"
	case hand <= 11:
		response = "H"
	case hand >= 17 && hand <= 20:
		response = "S"
	case hand >= 12 && hand <= 16 && dealerValue >= 7:
		response = "H"
	case hand >= 12 && hand <= 16 && dealerValue < 7:
		response = "S"
	case hand == 21 && dealerValue < 10:
		response = "W"
	case hand == 21 && dealerValue >= 10:
		response = "S"
	}

	return response
}
