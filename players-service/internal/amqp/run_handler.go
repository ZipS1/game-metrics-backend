package amqp

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func RunHandler(exchange, queue, event string, timeout time.Duration, handler func(amqp.Delivery, zerolog.Logger), logger zerolog.Logger) error {
	ch := brokerState.ch
	if err := ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	q, err := ch.QueueDeclare(
		queue,
		false,
		false,
		true,
		false,
		map[string]interface{}{
			"x-message-ttl": int32(timeout.Milliseconds()),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	if err = ch.QueueBind(
		q.Name,
		event,
		exchange,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	messages, err := ch.Consume(
		q.Name,
		"activity-service",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to consume: %w", err)
	}

	go func() {
		for msg := range messages {
			handler(msg, logger)
		}
	}()

	return nil
}
