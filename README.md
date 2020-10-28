<h1 align="center">Welcome to go-mq-boilerplate 👋</h1>
<p>
  <img alt="documentation: yes" src="https://img.shields.io/badge/Documentation-Yes-green.svg" />
  <img alt="maintained: yes" src="https://img.shields.io/badge/Maintained-Yes-green.svg" />
</p>

>This project is boilerplate for Message Queueing Processing, in this project using MySQL Database, and using RabbitMQ Message Broker, and using AMQP for interacting with the RabbitMQ.



### Tech Stack

- Programming Language : Go
- Message Broker : RabbitMQ
- Database : MySQL
- Query Builder : DBR
- Architecture : Clean Architecture



### Getting Started

- Clone this project & Rename everything as your project

- Then run these command

  ```bash
  go mod tidy
  go mod download
  go mod verify
  ```

- Create .env file, here is the example

  ```bash
  DB_DRIVER="mysql"
  DB_HOST="localhost"
  DB_NAME="foobar"
  DB_PASSWORD="password"
  DB_PORT="3306"
  DB_USER="root"
  MQ_PORT="5672"
  MQ_USERNAME="guest"
  MQ_PASSWORd="guest"
  MQ_HOST="localhost"
  MODE="mq"
  ```

- Create your database, and set the name inside .env file

- Start your RabbitMQ service, and set the host,user,pass for connect to your RabbitMQ

- Compiling project & Migrate tables defined inside db/schemas folder

  ```bash
  go build -o mq.exe main.go
  mq.exe migrate --loadenv .env
  ```

- Run the server

  ```bash
  mq.exe run-server
  ```