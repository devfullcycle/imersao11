package repository

import (
	"context"
	"database/sql"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"
	"github.com/devfullcycle/imersao10-consolidacao/internal/infra/db"
)

type MyTeamRepository struct {
	dbConn *sql.DB
	*db.Queries
}

func NewMyTeamRepository(dbConn *sql.DB) *MyTeamRepository {
	return &MyTeamRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (m *MyTeamRepository) AddScore(ctx context.Context, myTeam *entity.MyTeam, score float64) error {
	// check if my team exists for update
	_, err := m.FindByIDForUpdate(ctx, myTeam.ID)
	if err != nil {
		return err
	}
	myTeam.Score = score + myTeam.Score
	err = m.Queries.UpdateMyTeamScore(ctx, db.UpdateMyTeamScoreParams{
		ID:    myTeam.ID,
		Score: myTeam.Score,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MyTeamRepository) FindByID(ctx context.Context, id string) (*entity.MyTeam, error) {
	team, err := m.Queries.FindMyTeamById(ctx, id)
	if err != nil {
		return nil, err
	}

	players, err := m.GetPlayersByMyTeamID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.MyTeam{
		ID:    team.ID,
		Name:  team.Name,
		Score: team.Score,
		Players: func() []string {
			var output []string
			for _, player := range players {
				output = append(output, player.ID)
			}
			return output
		}(),
	}, nil
}

func (m *MyTeamRepository) FindByIDForUpdate(ctx context.Context, id string) (*entity.MyTeam, error) {
	team, err := m.Queries.FindMyTeamByIdForUpdate(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.MyTeam{
		ID:    team.ID,
		Name:  team.Name,
		Score: team.Score,
	}, nil
}

func (m *MyTeamRepository) Create(ctx context.Context, myTeam *entity.MyTeam) error {
	err := m.Queries.CreateMyTeam(ctx, db.CreateMyTeamParams{
		ID:    myTeam.ID,
		Name:  myTeam.Name,
		Score: myTeam.Score,
	})
	return err
}

func (m *MyTeamRepository) FindAllPlayers(ctx context.Context, teamID string) ([]entity.Player, error) {
	players, err := m.Queries.GetPlayersByMyTeamID(ctx, teamID)
	if err != nil {
		return nil, err
	}

	var output []entity.Player
	for _, player := range players {
		output = append(output, entity.Player{
			ID:    player.ID,
			Name:  player.Name,
			Price: player.Price,
		})
	}

	return output, nil
}

func (m *MyTeamRepository) SavePlayers(ctx context.Context, myTeam *entity.MyTeam) error {
	// delete all players
	err := m.Queries.DeleteAllPlayersFromMyTeam(ctx, myTeam.ID)
	if err != nil {
		return err
	}

	// add all players
	for _, playerId := range myTeam.Players {
		err = m.addPlayer(ctx, myTeam.ID, playerId)
		if err != nil {
			return err
		}
	}

	err = m.Queries.UpdateMyTeamScore(ctx, db.UpdateMyTeamScoreParams{
		ID:    myTeam.ID,
		Score: myTeam.Score,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MyTeamRepository) addPlayer(ctx context.Context, teamID string, playerID string) error {
	err := m.Queries.AddPlayerToMyTeam(ctx, db.AddPlayerToMyTeamParams{
		PlayerID: playerID,
		MyTeamID: teamID,
	})
	return err
}
