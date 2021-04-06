package utils

import (
	"math/rand"
	"time"
)

func InitRand() {
	s := rand.Seed(time.Now().UnixNano())
}
