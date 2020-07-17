package message_queue

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"math/rand"
	"server/model"
	"time"
)

func NewOrder(ord model.Order) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()
	amqpChan, err := conn.Channel()
	if err != nil {
		return err
	}
	defer amqpChan.Close()

	queue, err := amqpChan.QueueDeclare("goodlist",true,false,false,false,nil)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	body, err := json.Marshal(ord)
	if err != nil {
		return err
	}
	err = amqpChan.Publish("",queue.Name,false,false,amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:		   body,
	})
	if err != nil {
		return err
	}
	return nil
}
