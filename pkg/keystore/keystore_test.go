package keystore

import (
	"reflect"
	"testing"
)

func TestCertInfo(t *testing.T) {
	ci := CertInfo{}
	if reflect.TypeOf(ci.Subject).String() != "string" {
		t.Errorf("Incorrect type for field: Subject. Expected field type: %s. Want: string", reflect.TypeOf(ci.Subject).String())
	}
	if reflect.TypeOf(ci.NotAfter).String() != "string" {
		t.Errorf("Incorrect type for field: NotAfter. Expected field type: %s. Want: string", reflect.TypeOf(ci.Subject).String())
	}
}
