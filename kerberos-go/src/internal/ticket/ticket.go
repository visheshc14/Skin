package ticket

import (
	"bytes"
	"internal/crypt"
	"internal/network"
	"encoding/json"
	"encoding/gob"
	"fmt"
	"time"
)

type Ticket struct {
	UsernameA string
	UsernameB string
	Sharedkey crypt.SharedKey
	Lifespan time.Duration
	Timestamp time.Time
}

type TicketRequest struct {
	Username string
	Resource string
	Tickgt crypt.Encrypted
}

type TicketReply struct {
	Resource string
	Sharedkey crypt.SharedKey
	Tick_B crypt.Encrypted
}


func NewTicket(usernamea string, usernameb string, sharedkey crypt.SharedKey, lifespan time.Duration, timestamp time.Time) Ticket {
	ticket := Ticket {usernamea, usernameb, sharedkey, lifespan, timestamp}
	return ticket
}

func NewTicketRequest(username string, resource string, tickgt crypt.Encrypted) (buf *bytes.Buffer) {
	request := TicketRequest {username, resource, tickgt}
	req := new(bytes.Buffer)
	encoder := gob.NewEncoder(req)
	encoder.Encode(request)

	kdcrequest := network.NewKDCRequest("ticket", req)
	buf = new(bytes.Buffer)
	encoder2 := gob.NewEncoder(buf)
	encoder2.Encode(kdcrequest)
	return
}


func NewTicketReply(resource string, sharedkey crypt.SharedKey, tick_b crypt.Encrypted, sessionkey crypt.SessionKey) (rep *bytes.Buffer) {
	repobj := TicketReply {resource, sharedkey, tick_b}
	buf0 := new(bytes.Buffer)
	encoder0 := gob.NewEncoder(buf0)
	encoder0.Encode(repobj)

	reply := crypt.Encrypt(*buf0, sessionkey)
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(reply)
	return
}


func ToString(tick Ticket) (string) {
	out, err := json.Marshal(tick)
	if err != nil {
		panic(err)
	}
	return string(out)
}


func FromString(s string) (Ticket) {
	tick := Ticket{}
	json.Unmarshal([]byte(s), &tick)
	return tick
}


func TickettoBytes(tick Ticket) (buf *bytes.Buffer) {
	buf = new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(tick)
	return
}

func BytestoTicket(buf *bytes.Buffer) (tick Ticket) {
	tick = *new(Ticket)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(tick)
	return
}


func TicketRequesttoBytes(request TicketRequest) (buf *bytes.Buffer) {
	buf = new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(request)
	return
}

func BytestoTicketRequest(buf *bytes.Buffer) (request TicketRequest) {
	request = *new(TicketRequest)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(request)
	return
}


func TicketReplytoBytes(reply TicketReply) (buf *bytes.Buffer) {
	buf = new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(reply)
	return
}

func BytestoTicketReply(buf *bytes.Buffer) (reply TicketReply) {
	reply = *new(TicketReply)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(reply)
	return
}


func EncryptTicket(tick Ticket, k crypt.Key) (encrypted crypt.Encrypted) {
	fmt.Printf("Encrypting the Ticket: Starting\n")
	encrypted = crypt.Encrypt(*TickettoBytes(tick), k)
	fmt.Printf("Encrypting the Ticket: Complete\n")
	return encrypted
}


func DecryptTicket(b crypt.Encrypted, k crypt.Key) (tick Ticket) {
	fmt.Printf("Decrypting the Ticket: Starting\n")
	decrypted := crypt.Decrypt(b, k)
	tick = BytestoTicket(&decrypted)
	fmt.Printf("Decrypting the Ticket: Complete\n")
	return tick
}
