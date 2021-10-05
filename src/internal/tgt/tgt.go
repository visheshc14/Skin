package tgt

import (
	"bytes"
	"internal/crypt"
	"internal/network"
	"encoding/json"
	"encoding/gob"
	"fmt"
	"time"
)

type TGT struct {
	Username string
	Sessionkey crypt.SessionKey
	Lifespan time.Duration
	Timestamp time.Time
}

type TGTRequest struct {
	Username string
}

type TGTReply struct {
	Session_A crypt.SessionKey
	Tickgt crypt.Encrypted
}

func NewTGT(username string, sessionkey crypt.SessionKey, lifespan time.Duration, timestamp time.Time) TGT {
	tgt := TGT{username, sessionkey, lifespan, timestamp}
	return tgt
}


func NewTGTRequest(username string) (buf *bytes.Buffer) {
	request := TGTRequest{username}
	req := TGTRequesttoBytes(request)
	print(req.String())
	printi(req.Len())


	buf = network.NewKDCRequest("tgt", req)
	return
}

func NewTGTReply(sessionkey crypt.SessionKey, tickgt crypt.Encrypted, key_a crypt.Key) (buf bytes.Buffer){
	repobj := TGTReply {sessionkey, tickgt}
	buf0 := new(bytes.Buffer)
	encoder0 := gob.NewEncoder(&buf)
	encoder0.Encode(repobj)

	buf = crypt.Encrypt(*buf0, key_a)
	return
}


func ToString(tgt TGT) (string) {
	out, err := json.Marshal(tgt)
	if err != nil {
		panic(err)
	}
	return string(out)
}


func FromString(s string) (TGT) {
	tgt := TGT{}
	json.Unmarshal([]byte(s), &tgt)
	return tgt
}


func TGTtoBytes(tgt TGT) (buf *bytes.Buffer) {
	buf = new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(tgt)
	return
}

func BytestoTGT(buf *bytes.Buffer) (tickgt TGT) {
	tickgt = *new(TGT)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(tickgt)
	return
}


func TGTRequesttoBytes(request TGTRequest) (buf *bytes.Buffer) {
	buf = new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(request)
	return
}

func BytestoTGTRequest(buf *bytes.Buffer) (request TGTRequest) {
	request = *new(TGTRequest)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(request)
	return
}


func TGTReplytoBytes(reply TGTReply) (buf *bytes.Buffer) {
	buf = new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(reply)
	return
}

func BytestoTGTReply(buf *bytes.Buffer) (reply TGTReply) {
	reply = *new(TGTReply)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(reply)
	return
}



func EncryptTGT(tgt TGT, k crypt.Key) (encrypted crypt.Encrypted) {
	fmt.Printf("Encrypting the Ticket-Granting-Ticket: Starting\n")
	encrypted = crypt.Encrypt(*TGTtoBytes(tgt), k)
	fmt.Printf("Encrypting the Ticket-Granting-Ticket: Complete\n")
	return
}


func DecryptTGT(b crypt.Encrypted, k crypt.Key) (tgt TGT) {
	//fmt.Printf("Decrypting the Ticket-Granting-Ticket: Starting\n")
	decrypted := crypt.Decrypt(b, k)
	tgt = BytestoTGT(&decrypted)
	//fmt.Printf("Decrypting the Ticket-Granting-Ticket: Complete\n")
	return
}
func print(s string) {
	fmt.Printf("%s\n", s)
}
func printi(i int) {
	fmt.Printf("%d\n", i)
}
