package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Ticket struct {
	gorm.Model
	Name string
	Begin time.Time
	End time.Time
	Site string
	Total int
	Left int
}

type Order struct {
	gorm.Model
	Uid uint
	Fid uint
}

func NewOrde(ord Order) error {
	var ticket Ticket
	DB.Find(&ticket).Where("id=?",ord.Fid)
	DB.Model(&ticket).Update("left",ticket.Left-1)
	DB.Create(&ord)
	return nil
}

func NewTicket(ticket Ticket) error {
	db := DB.Create(&ticket)
	return db.Error
}
