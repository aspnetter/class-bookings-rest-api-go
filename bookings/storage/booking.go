package storage

import (
	"time"
)

type Booking struct {
	ID          int
	MemberName  string
	ClassID     int
	Date        time.Time
	CreatedDate time.Time
}
