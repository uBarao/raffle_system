package models

import (
	"errors"
	"github/raffle_system/utils"
)

type Raffle struct {
	Bets         []Bet
	nextRegister int
}

var ActiveRaffle Raffle

func Start() *Raffle {
	ActiveRaffle = Raffle{nextRegister: 1000}
	return &ActiveRaffle
}

func (r *Raffle) AddBet(name string, cpf string, numbers []int8) (Bet, error) {
	if len(numbers) != 5 {
		return Bet{}, errors.New("not five numbers bet")
	}
	if isRepeating(numbers) {
		return Bet{}, errors.New("repeating bet numbers")
	}
	if isOutOfRange(numbers) {
		return Bet{}, errors.New("out of range bet numbers")
	}
	bet := Bet{
		Id:      r.nextRegister,
		Name:    name,
		CPF:     cpf,
		Numbers: numbers,
	}
	r.nextRegister++
	r.Bets = append(r.Bets, bet)
	return bet, nil
}

func (r *Raffle) FindBetByName(name string) []Bet {
	bets := []Bet{}

	for _, bet := range r.Bets {
		if bet.Name == name {
			bets = append(bets, bet)
		}
	}

	return bets
}

func (r *Raffle) FindBetByCPF(cpf string) []Bet {
	bets := []Bet{}

	for _, bet := range r.Bets {
		if bet.CPF == cpf {
			bets = append(bets, bet)
		}
	}

	return bets
}

// Numeros sorteados, apostas vencedoras
func (r *Raffle) Run() ([]int8, []Bet) {
	// Initial numbers
	drawnNumbers := RandomBetNumbers(5)
	winners := r.calculateWinners(drawnNumbers)

	for i := 0; i < 25 && len(winners) <= 0; i++ {
		newNumber := utils.UnrepeatableRandomInt8InRange(drawnNumbers, 50)
		drawnNumbers = append(drawnNumbers, newNumber)

		winners = r.calculateWinners(drawnNumbers)
	}

	return drawnNumbers, winners
}

func (r *Raffle) calculateWinners(drawnNumbers []int8) []Bet {
	winners := []Bet{}

	for _, bet := range r.Bets {
		if isWinner(bet, drawnNumbers) {
			winners = append(winners, bet)
		}
	}
	return winners
}

func isRepeating(numbers []int8) bool {
	for i := range numbers {
		for j := range numbers {
			if i != j && numbers[i] == numbers[j] {
				return true
			}
		}
	}
	return false
}

func isOutOfRange(numbers []int8) bool {
	for _, number := range numbers {
		if number <= 0 || number > 50 {
			return true
		}
	}
	return false
}

func isWinner(bet Bet, numbers []int8) bool {
	for _, betNumber := range bet.Numbers {
		if !isPresent(betNumber, numbers) {
			return false
		}
	}
	return true
}

func isPresent(betNumber int8, numbers []int8) bool {
	for _, drawnNumber := range numbers {
		if drawnNumber == betNumber {
			return true
		}
	}
	return false
}
