package mdm

import "encoding/hex"

// CheckinCommand represents an MDM checkin command struct
type CheckinCommand struct {
	// MessageType can be either Authenticate,
	// TokenUpdate or CheckOut
	MessageType string
	Topic       string
	UDID        string
	auth
	update
}

// Authenticate Message Type
type auth struct {
	OSVersion    string
	BuildVersion string
	ProductName  string
	SerialNumber string
	IMEI         string
	MEID         string
}

// TokenUpdate Mesage Type
type update struct {
	Token                 hexData
	PushMagic             string
	UnlockToken           hexData
	AwaitingConfiguration bool
}

// DEPEnrollmentRequest is a request sent
// by the device to an MDM server during
// DEP Enrollment
type DEPEnrollmentRequest struct {
	Language string `plist:"LANGUAGE"`
	Product  string `plist:"PRODUCT"`
	Serial   string `plist:"SERIAL"`
	UDID     string `plist:"UDID"`
	Version  string `plist:"VERSION"`
	IMEI     string `plist:"IMEI,omitempty"`
	MEID     string `plist:"MEID,omitempty"`
}

// data decodes to []byte,
// we can then attach a string method to the type
// Tokens are encoded as Hex Strings
type hexData []byte

func (d hexData) String() string {
	return hex.EncodeToString(d)
}
