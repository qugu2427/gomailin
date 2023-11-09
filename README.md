# gomailin
A simple golang module for accepting email sent via SMTP.<br>
Mailin provides two ways to listen for emails: ```SimpleListen()``` and ```Listen()```

## SimpleListen()
Sometimes we do not care about the nitty-gritty details of email... Sometimes we just want to recieve emails!<br>
SimpleListen() performs NO AUTHENTICATION. SimpleListen() does not check whether emails are forged, properly formatted, spam, etc.

```go
    mailin.SimpleListen(
        listenUrl, // host and port to listen for emails (ex: "0.0.0.0:25")
        tlsCrt, 
        tlsKey, // if tlsCrt and tlsKey are both "", mailin will listen over unencrypted TCP
        emailHandler // a func(mailin.Email) to handle recieved emails
        )
```

## A SimpleListen Example
```go
package main

import mailin, "github.com/qugu2427/go-mailin"

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

	err := mailin.SimpleListen("0.0.0.0:25", tlsCrt, tlsKey, handleEmail)
	panic(err)
}
```
**IMPORTANT**: simple listen performs NO AUTENTICATION
