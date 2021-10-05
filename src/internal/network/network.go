package network

import (
	"bytes"
	"fmt"
	"encoding/gob"
	"encoding/json"
	"net"

)

type KDCRequest struct {
	Type string
	Request bytes.Buffer
}


func NewKDCRequest(req_type string, req *bytes.Buffer) (buf *bytes.Buffer) {
	print("---NewKDCRequest---")
	print(req.String())
	printi(req.Len())
	print("---NewKDCRequest Complete---")
	kdcreq := KDCRequest {req_type, *req}
	buf = KDCRequesttoBytes(kdcreq)
	return
}


func ToString(kdc KDCRequest) (string) {
	out, err := json.Marshal(kdc)
	if err != nil {
		panic(err)
	}
	return string(out)
}


func FromString(s string) (KDCRequest) {
	kdc := KDCRequest{}
	json.Unmarshal([]byte(s), &kdc)
	return kdc
}


func KDCRequesttoBytes(kdcreq KDCRequest) (buf *bytes.Buffer) {
	print(ToString(kdcreq))
	buf = new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(kdcreq)
	if err != nil {
		panic(err)
	}
	return
}

func BytestoKDCRequest(buf *bytes.Buffer) (kdcreq KDCRequest) {
	print("---BytestoKDCRequest: Buf---")
	print(buf.String())
	printi(buf.Len())
	print("---BytestoKDCRequest: Buf---")

	kdcreq = *new(KDCRequest)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(&kdcreq)
	if err != nil {
		panic(err)
	}

	print("---BytestoKDCRequest: kdcreq---")
	print(ToString(kdcreq))
	print(kdcreq.Type)
	print(kdcreq.Request.String())
	print("---BytestoKDCRequest: kdcreq---")
	return
}



func connect() (conn net.Conn){
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to 127.0.0.1:8000\n")
	return
}


func Read(conn net.Conn) (recv *bytes.Buffer) {
	tmp := make([]byte, 500)
	in, _ := conn.Read(tmp)
	recv = bytes.NewBuffer(tmp[:in+1])

	print("---Network Read---")
	print(recv.String())
	printi(recv.Len())
	printi(in)
	print("---Network Read---")


	return
}


func Write(conn net.Conn, send *bytes.Buffer) {
	sent, _ := conn.Write(send.Bytes())
	print("---Network Write---")
	print(send.String())
	printi(sent)
	print("---Network Write---")
}



func Send_Receive(send *bytes.Buffer) (recv *bytes.Buffer) {
	print("---Network SendReceive---")
	print(send.String())
	printi(send.Len())
	print("---Network SendReceive---")
	conn := connect()
	Write(conn, send)
	recv = Read(conn)
	return
}




func print(s string) {
	fmt.Printf("%s\n", s)
}
func printi(i int) {
	fmt.Printf("%d\n", i)
}
