package main

import (
	"fmt"
	"internal/crypt"
	"internal/tgt"
	"time"
)

func main() {
	key := "123456"
	username := "Skin"
	sessionkey := crypt.GenKey()
	lifespan, _ := time.ParseDuration("8h")
	timestamp := time.Now()
	tick := tgt.New(username, sessionkey, lifespan, timestamp)
	tgt.EncryptTGT(tick, key)
	fmt.Printf("hey\n")
}
