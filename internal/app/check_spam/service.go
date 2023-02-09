package check_spam

import (
	"math/rand"
	"time"
)

type Implementation struct {
}

func NewService() *Implementation {
	rand.Seed(time.Now().UnixNano())
	return &Implementation{}
}
