package keystore

import (
	"os"
	"reflect"
	"regexp"
	"testing"
)

var fineCert string = `-----BEGIN CERTIFICATE-----
MIICljCCAX4CCQDy7FCnVHvdlTANBgkqhkiG9w0BAQsFADANMQswCQYDVQQDDAJD
TjAeFw0yMzA0MjkxMzIzNDZaFw0zMzA0MjYxMzIzNDZaMA0xCzAJBgNVBAMMAkNO
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuRWlEhgSJQCKgTVe42Tp
akmohFN2Hq315UKTLkAcgDLLUt2cNJ5VkKvrFFrEqQ36sgUqtJb6RGPr/KcIlDL6
m9p4hXw/aW3/gECPwPmXomY6BwZPDe7JfwA59GR7WyKd0h8CKZv59N1RCEfaINqv
iCtkOgFHt4tdQZ/c7GqN1NLA0z/wP2o3OKlPdvaLWV/5KxCDvrCFqfi1tWuZk9ZB
RB/FzQNnVHzs4VDBldAVYc7rbAc8NjCur8ZCIHW6ECSAdl4PNYG668VoBv5o3DPm
zuaeAmC8igtemriM1dqWKetHsPzfbJIIXw2+CSCB8WSuc+ejfA2PB6RZQJB/0Mge
EwIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBxHbIv4bA3ho2foFd4WlrtjAtlnCM6
SFIvwjAB2J7p6Q/uxNzOKQUmLyLCy+Bg0OpvY7DJ2ihfH80UiiiECYbHkCiW195l
VWnhlmRQEAhfaixahEApodVefMhb+PfpRatOdKc95ptTgbLN7AcEveO8EcyaQQBt
4wb+Zw8d5gdH5o4tYspvHgetWdAnSqm9Ot4DxgsApR/41HLJ0Xr91zofooQ0nt/N
qmunodnqxRYjQ+1krja9btEJhBfKYR0nQ1s5Nd3V73qUKaZkPJRGtPbOChNPfULq
Hc2eh5uGacPKXdXJnD4aOYg00Otcml+wyIYeFnnV2ixUm4ZqUGgnmjEa
-----END CERTIFICATE-----
`
var wrongYaml string = `wrong: yaml`

func TestCertInfo(t *testing.T) {
	ci := CertInfo{}
	if reflect.TypeOf(ci.Subject).String() != "string" {
		t.Errorf("Incorrect type for field: Subject. Expected field type: %s. Want: string", reflect.TypeOf(ci.Subject).String())
	}
	if reflect.TypeOf(ci.NotAfter).String() != "string" {
		t.Errorf("Incorrect type for field: NotAfter. Expected field type: %s. Want: string", reflect.TypeOf(ci.Subject).String())
	}
}

func TestGetFiles(t *testing.T) {
	os.MkdirAll(certPath, os.ModePerm)
	os.WriteFile(certPath+"fine.pem", []byte(fineCert), 0644)
	os.WriteFile(certPath+"wrong.yaml", []byte(wrongYaml), 0644)
	ci := getFiles()
	if len(ci) == 0 {
		t.Errorf("Error with getFiles. Function return empty array.")
	}
	for _, wrong := range ci {
		if matched, _ := regexp.MatchString(".yaml", wrong); matched {
			t.Errorf("Error with getFiles. Wrong file extensions in array. Want: .pem, expected: %s", wrong)
		}

	}

}

func TestGetCertInfo(t *testing.T) {
	ci := GetCertInfo()
	if len(ci) == 0 {
		t.Errorf("Error with GetCertInfo. Function return empty array.")
	}
}
