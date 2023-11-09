# gomailin
A simple golang module for accepting email sent via SMTP.

## A Simple Example
```go
package main

import mailin, "github.com/qugu2427/gomailin"

func handleEmail(email mailin.Email) {
	fmt.Println(email)
}

func main() {

    /* 
        When crt and key are empty strings
        we will listen over unencrypted TCP
    */
    tlsCrt := ""
    tlsKey := ""

	err := mailin.Listen("0.0.0.0:25", tlsCrt, tlsKey, handleEmail)
	panic(err)
}
```
**IMPORTANT**: simple listen performs NO AUTENTICATION
