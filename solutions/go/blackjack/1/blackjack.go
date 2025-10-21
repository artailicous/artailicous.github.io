package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Calculate player's hand value
	playerValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	// Strategy rules:
	// 1. If you have a pair of aces you must always split them
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	// 2. If you have a Blackjack (21), and dealer doesn't have ace/face/ten then auto win
	if playerValue == 21 {
		if dealerValue == 11 || dealerValue == 10 { // ace, jack, queen, king, ten
			return "S" // stand and wait
		}
		return "W" // automatically win
	}

	// 3. If cards sum to [17, 20] always stand
	if playerValue >= 17 && playerValue <= 20 {
		return "S"
	}

	// 4. If cards sum to [12, 16] stand unless dealer has 7 or higher
	if playerValue >= 12 && playerValue <= 16 {
		if dealerValue >= 7 {
			return "H" // hit
		}
		return "S" // stand
	}

	// 5. If cards sum to 11 or lower always hit
	if playerValue <= 11 {
		return "H"
	}

	// This should never be reached based on the rules, but return hit as default
	return "H"
}
