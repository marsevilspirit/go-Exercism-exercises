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
        case "joker":
            return 0
        default:
            return 0
        }
}

// firstturn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    playerSum := calculateSum(card1, card2) // 计算玩家手牌点数
    dealerSum := calculateSum(dealerCard)   // 计算庄家手牌点数

    if playerSum == 21 {
        return "blackjack"
    } else if playerSum == 22 && dealerSum == 11 {
        return "p" // Stand
    } else if playerSum == 21 && dealerSum == 11 {
        return "S" // Stand
    } else if playerSum == 20 && dealerSum == 11 {
        return "S" // Stand
    } else if playerSum == 16 && dealerSum >= 7 {
        return "S" // Stand
    } else if playerSum == 15 && dealerSum >= 7 {
        return "S" // Stand
    } else if playerSum == 14 && dealerSum >= 7 {
        return "S" // Stand
    } else if playerSum == 13 && dealerSum >= 7 {
        return "S" // Stand
    } else if playerSum == 12 && dealerSum >= 4 && dealerSum <= 6 {
        return "S" // Stand
    } else if playerSum == 11 {
        return "double"
    } else if card1 == card2 {
        return "P" // Split
    } else {
        return "H" // Hit (default)
    }
}

func calculateSum(cards ...string) int {
    sum := 0
    for _, card := range cards {
        sum += ParseCard(card)
    }
    return sum
}
