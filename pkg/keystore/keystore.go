package keystore

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

type CertInfo struct {
	Subject  interface{}
	NotAfter string
}

func getFiles() []string {
	var filenames []string
	files, _ := ioutil.ReadDir("/etc/keeper/")
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames
}

func GetCertInfo() []CertInfo {
	ci := []CertInfo{}
	files := getFiles()
	for _, file := range files {
		bs, _ := os.ReadFile("/etc/keeper/" + file)
		block, _ := pem.Decode(bs)
		if block == nil {
			log.Fatal("failed to parse PEM block containing the public key")
		}
		cert, _ := x509.ParseCertificate(block.Bytes)
		notAfterFormat := cert.NotAfter.Format("02.01.2006")
		ci = append(ci, struct {
			Subject  interface{}
			NotAfter string
		}{
			Subject:  cert.Subject.CommonName,
			NotAfter: string(notAfterFormat),
		})
	}
	return ci
}
