package amqp

import (
	"context"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func SendMessage(message []byte, logger zerolog.Logger) error {
	if !connConfig.isInitialized() {
		return errors.New("AMQP connection is not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := connConfig.ch.PublishWithContext(ctx,
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)

	if err != nil {
		logger.Error().Err(err).Msg("Failed to publish message")
		return err
	}

	logger.Info().Msg("Message published successfully")
	return nil
}
