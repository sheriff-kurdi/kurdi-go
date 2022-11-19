package entities

import "time"

type Entity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
