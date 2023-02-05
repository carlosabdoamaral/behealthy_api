package rabbit

import (
	"fmt"

	"github.com/carlosabdoamaral/behealthy_api/common"
)

func StartConsuming() {
	msgs, err := common.RabbitChannel.Consume(
		common.RabbitQueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		common.LogError(err.Error())
		return
	}

	fmt.Println(msgs)

	forever := make(chan bool)
	go func() {
		fmt.Println("[*] Waiting for new messages")

		// for m := range msgs {
		// if m.Type == common.CREATE_OCCURRENCE_RABBIT_MSG_TYPE {
		// 	ConsumeCreateOccurrenceRequest(context.Background(), m.Body)
		// }
		// }
	}()

	<-forever
}
