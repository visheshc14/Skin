package crypt

import (
	"bytes"
	"fmt"
	"math/rand"
)

type Key = string
type SharedKey = Key
type SessionKey = string
type Encrypted = bytes.Buffer

func DeriveK(password string) (output Key) {
	// Just reverse it
	runes := []rune(password)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	output = string(runes)
	return
}

func GenKey() (output Key) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var n = 6
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
	output = string(b)
	print("In GenKey")
	print(output)
	return
}


func Encrypt(input bytes.Buffer, key Key) (output Encrypted) {
	fmt.Printf("Encrypt: Starting\n")
	output = input
	//output = xor(input, key)
	fmt.Printf("Encrypt: Complete\n")
	return
}

func Decrypt(input Encrypted, key Key) (output bytes.Buffer) {
	//fmt.Printf("Decrypt: Starting\n")
	output = input
	//output = xor(input, key)
	//fmt.Printf("Decrypt: Complete\n")
	return
}

func xor(input string, key Key) (output string) {
	output = ""
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i % len(key)])
	}

	return
}


func print(s string) {
	fmt.Printf("%s\n", s)
}
