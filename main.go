package main

import (
	"server/api"
	"server/message_queue"
	"server/model"
)

func main() {
	model.MysqlInit()
	go message_queue.OpenConsumer()
	api.Set_Router()
}
