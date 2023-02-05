package rabbit

import (
	"github.com/username/.../backend/common"
	"github.com/streadway/amqp"
)

func Connect() {
	conn, err := amqp.Dial(common.RabbitURL)
	if err != nil {
	}

	ch, err := conn.Channel()
	if err != nil {
	}

	common.RabbitConn = conn
	common.RabbitChannel = ch
}