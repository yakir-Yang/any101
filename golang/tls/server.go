package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
)

func NewTLSServer() {
	// Used for TLS client to verify whether the TLS server is ok
	cert, err := tls.LoadX509KeyPair("certs/server/server.crt", "certs/server/server.key")
	if err != nil {
		log.Println(err)
		return
	}

	// Used for TLS server to verify whether the TLS client is trust
	caCert, err := ioutil.ReadFile("certs/ca/ca.crt")
	if err != nil {
		log.Println(err)
		return
	}

	certPool := x509.NewCertPool()

	ok := certPool.AppendCertsFromPEM(caCert)
	if !ok {
		log.Println("failed to parse client root certificate")
		return
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	ln, err := tls.Listen("tcp", ":4443", config)
	if err != nil {
		log.Println("SERVER: ", err)
		return
	}
	defer ln.Close()

	log.Println("SERVER: Listen at 127.0.0.1:4443")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("SERVER: ", err)
			return
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println("SERVER: ", err)
			return
		}

		println("SERVER: ", msg)

		_, err = conn.Write([]byte("world\n"))
		if err != nil {
			log.Println("SERVER: ", err)
			return
		}
	}
}
