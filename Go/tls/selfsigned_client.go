package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	certPool := x509.NewCertPool()

	file, err := os.Open("cert.pem")
	if err != nil {
		log.Fatalf("failed to open certificate file for reading: %v", err)
	}
	pemCert, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read certfile: %v", err)
	}

	ok := certPool.AppendCertsFromPEM(pemCert)
	if !ok {
		log.Fatalf("client: failed to parse cert!")
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool, // Add certPool with server cert to client's trusted CAs
			},
		},
	}

	resp, err := client.Get("https://localhost:8765")
	if err != nil {
		log.Fatalf("client: failed to dail server: %v", err)
	}
	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)
	fmt.Println(string(html))
}
