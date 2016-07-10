package mdm

import (
	"github.com/groob/plist"
	"github.com/micromdm/mdm/test"
	"testing"
)

func TestQueryResponseMac(t *testing.T) {
	response := &Response{}
	plistBuf := []byte(test.OSXElCapQueryResponses)
	err := plist.Unmarshal(plistBuf, response)

	if err != nil {
		t.Fatal(err)
	}
}

func TestQueryResponseIpadIOS8(t *testing.T) {
	response := &Response{}
	plistBuf := []byte(test.IOS8IpadQueryResponses)
	err := plist.Unmarshal(plistBuf, response)

	if err != nil {
		t.Fatal(err)
	}
}

func TestQueryResponseIphoneIOS8(t *testing.T) {
	response := &Response{}
	plistBuf := []byte(test.IOS8IphoneQueryResponses)
	err := plist.Unmarshal(plistBuf, response)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSecurityInfoMac(t *testing.T) {
	response := &Response{}
	plistBuf := []byte(test.OSXElCapSecurityInfoNoFDE)
	err := plist.Unmarshal(plistBuf, response)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSecurityInfoIpadIOS8(t *testing.T) {
	response := &Response{}
	plistBuf := []byte(test.IOS8IpadSecurityInfo)
	err := plist.Unmarshal(plistBuf, response)

	if err != nil {
		t.Fatal(err)
	}
}
