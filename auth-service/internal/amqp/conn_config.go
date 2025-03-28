package amqp

import (
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConnConfig struct {
	ch        *amqp.Channel
	initOnce  sync.Once
	closeOnce sync.Once
	conn      *amqp.Connection
}

func (c *ConnConfig) isInitialized() bool {
	if c == nil {
		return false
	}
	return c.ch != nil && !c.ch.IsClosed()
}
