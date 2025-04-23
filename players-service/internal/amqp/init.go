package amqp

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchangeName = "players-service"
)

var (
	brokerState   *amqpState
	brokerTimeout time.Duration
)

func Init(uri string, timeout time.Duration) (func(), error) {
	var state amqpState
	var initErr error

	brokerTimeout = timeout
	state.initOnce.Do(func() {
		state.conn, initErr = amqp.Dial(uri)
		if initErr != nil {
			initErr = fmt.Errorf("failed to connect to message broker: %w", initErr)
			return
		}

		state.ch, initErr = state.conn.Channel()
		if initErr != nil {
			state.conn.Close()
			initErr = fmt.Errorf("failed to create channel: %w", initErr)
			return
		}

		if err := state.ch.ExchangeDeclare(
			exchangeName,
			"fanout",
			true,
			false,
			false,
			false,
			nil,
		); err != nil {
			state.ch.Close()
			state.conn.Close()
			initErr = fmt.Errorf("failed to declare exchange: %w", err)
			return
		}
	})

	brokerState = &state

	if initErr != nil {
		return nil, initErr
	}

	return getCloseFunc(&state), nil
}

func getCloseFunc(cfg *amqpState) func() {
	return func() {
		cfg.closeOnce.Do(func() {
			if cfg.ch != nil {
				cfg.ch.Close()
			}
			if cfg.conn != nil {
				cfg.conn.Close()
			}
		})
	}
}
