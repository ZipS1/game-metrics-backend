package amqp

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchangeName = "auth-service"
)

var (
	connConfig *ConnConfig
	timeout    time.Duration
)

func Init(brokerUri string, brokerTimeout time.Duration) (func(), error) {
	var cfg ConnConfig
	var initErr error

	timeout = brokerTimeout

	cfg.initOnce.Do(func() {
		cfg.conn, initErr = amqp.Dial(brokerUri)
		if initErr != nil {
			initErr = fmt.Errorf("failed to connect to message broker: %w", initErr)
			return
		}

		cfg.ch, initErr = cfg.conn.Channel()
		if initErr != nil {
			cfg.conn.Close()
			initErr = fmt.Errorf("failed to create channel: %w", initErr)
			return
		}

		if err := cfg.ch.ExchangeDeclare(
			exchangeName,
			"fanout",
			true,
			false,
			false,
			false,
			nil,
		); err != nil {
			cfg.ch.Close()
			cfg.conn.Close()
			initErr = fmt.Errorf("failed to declare exchange: %w", err)
			return
		}
	})

	connConfig = &cfg

	if initErr != nil {
		return nil, initErr
	}

	closeFunc := func() {
		cfg.closeOnce.Do(func() {
			if cfg.ch != nil {
				cfg.ch.Close()
			}
			if cfg.conn != nil {
				cfg.conn.Close()
			}
		})
	}

	return closeFunc, nil
}
