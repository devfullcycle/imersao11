package entity

import "github.com/google/uuid"

type GameAction struct {
	ID         string
	PlayerID   string
	PlayerName string
	TeamID     string
	Minute     int
	Action     string
	Score      int
}

func NewGameAction(playerID string, minute int, action string, score int, teamID string) *GameAction {
	return &GameAction{
		ID:       uuid.New().String(),
		PlayerID: playerID,
		Minute:   minute,
		TeamID:  teamID,
		Action:   action,
		Score:    score,
	}
}
