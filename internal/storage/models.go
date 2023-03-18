package storage

import (
	"github.com/mr-linch/go-tg"
	"time"
)

type User struct {
	Id        tg.UserID   `json:"u_id"`
	Username  tg.Username `json:"username"`
	FirstName string      `json:"firstname"`
	LastName  string      `json:"second_name"`
}

type RequestRecord struct {
	User        User      `json:"u_id"`
	Id          int       `json:"r_id"`
	RequestType string    `json:"r_type"`
	RequestArgs string    `json:"r_args"`
	RequestTime time.Time `json:"r_time"`
}
