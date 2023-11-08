# go-mailin
A simple golang module for accepting email sent via SMTP.

## A Simple Example
```go
package main

import mailin, "github.com/qugu2427/go-mailin"

func handleEmail(e mailin.Email) {
	fmt.Println(email)
}

func main() {

    /* 
        When crt and key are empty strings
        we will listen over unencrypted TCP
    */
    tlsCrt := ""
    tlsKet := ""

	err := mailin.Listen("0.0.0.0:25", "", "", handleEmail)
	panic(err)
}
```
\*\*\* **IMPORTANT**: mailin performs NO VERIFICATION by default \*\*\*

## A More Secure Example
TODO
