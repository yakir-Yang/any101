package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func NewTLSClient() {
	// Used for TLS server to verify whether the TLS client is ok
	cert, err := tls.LoadX509KeyPair("certs/client/client.crt", "certs/client/client.key")
	if err != nil {
		log.Println(err)
		return
	}

	// Used for TLS client to verify whether the TLS server is trust
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

	conf := &tls.Config{
		RootCAs: certPool,
		//InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: false,
	}

	log.Println("CLIENT: connecting to 127.0.0.1:4443\n")

	conn, err := tls.Dial("tcp", "127.0.0.1:4443", conf)
	if err != nil {
		log.Println("CLIENT: ", err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println("CLIENT: ", err)
		return
	}

	buf := make([]byte, 100)

	n, err = conn.Read(buf)
	if err != nil {
		log.Println("CLIENT: ", err)
		return
	}

	println("CLIENT: ", string(buf[:n]))
}
