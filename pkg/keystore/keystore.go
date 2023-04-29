package keystore

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

type certInfo struct {
	subject  struct {
        notAfter string
    }
}

func getFiles() []string {
	var filenames []string
	files, _ := ioutil.ReadDir("/tmp/")
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames
}

func GetCertInfo() {
    for
	bs, _ := os.ReadFile("cert.pem") // handle error

	block, _ := pem.Decode(bs)
	if block == nil {
		log.Fatal("failed to parse PEM block containing the public key")
	}

	cert, _ := x509.ParseCertificate(block.Bytes) // handle error

	log.Printf("Subject:   %q", cert.NotAfter)

	// Subject:   "CN=*.google.com"
}
