package usecase

import (
	"context"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/repository"
	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/service"
	"github.com/devfullcycle/imersao10-consolidacao/pkg/uow"
)

type MyTeamChoosePlayersInput struct {
	ID        string   `json:"my_team_id"`
	PlayersID []string `json:"players"`
}

type MyTeamChoosePlayersUseCase struct {
	Uow uow.UowInterface
}

func NewMyTeamChoosePlayersUseCase(uow uow.UowInterface) *MyTeamChoosePlayersUseCase {
	return &MyTeamChoosePlayersUseCase{
		Uow: uow,
	}
}

func (u *MyTeamChoosePlayersUseCase) Execute(ctx context.Context, input MyTeamChoosePlayersInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		playerRepo := u.getPlayerRepository(ctx)
		myTeamRepo := u.getMyTeamRepository(ctx)
		myTeam, err := myTeamRepo.FindByID(ctx, input.ID)
		if err != nil {
			return err
		}
		myPlayers, err := playerRepo.FindAllByIDs(ctx, myTeam.Players)
		if err != nil {
			return err
		}

		players, err := playerRepo.FindAllByIDs(ctx, input.PlayersID)
		if err != nil {
			return err
		}
		service.ChoosePlayers(myTeam, myPlayers, players)
		err = myTeamRepo.SavePlayers(ctx, myTeam)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *MyTeamChoosePlayersUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := u.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (u *MyTeamChoosePlayersUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := u.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
