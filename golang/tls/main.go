package main

import (
	"time"
)

func main() {
	go NewTLSServer()

	time.Sleep(time.Second)

	go NewTLSClient()

	select {}
}
