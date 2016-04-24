package mdm

import (
	"fmt"

	"github.com/satori/go.uuid"
)

// CommandRequest represents an MDM command request
type CommandRequest struct {
	// Included with every payload
	RequestType string `json:"request_type"`
	UDID        string `json:"udid"`
	// DeviceInformation request
	Queries []string `json:"queries,omitempty"`
	InstallApplication
	AccountConfiguration
}

// Payload is an MDM payload
type Payload struct {
	CommandUUID string
	Command     *command
}

type command struct {
	RequestType string
	DeviceInformation
	InstallApplication
	AccountConfiguration
}

// InstallApplication is a InstallApplication MDM Comand
type InstallApplication struct {
	ITunesStoreID   int    `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	Identifier      string `plist:",omitempty" json:"identifier,omitempty"`
	ManifestURL     string `plist:",omitempty" json:"manifest_url,omitempty"`
	ManagementFlags int    `plist:",omitempty" json:"management_flags,omitempty"`
	NotManaged      bool   `plist:",omitempty" json:"not_managed,omitempty"`
	// TODO: add remaining optional fields
}

// DeviceInformation is a DeviceInformation MDM Command
type DeviceInformation struct {
	Queries []string `plist:",omitempty" json:"queries,omitempty"`
}

// AccountConfiguration is an MDM command to create a primary user on OS X
// It allows skipping the UI to set up a user.
type AccountConfiguration struct {
	SkipPrimarySetupAccountCreation     bool           `plist:",omitempty" json:"skip_primary_setup_account_creation,omitempty"`
	SetPrimarySetupAccountAsRegularUser bool           `plist:",omitempty" json:"skip_primary_setup_account_as_regular_user,omitempty"`
	AutoSetupAdminAccounts              []AdminAccount `plist:",omitempty" json:"auto_setup_admin_accounts,omitempty"`
}

// AdminAccount is the configuration for the
// Admin account created during Setup Assistant
type AdminAccount struct {
	ShortName    string `plist:"shortName" json:"short_name"`
	FullName     string `plist:"fullName,omitempty" json:"short_name,omitempty"`
	PasswordHash data   `plist:"passwordHash" json:"password_hash"`
	Hidden       bool   `plist:"hidden,omitempty" json:"hidden,omitempty"`
}

type data []byte

func newPayload(requestType string) *Payload {
	u := uuid.NewV4()
	return &Payload{u.String(),
		&command{RequestType: requestType}}
}

// NewPayload creates an MDM Payload
func NewPayload(request *CommandRequest) (*Payload, error) {
	requestType := request.RequestType
	payload := newPayload(requestType)
	switch requestType {
	case "DeviceInformation":
		payload.Command.DeviceInformation.Queries = request.Queries
	case "ProfileList",
		"SecurityInfo",
		"CertificateList",
		"OSUpdateStatus",
		"DeviceConfigured",
		"AvailableOSUpdates":
		return payload, nil
	case "InstallApplication":
		payload.Command.InstallApplication = request.InstallApplication
	case "AccountConfiguration":
		payload.Command.AccountConfiguration = request.AccountConfiguration
	default:
		return nil, fmt.Errorf("Unsupported MDM RequestType %v", requestType)
	}
	return payload, nil
}
