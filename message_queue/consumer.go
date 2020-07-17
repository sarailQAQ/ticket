package message_queue

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"server/model"
)

func OpenConsumer() error{
	conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return  err
	}
	defer conn.Close()
	amqpChan, err := conn.Channel()
	if err != nil {
		return err
	}
	defer amqpChan.Close()

	queue, err := amqpChan.QueueDeclare("goodlist",true,false,false,false,nil)
	if err != nil {
		return  err
	}

	err = amqpChan.Qos(1,0,false)
	if err != nil {
		return err
	}
	msgChan,err := amqpChan.Consume(
			queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
	if err != nil {
		return err
	}

	stopChan := make(chan bool)

	go func() {
		for d := range msgChan {
			ord := &model.Order{}
			err = json.Unmarshal(d.Body,ord)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			model.NewOrde(*ord)
		}
	}()
	<-stopChan
	return nil
}
