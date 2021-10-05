package main

import (
	"bufio"
	"fmt"
	"internal/crypt"
	"internal/network"
	"internal/ticket"
	"internal/tgt"
	"os"
	"strings"
)

var kdc_addr string
var kdc_port int

func main() {
	kdc_addr = "127.0.0.1"
	kdc_port = 8000
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username_n, _ := reader.ReadString('\n')
	username := strings.TrimSuffix(username_n, "\n")
	fmt.Print("Password: ")
	password_n, _ := reader.ReadString('\n')
	password := strings.TrimSuffix(password_n, "\n")
	fmt.Print("Getting TGT from KDC\n")
	s_a, tickgt := login(username, password)

	fmt.Print("What resource would you like to access: ")
	resource_n, _ := reader.ReadString('\n')
	resource := strings.TrimSuffix(resource_n, "\n")
	//sharedkey, tick_b := requestTicket(username, resource, tickgt, s_a)
	_,_  = requestTicket(username, resource, tickgt, s_a)

	print("Received Ticket to the resource")
}



func login(username string, password string) (s_a crypt.Key, tickgt crypt.Encrypted){
	print("login: Starting")
	// Request the TGT from the KDC
	// KDC will respond with {S_A, TGT}K_A
	tgtRequest := tgt.NewTGTRequest(username)
	print("+++")
	print(tgtRequest.String())
	printi(tgtRequest.Len())
	print("+++")
	tgtResponse := network.Send_Receive(tgtRequest)

	// Derive K_A from the password
	key_a := crypt.DeriveK(password)

	// Decrypt the session key and the TGT
	decryptedTGTReply := crypt.Decrypt(*tgtResponse, key_a)
	tgtReply := tgt.BytestoTGTReply(&decryptedTGTReply)

	// Separate S_A and the TGT
	s_a = tgtReply.Session_A
	tickgt = tgtReply.Tickgt

	return
}


func requestTicket(username string, resource string, tickgt crypt.Encrypted, s_a crypt.SessionKey) (sharedkey crypt.SharedKey, tick_b crypt.Encrypted) {
	print("requestTicket: Starting")
	// Ask the KDC for a shared key
	ticketRequest := ticket.NewTicketRequest(username, resource, tickgt)
	print("+++")
	print(ticketRequest.String())
	printi(ticketRequest.Len())
	print("+++")
	ticketResponse := network.Send_Receive(ticketRequest)

	// Receive {B, K_AB, Tick_B}S_A from the server
	decryptedTicketReply := crypt.Decrypt(*ticketResponse, s_a)
	ticketReply := ticket.BytestoTicketReply(&decryptedTicketReply)

	// Parse the ticketReply object
	sharedkey = ticketReply.Sharedkey
	tick_b = ticketReply.Tick_B

	print("requestTicket: Complete")
	return
}


func print(s string) {
	fmt.Printf("%s\n", s)
}
func printi(i int) {
	fmt.Printf("%d\n", i)
}
