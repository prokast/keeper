package keystore

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var certPath string = "/mnt/keeper/"

type CertInfo struct {
	Subject  string
	NotAfter string
}

func getFiles() []string {
	var filenames []string
	files, _ := ioutil.ReadDir(certPath)
	for _, file := range files {
		if matched, _ := regexp.MatchString(".pem", file.Name()); matched {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames
}

func GetCertInfo() []CertInfo {
	ci := []CertInfo{}
	files := getFiles()
	for _, file := range files {
		bs, _ := os.ReadFile(certPath + file)
		block, _ := pem.Decode(bs)
		if block == nil {
			log.Fatal("failed to parse PEM block containing the public key")
		}
		cert, _ := x509.ParseCertificate(block.Bytes)
		notAfterFormat := cert.NotAfter.Format("02.01.2006")
		ci = append(ci, struct {
			Subject  string
			NotAfter string
		}{
			Subject:  cert.Subject.CommonName,
			NotAfter: string(notAfterFormat),
		})
	}
	return ci
}
