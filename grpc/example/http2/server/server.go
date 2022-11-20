package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	l, err := net.Listen("tcp", ":443")
	if err != nil {
		panic(err)
	}

	path, _ := os.Getwd()
	certFile := path + "/cert.pem"
	keyFile := path + "/key.pem"
	log.Fatal(http.ServeTLS(l, h, certFile, keyFile))
}
