package mq

import (
	"os"
	"strconv"

	"github.com/muhammadisa/go-mq-boilerplate/mq/routes"

	"github.com/joho/godotenv"
	"github.com/muhammadisa/go-mq-boilerplate/mq/utils/errhandler"
	"github.com/muhammadisa/gorabbitmq"
	"github.com/muhammadisa/gormdbcon"
)

// Run start running message queue service
func Run() {

	// Loading .env file
	err := godotenv.Load()
	errhandler.HandleError(err, true)

	// Load database credential env and use it
	db, err := gormdbcon.DBCredential{
		DBDriver:   os.Getenv("DB_DRIVER"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}.Connect()
	errhandler.HandleError(err, true)

	// Load debuging mode env
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	errhandler.HandleError(err, true)
	db.LogMode(debug)

	// Migrate and checking table fields changes
	Seed{DB: db}.Migrate()

	// Create connection and channel RabbitMQ
	connection, channel, err := gorabbitmq.Connector{
		Username: "guest",
		Password: "guest",
		Host:     "localhost",
		Port:     "5672",
	}.Dial()
	errhandler.HandleError(err, true)
	defer connection.Close()
	defer channel.Close()

	routes.MessageQueue{
		DB:           db,
		MQConnection: connection,
		MQChannel:    channel,
	}.NewMQ()

}
