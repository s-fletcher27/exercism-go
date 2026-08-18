package blackjack

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
	case "10", "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	val1 := ParseCard(card1)
	val2 := ParseCard(card2)
	dealerVal := ParseCard(dealerCard)
	sum := val1 + val2

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case sum == 21:
		if dealerVal < 10 {
			return "W"
		}
		return "S"
	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16:
		if dealerVal >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}