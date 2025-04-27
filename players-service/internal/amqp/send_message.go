package amqp

import (
	"context"
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func SendMessage(event string, payload map[string]interface{}, logger zerolog.Logger) error {
	if !brokerState.isInitialized() {
		return errors.New("AMQP connection is not initialized")
	}

	messageBody, err := json.Marshal(payload)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to serialize payload to JSON")
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), brokerTimeout)
	defer cancel()

	err = brokerState.ch.PublishWithContext(ctx,
		exchangeName,
		event,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBody,
		},
	)

	if err != nil {
		logger.Error().Str("event", event).Err(err).Msg("Failed to publish message")
		return err
	}

	logger.Info().Str("event", event).Msg("Message published successfully")
	return nil
}
