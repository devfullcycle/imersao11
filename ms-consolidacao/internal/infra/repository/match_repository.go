package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"
	"github.com/devfullcycle/imersao10-consolidacao/internal/infra/db"
)

type MatchRepository struct {
	dbConn *sql.DB
	*db.Queries
}

func NewMatchRepository(dbConn *sql.DB) *MatchRepository {
	return &MatchRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *MatchRepository) Create(ctx context.Context, match *entity.Match) error {
	return r.Queries.CreateMatch(ctx, db.CreateMatchParams{
		ID:        match.ID,
		MatchDate: sql.NullTime{Time: time.Now(), Valid: true},
		TeamAID:   sql.NullString{String: match.TeamA.ID, Valid: true},
		TeamBID:   sql.NullString{String: match.TeamB.ID, Valid: true},
		TeamAName: sql.NullString{String: match.TeamA.Name, Valid: true},
		TeamBName: sql.NullString{String: match.TeamB.Name, Valid: true},
	})
}

func (r *MatchRepository) SaveActions(ctx context.Context, match *entity.Match, score float64) error {
	err := r.deleteMatchActions(ctx, match.ID)
	if err != nil {
		return err
	}
	for _, action := range match.Actions {
		err := r.Queries.CreateAction(ctx, db.CreateActionParams{
			ID:       action.ID,
			MatchID:  match.ID,
			TeamID:   action.TeamID,
			PlayerID: action.PlayerID,
			Minute:   int32(action.Minute),
			Action:   action.Action,
			Score:    score,
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (r *MatchRepository) deleteMatchActions(ctx context.Context, matchID string) error {
	err := r.Queries.RemoveActionFromMatch(ctx, matchID)
	return err
}

func (r *MatchRepository) FindByID(ctx context.Context, id string) (*entity.Match, error) {
	match, err := r.Queries.FindMatchById(ctx, id)
	if err != nil {
		return nil, err
	}

	teamAResult := 0
	teamBResult := 0
	if match.Result.String != "" {
		matchResult := strings.Split(match.Result.String, "-")
		teamAResult, err = strconv.Atoi(matchResult[0])
		if err != nil {
			return nil, err
		}
		teamBResult, err = strconv.Atoi(matchResult[1])
		if err != nil {
			return nil, err
		}
	}
	// get actions
	actions, err := r.Queries.GetMatchActions(ctx, id)
	if err != nil {
		return nil, err
	}

	// convert actions to entity
	var gameActions []entity.GameAction
	for _, action := range actions {
		playerInfo, err := r.Queries.FindPlayerById(ctx, action.PlayerID)
		if err != nil {
			return nil, err
		}
		gameActions = append(gameActions, entity.GameAction{
			ID:         action.ID,
			PlayerID:   action.PlayerID,
			PlayerName: playerInfo.Name,
			Minute:     int(action.Minute),
			Action:     action.Action,
		})
	}

	return &entity.Match{
		ID:      match.ID,
		TeamA:   &entity.Team{ID: match.TeamAID.String, Name: match.TeamAName.String},
		TeamB:   &entity.Team{ID: match.TeamBID.String, Name: match.TeamBName.String},
		Result:  *entity.NewMatchResult(teamAResult, teamBResult),
		Date:    match.MatchDate.Time,
		Actions: gameActions,
	}, nil
}

func (r *MatchRepository) Update(ctx context.Context, match *entity.Match) error {
	matchResult := match.Result.GetResult()
	return r.Queries.UpdateMatch(ctx, db.UpdateMatchParams{
		ID:        match.ID,
		Result:    sql.NullString{String: matchResult, Valid: true},
		MatchDate: sql.NullTime{Time: match.Date, Valid: true},
		TeamAID:   sql.NullString{String: match.TeamA.ID, Valid: true},
		TeamAName: sql.NullString{String: match.TeamA.Name, Valid: true},
		TeamBID:   sql.NullString{String: match.TeamB.ID, Valid: true},
		TeamBName: sql.NullString{String: match.TeamB.Name, Valid: true},
	})
}

func (r *MatchRepository) FindAll(ctx context.Context) ([]*entity.Match, error) {
	matches, err := r.Queries.FindAllMatches(ctx)
	if err != nil {
		return nil, err
	}
	var gameActions []entity.GameAction
	var matchList []*entity.Match

	for _, match := range matches {
		teamAResult := 0
		teamBResult := 0
		if match.Result.String != "" {
			matchResult := strings.Split(match.Result.String, "-")
			teamAResult, err = strconv.Atoi(matchResult[0])
			if err != nil {
				return nil, err
			}
			teamBResult, err = strconv.Atoi(matchResult[1])
			if err != nil {
				return nil, err
			}
		}
		// get actions
		actions, err := r.Queries.GetMatchActions(ctx, match.ID)
		if err != nil {
			return nil, err
		}

		// convert actions to entity
		for _, action := range actions {
			playerInfo, err := r.Queries.FindPlayerById(ctx, action.PlayerID)
			if err != nil {
				return nil, err
			}
			gameActions = append(gameActions, entity.GameAction{
				ID:         action.ID,
				PlayerID:   action.PlayerID,
				PlayerName: playerInfo.Name,
				Minute:     int(action.Minute),
				Action:     action.Action,
				Score:      int(action.Score),
			})
		}

		matchList = append(matchList, &entity.Match{
			ID:      match.ID,
			TeamA:   &entity.Team{ID: match.TeamAID.String, Name: match.TeamAName.String},
			TeamB:   &entity.Team{ID: match.TeamBID.String, Name: match.TeamBName.String},
			TeamAID: match.TeamAID.String,
			TeamBID: match.TeamBID.String,
			Result:  *entity.NewMatchResult(teamAResult, teamBResult),
			Date:    match.MatchDate.Time,
			Actions: gameActions,
		})
	}
	return matchList, nil
}
