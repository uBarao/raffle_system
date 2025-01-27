package models

import "sort"

type Result struct {
	SortedNumbers []int8         `json:"sortedNumbers"`
	TotalRounds   int            `json:"totalRounds"`
	TotalWinners  int            `json:"totalWinners"`
	WinningBets   []Bet          `json:"winningBets"`
	BetsPerNumber []BettedNumber `json:"betsPerNumber"`
}

type BettedNumber struct {
	Number      int `json:"number"`
	TimesBetted int `json:"timesBetted"`
}

func CalculateResult(sortedNumbers []int8, winningBets []Bet, allBets []Bet) Result {
	result := Result{}

	result.SortedNumbers = sortedNumbers
	result.TotalRounds = len(sortedNumbers) - 4
	result.TotalWinners = len(winningBets)
	result.WinningBets = winningBets
	result.BetsPerNumber = calculateBetsPerNumber(allBets)

	return result
}

func calculateBetsPerNumber(bets []Bet) []BettedNumber {
	betCount := make(map[int8]int)

	for i := 1; i <= 50; i++ {
		betCount[int8(i)] = 0
	}

	for _, bet := range bets {
		for _, betNumber := range bet.Numbers {
			betCount[betNumber]++
		}
	}

	// adicionar na ordem que mais aparecem até o que menos aparece
	keys := make([]int8, 0, len(betCount))

	for key := range betCount {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return betCount[keys[i]] > betCount[keys[j]]
	})

	// monta o array de numeros e a quantidade de apostas
	betsPerNumber := []BettedNumber{}

	for _, number := range keys {
		betsPerNumber = append(betsPerNumber, BettedNumber{Number: int(number), TimesBetted: betCount[number]})
	}

	return betsPerNumber
}
