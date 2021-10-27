package storage

import (
	"time"
)

type Class struct {
	ID          int
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Capacity    int
	CreatedDate time.Time
}
