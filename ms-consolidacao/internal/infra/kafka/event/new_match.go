package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/devfullcycle/imersao10-consolidacao/internal/usecase"
	"github.com/devfullcycle/imersao10-consolidacao/pkg/uow"
)

type ProcessNewMatch struct{}

func (p ProcessNewMatch) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.MatchInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	addNewMatchUsecase := usecase.NewMatchUseCase(uow)
	err = addNewMatchUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
