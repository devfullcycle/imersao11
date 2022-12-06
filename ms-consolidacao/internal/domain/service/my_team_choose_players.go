package service

import (
	"errors"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"
)

var errNotEnoughMoney = errors.New("not enough money")

func ChoosePlayers(myTeam *entity.MyTeam, myPlayers []entity.Player, players []entity.Player) error {
	totalCost := 0.0
	totalEarned := calculateTotalEarned(myPlayers, players)

	for _, player := range players {
		if !playerInMyTeam(player, *myTeam) && playerInPlayersList(player, players) {
			totalCost += player.Price
		}
	}

	if totalCost > myTeam.Score+totalEarned {
		return errNotEnoughMoney
	}

	myTeam.Score += totalEarned - totalCost
	myTeam.Players = []string{}

	for _, player := range players {
		myTeam.Players = append(myTeam.Players, player.ID)
	}
	return nil
}

func playerInMyTeam(player entity.Player, myTeam entity.MyTeam) bool {
	for _, p := range myTeam.Players {
		if p == player.ID {
			return true
		}
	}
	return false
}

func playerInPlayersList(player entity.Player, players []entity.Player) bool {
	for _, p := range players {
		if p.ID == player.ID {
			return true
		}
	}
	return false
}

// get the difference between two slices
func calculateTotalEarned(myPlayers []entity.Player, players []entity.Player) float64 {
	var totalEarned float64
	for _, myPlayer := range myPlayers {
		if !playerInPlayersList(myPlayer, players) {
			totalEarned += myPlayer.Price
		}
	}
	return totalEarned
}
