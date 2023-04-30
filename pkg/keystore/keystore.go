package keystore

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type CertInfo struct {
	Subject  interface{}
	NotAfter time.Time
}

func getFiles() []string {
	var filenames []string
	files, _ := ioutil.ReadDir("./etc/certs")
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	fmt.Println(filenames)
	return filenames
}

func GetCertInfo() {
	ci := []CertInfo{}
	files := getFiles()
	for _, file := range files {
		bs, _ := os.ReadFile("/Users/traksel/git/go/src/keeper/etc/certs/" + file)
		block, _ := pem.Decode(bs)
		if block == nil {
			log.Fatal("failed to parse PEM block containing the public key")
		}
		cert, _ := x509.ParseCertificate(block.Bytes)
		ci = append(ci, struct {
			Subject  interface{}
			NotAfter time.Time
		}{
			Subject:  cert.Subject.CommonName,
			NotAfter: cert.NotAfter,
		})
	}
	// handle error

	log.Printf("Subject:   %q", ci)

	// Subject:   "CN=*.google.com"
}
