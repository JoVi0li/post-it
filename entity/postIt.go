package entity

import "time"

type PostItEntity struct {
	Id          string
	CreatedAt   time.Time
	Title       string
	Description string
}
