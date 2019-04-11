package main

import (
	"io"
	"log"
	"net"
	"os"

	"github.com/golangq/q"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		q.Q("error", err)
		log.Printf("error %v", err)
		os.Exit(1)
	}
	ch := make(chan struct{})
	go func() {
		_, err := io.Copy(con, os.Stdin)
		q.Q("stdin copy ended ", err)
		ch <- struct{}{}
	}()
	go func() {
		_, err := io.Copy(os.Stdout, con)
		q.Q("stdout copy ended ", err)
		ch <- struct{}{}
	}()
	q.Q("done")
	<-ch
}
