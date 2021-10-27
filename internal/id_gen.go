package internal

import (
	"math/rand"
	"time"
)

//the simplest version possible
func GetID() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10000
	return (rand.Intn(max-min) + min)
}
