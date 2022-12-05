package usecase

import (
	"context"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"
	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/repository"
	"github.com/devfullcycle/imersao10-consolidacao/pkg/uow"
)

type ActionAddInput struct {
	MatchID  string
	TeamID   string
	PlayerID string
	Minute   int
	Action   string
}

type ActionAddUseCase struct {
	Uow         uow.UowInterface
	ActionTable entity.ActionTableInterface
}

func (a *ActionAddUseCase) Execute(ctx context.Context, input ActionAddInput) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		matchRepo := a.getMatchRepository(ctx)
		myTeamRepo := a.getMyTeamRepository(ctx)
		playerRepo := a.getPlayerRepository(ctx)

		match, err := matchRepo.FindByID(ctx, input.MatchID)
		if err != nil {
			return err
		}

		score, err := a.ActionTable.GetScore(input.Action)
		if err != nil {
			return err
		}

		theAction := entity.NewGameAction(input.PlayerID, input.Minute, input.Action, score)
		match.Actions = append(match.Actions, *theAction)

		err = matchRepo.SaveActions(ctx, match, float64(score))
		if err != nil {
			return err
		}

		player, err := playerRepo.FindByID(ctx, input.PlayerID)
		if err != nil {
			return err
		}

		player.Price += float64(score)
		err = playerRepo.Update(ctx, player)
		if err != nil {
			return err
		}

		myTeam, err := myTeamRepo.FindByID(ctx, input.TeamID)
		if err != nil {
			return err
		}
		err = myTeamRepo.AddScore(ctx, myTeam, float64(score))
		if err != nil {
			return err
		}
		return nil
	})
}

func (a *ActionAddUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := a.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

func (a *ActionAddUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := a.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (a *ActionAddUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
