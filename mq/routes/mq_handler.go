package routes

import (
	dbr "github.com/gocraft/dbr/v2"
	"github.com/streadway/amqp"

	_foobarApi "github.com/muhammadisa/go-mq-boilerplate/mq/app/foobar/consume"
	_foobarRepo "github.com/muhammadisa/go-mq-boilerplate/mq/app/foobar/repository"
	_foobarUsecase "github.com/muhammadisa/go-mq-boilerplate/mq/app/foobar/usecase"
)

// MessageQueue struct
type MessageQueue struct {
	Sess         *dbr.Session
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
}

// IMQConfig interface
type IMQConfig interface {
	NewMQ()
}

// NewMQ message queue initialization
func (mq MessageQueue) NewMQ() {
	foreverFoobar := mq.initMQForFoobar()
	<-foreverFoobar
}

func (mq *MessageQueue) initMQForFoobar() chan bool {
	foobarRepo := _foobarRepo.NewORMFoobarRepo(mq.Sess)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo)
	return _foobarApi.NewFoobarConsumeHandler(mq.MQConnection, mq.MQChannel, foobarUsecase)
}
