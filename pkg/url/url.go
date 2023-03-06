package url

import (
	"crypto/tls"

	"github.com/kooberetis/keeper/pkg/configReader"
	"github.com/kooberetis/keeper/pkg/utils"
)

// Url configuration structure
type config []struct {
	Host        string
	Port        string
	Description string
	CertInfo    []struct {
		CN      string
		Expired string
	}
}

// Return config structure with data
func Url() config {
	cfg := config(configReader.ReadConfig().Url)
	for item := range cfg {
		conn, err := tls.Dial("tcp", cfg[item].Host+":"+cfg[item].Port, nil)
		if err != nil {
			utils.WarningLog.Println("Server doesn't support SSL certificate err: " + err.Error())
		}
		for idx := range conn.ConnectionState().PeerCertificates {
			expiry := conn.ConnectionState().PeerCertificates[idx].NotAfter.Format("2006-01-02")
			cn := conn.ConnectionState().PeerCertificates[idx].Issuer.CommonName
			cfg[item].CertInfo = append(cfg[item].CertInfo[:idx], struct {
				CN      string
				Expired string
			}{
				CN:      cn,
				Expired: expiry,
			})
		}
	}
	return cfg
}
