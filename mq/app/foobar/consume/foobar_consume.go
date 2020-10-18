package consume

import (
	"encoding/json"
	"fmt"

	"github.com/muhammadisa/go-mq-boilerplate/mq/models"

	"github.com/muhammadisa/go-mq-boilerplate/mq/app/foobar"
	"github.com/muhammadisa/go-mq-boilerplate/mq/utils/errhandler"
	"github.com/muhammadisa/gorabbitmq"
	"github.com/streadway/amqp"
)

// FoobarConsumeHandler struct
type FoobarConsumeHandler struct {
	foobarUsecase foobar.Usecase
}

// NewFoobarConsumeHandler initialize consumer
func NewFoobarConsumeHandler(
	connection *amqp.Connection,
	channel *amqp.Channel,
	foobarUsecase foobar.Usecase,
) chan bool {
	usecase := &FoobarConsumeHandler{
		foobarUsecase: foobarUsecase,
	}

	messageFoobar, err := gorabbitmq.Queue{
		QueueName: "foobar",
		Consumer:  "",
		AutoAck:   false,
		Exclusive: false,
		NoLocal:   true,
		NoWait:    true,
		Args: amqp.Table{
			"x-dead-letter-exchange": "foobar_exchanges",
		},
	}.Consume(channel)
	errhandler.HandleError(err, false)

	foreverFoobar := make(chan bool)
	go func() {
		for d := range messageFoobar {
			attract(d, usecase)
		}
	}()
	return foreverFoobar
}

func attract(delivery amqp.Delivery, usecase *FoobarConsumeHandler) {
	var foobar *models.Foobar = &models.Foobar{}

	// unmarshaling body
	err := json.Unmarshal(delivery.Body, &foobar)
	if err != nil {
		delivery.Nack(false, true)
	}
	// attract to db
	err = usecase.foobarUsecase.Store(foobar)
	if err != nil {
		delivery.Reject(false)
	}
	// complete
	err = delivery.Ack(false)
	fmt.Println(foobar)
}
