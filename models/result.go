package models

type Result struct {
	SortedNumbers []int8 `json:"sortedNumbers"`
	TotalRounds   int    `json:"totalRounds"`
	WinningBets   []Bet  `json:"winningBets"`
	TotalBets     []int8 `json:"totalBets"`
}

func CalculateResult(sortedNumbers []int8, winningBets []Bet) Result {
	result := Result{}
	result.SortedNumbers = sortedNumbers
	result.TotalRounds = len(sortedNumbers) - 4
	result.WinningBets = winningBets
	result.TotalBets = []int8{}
	return result
}
