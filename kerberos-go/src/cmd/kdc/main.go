package main

import (
	"bytes"
	"fmt"
	"internal/crypt"
	"internal/network"
	"internal/tgt"
	"internal/ticket"
	"net"
	"time"
)

var kdc_port int
var key_KDC string
var database map[string]string

func main() {
	kdc_port = 88
	key_KDC = "SecretKey1"
	database = make(map[string]string)
	addUser("Aditi", "Jain")
	addResource("fileserver", "password")

	fmt.Print("KDC is running\n")
	server()
}

func addUser(username string, password string) {
	database[username] = crypt.DeriveK(password)
}

func addResource(username string, password string) {
	addUser(username, password)
}


func genTGT(req *bytes.Buffer) (tgtReply *bytes.Buffer) {
	print("genTGT: Starting")

	// Parse the TGTRequest object
	tgtRequest := tgt.BytestoTGTRequest(req)
	username := tgtRequest.Username

	// Invent fresh S_A
	s_a := crypt.GenKey()

	// Lookup K_A in database
	key_a := database[username]

	// Generate the new TGT
	// TGT := {A, S_A, exptime}K_KDC
	lifespan, _ := time.ParseDuration("8h")
	tickgt := tgt.EncryptTGT(tgt.NewTGT(username, s_a, lifespan, time.Now()), key_KDC)

	// Create the TGTReply
	tgtRep := tgt.NewTGTReply(s_a, tickgt, key_a)
	tgtReply = &tgtRep
	return
}


func genTicket(req *bytes.Buffer) (ticketReply *bytes.Buffer) {
	print("genTicket: Starting")

	// Parse the TicketRequest object
	ticketRequest := ticket.BytestoTicketRequest(req)
	username := ticketRequest.Username
	resource := ticketRequest.Resource
	tickgt_encrypted := ticketRequest.Tickgt

	// Parse the TGT received
	tickgt := tgt.DecryptTGT(tickgt_encrypted, key_KDC)
	//tgtUsername := tickgt.Username
	sessionkey := tickgt.Sessionkey
	//lifespan := tickgt.Lifespan
	//timestamp := tickgt.Timestamp

	// TODO: verify the usernames match

	// Generate a fresh shared key
	key_AB := crypt.GenKey()

	// Generate tick_B
	// Look up the resource key
	key_b := database[resource]
	life, _ := time.ParseDuration("8h")
	tick_b := ticket.NewTicket(username, resource, key_AB, life, time.Now())
	tick_b_encrypted := ticket.EncryptTicket(tick_b, key_b)

	// Generate the ticketReply
	ticketReply = ticket.NewTicketReply(resource, key_AB, tick_b_encrypted, sessionkey)
	return
}

func server() {
	var response *bytes.Buffer
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		print("Got a connection")
		if err != nil {
			panic(err)
		}

		requestBuffer := network.Read(conn)

		print("---Server Loop---")
		print(requestBuffer.String())
		printi(requestBuffer.Len())
		print("---Server Loop---")

		// Parse the KDC request
		kdcrequest := network.BytestoKDCRequest(requestBuffer)

		//print(network.ToString(kdcrequest))
		req_type := kdcrequest.Type
		req_details := kdcrequest.Request
		print("parsed the network input")
		print(req_type)

		if req_type == "tgt" {
			response = genTGT(&req_details)
		} else if req_type == "ticket" {
			response = genTicket(&req_details)
		}

		network.Write(conn, response)
	}
}

func print(s string) {
	fmt.Printf("%s\n", s)
}
func printi(i int) {
	fmt.Printf("%d\n", i)
}
