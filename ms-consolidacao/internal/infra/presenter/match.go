package presenter

import "github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"

type Matches []Match

type Match struct {
	ID      string        `json:"id"`
	TeamA   string        `json:"team_a"`
	TeamB   string        `json:"team_b"`
	TeamAID string        `json:"team_a_id"`
	TeamBID string        `json:"team_b_id"`
	Date    string        `json:"match_date"`
	Status  string        `json:"status"`
	Result  string        `json:"result"`
	Actions []MatchAction `json:"actions"`
}

type MatchAction struct {
	PlayerID   string `json:"player_id"`
	PlayerName string `json:"player_name"`
	Minute     int    `json:"minutes"`
	Action     string `json:"action"`
	Score      int    `json:"score"`
}

func NewMatchPresenter(match *entity.Match) Match {
	matchPresenter := Match{
		ID:      match.ID,
		TeamA:   match.TeamA.Name,
		TeamB:   match.TeamB.Name,
		TeamAID: match.TeamA.ID,
		TeamBID: match.TeamB.ID,
		Date:    match.Date.Format("2006-01-02"),
		Status:  match.Status,
		Result:  match.Result.GetResult(),
	}

	var actions []MatchAction
	for _, action := range match.Actions {
		actions = append(actions, MatchAction{
			PlayerID:   action.PlayerID,
			PlayerName: action.PlayerName,
			Minute:     action.Minute,
			Action:     action.Action,
			Score:      action.Score,
		})
	}
	matchPresenter.Actions = actions
	return matchPresenter
}
