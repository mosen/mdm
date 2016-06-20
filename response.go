package mdm

// Response is an MDM Command Response
type Response struct {
	UDID           string
	UserID         *string `json:"user_id,omitempty" plist:"UserID,omitempty"`
	Status         string
	CommandUUID    string
	RequestType    string         `json:"request_type,omitempty" plist:",omitempty"`
	QueryResponses QueryResponses `json:"query_responses,omitempty" plist:",omitempty"`
}

// CommonQueryResponses has a list of query responses common to all device types
type CommonQueryResponses struct {
	UDID                  string
	Languages             []string          // ATV 6+
	Locales               []string          // ATV 6+
	DeviceID              string            // ATV 6+
	OrganizationInfo      map[string]string // IOS 7+
	LastCloudBackupDate   string            // IOS 8+
	AwaitingConfiguration bool              // IOS 9+
	// iTunes
	ITunesStoreAccountIsActive bool   // IOS 7+ OSX 10.9+
	ITunesStoreAccountHash     string // IOS 8+ OSX 10.10+

	// Device
	DeviceName                    string
	OSVersion                     string
	BuildVersion                  string
	ModelName                     string
	Model                         string
	ProductName                   string
	SerialNumber                  string
	DeviceCapacity                float32
	AvailableDeviceCapacity       float32
	BatteryLevel                  float32 // IOS 5+
	CellularTechnology            int     // IOS 4+
	IsSupervised                  bool    // IOS 6+
	IsDeviceLocatorServiceEnabled bool    // IOS 7+
	IsActivationLockEnabled       bool    // IOS 7+ OSX 10.9+
	IsDoNotDisturbInEffect        bool    // IOS 7+
	EASDeviceIdentifier           string  // IOS 7 OSX 10.9
	IsCloudBackupEnabled          bool    // IOS 7.1

	// Network
	BluetoothMAC string
	WiFiMAC      string
	EthernetMACs []string // Surprisingly works in IOS
}

// AtvQueryResponses contains AppleTV QuerryResponses
type AtvQueryResponses struct {
}

// IosQueryResponses contains iOS QueryResponses
type IosQueryResponses struct {
	IMEI                 string
	MEID                 string
	ModemFirmwareVersion string
	IsMDMLostModeEnabled bool // IOS 9.3
	MaximumResidentUsers int  // IOS 9.3

	// Network
	ICCID                    string // IOS
	CurrentCarrierNetwork    string
	SIMCarrierNetwork        string
	SubscriberCarrierNetwork string
	CarrierSettingsVersion   string
	PhoneNumber              string
	VoiceRoamingEnabled      bool
	DataRoamingEnabled       bool
	IsRoaming                bool
	PersonalHotspotEnabled   bool
	SubscriberMCC            string
	SubscriberMNC            string
	CurrentMCC               string
	CurrentMNC               string
}

// MacosQueryResponses contains macOS queryResponses
type MacosQueryResponses struct {
	OSUpdateSettings   map[string]string // OSX 10.11+
	LocalHostName      string            // OSX 10.11
	HostName           string            // OSX 10.11
	ActiveManagedUsers []string          // OSX 10.11
}

// QueryResponses is a DeviceInformation MDM Command Response
type QueryResponses struct {
	CommonQueryResponses
	MacosQueryResponses
	IosQueryResponses
	AtvQueryResponses
}
