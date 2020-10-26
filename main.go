package main

import (
	"log"
	"os"

	"github.com/muhammadisa/go-mq-boilerplate/mq"
	"github.com/muhammadisa/gosqlexec"
	"github.com/urfave/cli"
)

func main() {
	qe := gosqlexec.GoSQLExec{
		AlterQuery:  "db/alter/alter_tables.sql",
		DropQuery:   "db/drop/drop_tables.sql",
		CustomQuery: "db/query/custom_query.sql",
		Schemas: []string{
			"db/schemas/foobars.sql",
		},
	}

	app := cli.NewApp()
	app.Name = "Message Broker Service"
	app.Usage = "Message broker service CLI tools"

	app.Commands = []cli.Command{
		gosqlexec.MigrateCommand(qe),
		gosqlexec.DropTablesCommand(qe),
		gosqlexec.AlterTablesCommand(qe),
		gosqlexec.CustomQueryExecCommand(qe),
		{
			Name:  "run-server",
			Usage: "Start Server",
			Action: func(c *cli.Context) error {
				mq.Run()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
