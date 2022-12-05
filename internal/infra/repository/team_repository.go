package repository

import (
	"context"
	"database/sql"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"
	"github.com/devfullcycle/imersao10-consolidacao/internal/infra/db"
)

type TeamRepository struct {
	dbConn *sql.DB
	*db.Queries
	Repository
}

func NewTeamRepository(dbConn *sql.DB) *TeamRepository {
	return &TeamRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *TeamRepository) AddScore(ctx context.Context, player *entity.Player, score float64) error {
	err := r.Queries.AddScoreToTeam(ctx, db.AddScoreToTeamParams{
		ID:    player.ID,
		Score: score,
	})
	return err
}

func (r *TeamRepository) FindByID(ctx context.Context, id string) (*entity.Team, error) {
	team, err := r.Queries.FindTeamById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Team{
		ID:   team.ID,
		Name: team.Name,
	}, nil
}
