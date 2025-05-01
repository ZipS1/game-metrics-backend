package amqp

import (
	"context"
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func SendMessage(event string, payload map[string]interface{}, logger zerolog.Logger) {
	if !brokerState.isInitialized() {
		err := errors.New("AMQP connection is not initialized")
		logger.Error().Err(err).Send()
	}

	messageBody, err := json.Marshal(payload)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to serialize payload to JSON")
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
	}

	logger.Info().Str("event", event).Msg("Message published successfully")
}
