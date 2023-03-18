package storage

import "time"

type User struct {
	Id int
}

type RequestRecord struct {
	User        User      `json:"user"`
	Id          int       `json:"id"`
	RequestTime time.Time `json:"request_time"`
	RequestType string    `json:"request_type"`
	RequestArgs string    `json:"request_args"`
	Response    string    `json:"response"`
}
