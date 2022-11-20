package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	path, _ := os.Getwd()
	caCert, err := ioutil.ReadFile(path + "cert.pem")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	client := http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: cfg,
		},
	}

	resp, err := client.Get("/demo")
}
