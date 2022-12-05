package usecase

import (
	"context"
	"time"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"
	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/repository"
	"github.com/devfullcycle/imersao10-consolidacao/pkg/uow"
)

type MatchInput struct {
	ID      string
	Date    time.Time
	TeamAID string
	TeamBID string
}

type MatchUseCase struct {
	Uow uow.UowInterface
}

func (a *MatchUseCase) Execute(ctx context.Context, input MatchInput) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		matchRepository := a.getMatchRepository(ctx)
		teamRepository := a.getTeamRepository(ctx)

		teamA, err := teamRepository.FindByID(ctx, input.TeamAID)
		if err != nil {
			return err
		}
		teamB, err := teamRepository.FindByID(ctx, input.TeamBID)
		if err != nil {
			return err
		}

		match := entity.NewMatch(input.ID, teamA, teamB, input.Date)
		err = matchRepository.Create(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
}

func (a *MatchUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := a.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

func (a *MatchUseCase) getTeamRepository(ctx context.Context) repository.TeamRepositoryInterface {
	teamRepository, err := a.Uow.GetRepository(ctx, "TeamRepository")
	if err != nil {
		panic(err)
	}
	return teamRepository.(repository.TeamRepositoryInterface)
}
