package userEntities

import "time"

type User struct {
	Id          string
	Name        string
	Mail        string
	PhoneNumber string
	Created     time.Time
	LastUpdated time.Time
}

