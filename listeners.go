package main

// package mailin

import (
	"crypto/tls"
	"net"
)

func startListenerTCP(
	listenUrl string,
	connHandler ConnHandler,
	emailHandler EmailHandler) error {
	listener, err := net.Listen("tcp", listenUrl)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		// TODO PROD set read deadline
		if err != nil {
			continue
		}
		go connHandler(conn, emailHandler)
	}
}

func startListenerTLS(
	listenUrl string,
	serverCrtPath string,
	serverKeyPath string,
	connHandler ConnHandler,
	emailHandler EmailHandler) error {
	cert, err := tls.LoadX509KeyPair(serverCrtPath, serverKeyPath)
	if err != nil {
		panic(err)
	}

	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	listener, err := tls.Listen("tcp", listenUrl, &tlsConfig)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		// TODO PROD set read deadline
		if err != nil {
			continue
		}
		go connHandler(conn, emailHandler)
	}
}

func Listen(
	listenUrl string,
	tlsCrtPath string,
	tlsKeyPath string,
	emailHandler EmailHandler,
) error {
	if tlsCrtPath == "" && tlsKeyPath == "" {
		return startListenerTCP(listenUrl, handleConnection, emailHandler)
	} else {
		return startListenerTLS(listenUrl, tlsCrtPath, tlsKeyPath, handleConnection, emailHandler)
	}
}
