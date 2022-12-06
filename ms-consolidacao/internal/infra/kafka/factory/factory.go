package factory

import "github.com/devfullcycle/imersao10-consolidacao/internal/infra/kafka/event"

func CreateProcessMessageStrategy(topic string) event.ProcessEventStrategy {
	switch topic {
	case "chooseTeam":
		// {"my_team_id":"1", "players":["1","2","3","4","5"]}
		return event.ProcessChooseTeam{}
	case "newPlayer":
		// {"id": "10","name": "Wesley","initial_price": 10.5}
		return event.ProcessNewPlayer{}
	case "newMatch":
		// {"id":"3","match_date":"2021-05-01T00:00:00Z","team_a_id":"1","team_b_id":"2"}
		return event.ProcessNewMatch{}
	case "newAction":
		// {"match_id":"3","team_id":"1","player_id":"1","action":"goal","minutes":10}
		return event.ProcessNewAction{}
	case "matchUpdateResult":
		// {"match_id":"1","result":"2-0"}
		return event.ProcessMatchUpdateResult{}
	}
	return nil
}
