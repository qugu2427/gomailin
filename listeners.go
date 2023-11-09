package mailin

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
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
		if err != nil {
			logMsg := fmt.Sprintf("ERROR: failed TCP(no TLS) listen with %s (%s)", conn.RemoteAddr(), err)
			LogHandler(logMsg)
			continue
		}
		conn.SetDeadline(time.Now().Add(10 * time.Second))
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
		if err != nil {
			logMsg := fmt.Sprintf("ERROR: failed TCP/TLS listen with %s (%s)", conn.RemoteAddr(), err)
			LogHandler(logMsg)
			continue
		}
		conn.SetDeadline(time.Now().Add(10 * time.Second))
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
