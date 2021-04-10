package main

import (
	"crypto/x509"
	"fmt"
	"strings"
)

type Bundle struct {
	RootCA         x509.Certificate
	IntermediaryCA x509.Certificate
}

func (b *Bundle) CheckValidity() error {

	// TODO Axchtually
	// func (* Certifiate).Verify(chains [][]*Certificate, err, error){...}

	return fmt.Errorf("OH FUCK OH NO INVALID")

}

func main() {

	// Generate self-signed certificate
	// privkey, err := ecsda.GenerateKey

	// Parse and verifying x509 certificate chain

	//TODO Chain means that we need to use a linked list? Or... just use a slice w/ ?

	rootca := &x509.Certificate{}
	intermediaryca := &x509.Certificate{}

	certs := []*x509.Certificate{rootca, intermediaryca}

	// Here is the cert bundle
	// chain := make([]x509.Certificate, 2)
	// chain = append(certs, cas...)

	// Issuer (CA) of each certificate matches the subject of the next certificate in the chain
	certPool, err := x509.SystemCertPool()
	if err != nil {
		fmt.Println("Failed to parse system certificate pool!")
	}

	//
	//	fmt.Sprintf("Total Domains/Subjects found in system CA pool: %v", len(certPool.Subjects[0:]))

	//
	for _, cert := range certPool.Subjects() {
		var certstring []string
		for _, sub := range cert {
			//fmt.Printf(strings.Join(string(sub), " "))
			certstring = append(certstring, string(sub))
		}
		fmt.Printf("%v", strings.Join(certstring, ""))
		fmt.Println("------------------------")
	}
	fmt.Println(certs)
}
