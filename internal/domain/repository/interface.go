package repository

import (
	"context"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"
)

type MyTeamRepositoryInterface interface {
	FindByID(ctx context.Context, id string) (*entity.MyTeam, error)
	FindByIDForUpdate(ctx context.Context, id string) (*entity.MyTeam, error)
	AddScore(ctx context.Context, team *entity.MyTeam, score float64) error
	Create(ctx context.Context, team *entity.MyTeam) error
	FindAllPlayers(ctx context.Context, teamID string) ([]entity.Player, error)
	SavePlayers(ctx context.Context, myTeam *entity.MyTeam) error
}

type MatchRepositoryInterface interface {
	Create(ctx context.Context, match *entity.Match) error
	SaveActions(ctx context.Context, match *entity.Match, score float64) error
	FindByID(ctx context.Context, id string) (*entity.Match, error)
	Update(ctx context.Context, match *entity.Match) error
	FindAll(ctx context.Context) ([]*entity.Match, error)
}

type PlayerRepositoryInterface interface {
	FindByID(ctx context.Context, id string) (*entity.Player, error)
	FindByIDForUpdate(ctx context.Context, id string) (*entity.Player, error)
	Create(ctx context.Context, player *entity.Player) error
	Update(ctx context.Context, player *entity.Player) error
	FindAll(ctx context.Context) ([]*entity.Player, error)
	FindAllByIDs(ctx context.Context, ids []string) ([]entity.Player, error)
}

type TeamRepositoryInterface interface {
	FindByID(ctx context.Context, id string) (*entity.Team, error)
}
