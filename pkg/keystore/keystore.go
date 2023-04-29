package keystore

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func GetCertInfo() {
	bs, _ := os.ReadFile("cert.pem") // handle error

	block, _ := pem.Decode(bs)
	if block == nil {
		log.Fatal("failed to parse PEM block containing the public key")
	}

	cert, _ := x509.ParseCertificate(block.Bytes) // handle error

	log.Printf("Subject:   %q", cert.NotAfter)

	// Subject:   "CN=*.google.com"
}
