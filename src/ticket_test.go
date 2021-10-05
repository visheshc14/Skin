package main

import (
	"fmt"
	"internal/crypt"
	"internal/ticket"
	"time"
)

func main() {
	key := "123456"
	usernamea := "Skin"
	usernameb := "Evangelion"
	sharedkey := crypt.GenKey()
	lifespan, _ := time.ParseDuration("8h")
	timestamp := time.Now()
	tick := ticket.New(usernamea, usernameb, sharedkey, lifespan, timestamp)
	ticket.EncryptTicket(tick, key)
	fmt.Printf("hey\n")
}
