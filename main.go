package main

import (
	"fmt"
)

func handleEmail(email Email) {
	fmt.Println(email)
}

func main() {

	/*
	   When crt and key are empty strings
	   we will listen over unencrypted TCP
	*/
	tlsCrt := ""
	tlsKey := ""

	err := Listen("0.0.0.0:25", tlsCrt, tlsKey, handleEmail)
	panic(err)
}
