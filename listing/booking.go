package listing

import (
	"time"
)

type Booking struct {
	ID         int       `json:"id"`
	MemberName string    `json:"name"`
	ClassID    int       `json:"class_id"`
	ClassName  string    `json:"class_name"`
	Date       time.Time `json:"date"`
}
