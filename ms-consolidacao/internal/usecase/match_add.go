package usecase

import (
	"context"
	"time"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"
	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/repository"
	"github.com/devfullcycle/imersao10-consolidacao/pkg/uow"
)

type MatchInput struct {
	ID      string    `json:"id"`
	Date    time.Time `json:"match_date"`
	TeamAID string    `json:"team_a_id"`
	TeamBID string    `json:"team_b_id"`
}

type MatchUseCase struct {
	Uow uow.UowInterface
}

func NewMatchUseCase(uow uow.UowInterface) *MatchUseCase {
	return &MatchUseCase{
		Uow: uow,
	}
}

func (u *MatchUseCase) Execute(ctx context.Context, input MatchInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		matchRepository := u.getMatchRepository(ctx)
		teamRepository := u.getTeamRepository(ctx)

		teamA, err := teamRepository.FindByID(ctx, input.TeamAID)
		if err != nil {
			return err
		}
		teamB, err := teamRepository.FindByID(ctx, input.TeamBID)
		if err != nil {
			return err
		}

		match := entity.NewMatch(input.ID, teamA, teamB, input.Date)

		// Create match
		err = matchRepository.Create(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *MatchUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := u.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

func (u *MatchUseCase) getTeamRepository(ctx context.Context) repository.TeamRepositoryInterface {
	teamRepository, err := u.Uow.GetRepository(ctx, "TeamRepository")
	if err != nil {
		panic(err)
	}
	return teamRepository.(repository.TeamRepositoryInterface)
}
