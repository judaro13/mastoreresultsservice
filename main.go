package main

import (
	"errors"
	"fmt"
	"judaro13/miaguila/storeresultsservice/process"
	"judaro13/miaguila/storeresultsservice/store"
	"log"
	"os"

	"github.com/judaro13/masharedmodels/utils"

	"github.com/streadway/amqp"
)

func main() {

	validateEnvVars()

	db := store.ConnectToDB()

	conn, err := amqp.Dial(os.Getenv("RABBIT_URL"))
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		os.Getenv("RABBIT_STORE_DATA_QUEUE"), "", true, // auto-ackc
		false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Print("\n processing message.....")
			err := process.Data(db, d.Body)
			if err != nil {
				utils.Error(err)
				d.Ack(false)
				fmt.Print("error\n")
			}
			fmt.Print("done\n")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func validateEnvVars() {
	vars := []string{"DATABASE_URL", "RABBIT_URL", "RABBIT_STORE_DATA_QUEUE"}
	for _, val := range vars {
		if len(val) == 0 {
			panic(errors.New("not found " + val + " environment variable"))
		}
	}
}
