package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

//TODO Generating chain
// using parent field: x509.CreateCertificate(io.Reader, template, partent *Certificate,...)

// Generating self-signed certificates

func main() {

	// Make new private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate private key!: %v", err)
	}

	// Create certificate template
	//
	// This is to get the largest possible value of a _signed_ integer (int64).
	// We can do this by taking a int64 and shifting it left as far as possible.
	//
	// tldr; this is really just 2^128
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatal("Failed to generate serial number!")
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Tyrell Corp"},
		},
		DNSNames:    []string{"localhost"},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(3 * time.Hour),
		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	// x509.CreateCertificate() returns slice of certifcate in DER encoding
	// Passing the template in the parent field denotes a self-signing
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatalf("Failed to generate certificate: %v", err)
	}

	// pem = encode(DER-encoded-data)
	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if pemCert == nil {
		log.Fatal("Failed to encode cert to PEM")
	}

	// Serialize certificate to file
	err = os.WriteFile("cert.pem", pemCert, 0644)
	if err != nil {
		log.Fatalf("Failed to write PEM-encoded cert to file: %v", err)
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}

	// Serialize private key to file
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	err = os.WriteFile("key.pem", pemKey, 0644)
	if err != nil {
		log.Fatalf("Failed to write PEM-encoded key to file: %v", err)
	}

	// Now, use it for a HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprint(w, "WOOP")
	})

	srv := &http.Server{
		Addr:    "127.0.0.1:8765",
		Handler: mux,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	err = srv.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("HTTPTLS server failed to start: %v", err)
	}

	//TESTING NOTES

	// This will yell...
	// curl --tlsv1.3 -v  https://localhost:8765

	// ... unless we provide the cert
	// curl --tlsv1.3 --cacert cert.pem -v https://localhost:8765

}
