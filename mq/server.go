package mq

import (
	"os"

	"github.com/gocraft/dbr/dialect"
	dbr "github.com/gocraft/dbr/v2"
	"github.com/joho/godotenv"
	"github.com/muhammadisa/go-mq-boilerplate/mq/routes"
	"github.com/muhammadisa/go-mq-boilerplate/mq/utils/errhandler"
	"github.com/muhammadisa/godbconn"
	"github.com/muhammadisa/gorabbitmq"
)

// Run start running message queue service
func Run() {

	// Loading .env file
	err := godotenv.Load()
	errhandler.HandleError(err, true)

	// Load database credential env and use it
	db, err := godbconn.DBCred{
		DBDriver:   os.Getenv("DB_DRIVER"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}.Connect()
	errhandler.HandleError(err, true)
	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}
	conn.SetMaxOpenConns(10)
	session := conn.NewSession(nil)
	session.Begin()

	// Create connection and channel RabbitMQ
	connection, channel, err := gorabbitmq.Connector{
		Username: os.Getenv("MQ_USERNAME"),
		Password: os.Getenv("MQ_PASSWORD"),
		Host:     os.Getenv("MQ_HOST"),
		Port:     os.Getenv("MQ_PORT"),
	}.Dial()
	errhandler.HandleError(err, true)
	defer connection.Close()
	defer channel.Close()

	routes.MessageQueue{
		Sess:         session,
		MQConnection: connection,
		MQChannel:    channel,
	}.NewMQ()

}
