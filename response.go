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
	UDID                  string            `json:"udid"`
	Languages             []string          `json:"languages,omitempty"`              // ATV 6+
	Locales               []string          `json:"locales,omitempty"`                // ATV 6+
	DeviceID              string            `json:"device_id"`                        // ATV 6+
	OrganizationInfo      map[string]string `json:"organization_info,omitempty"`      // IOS 7+
	LastCloudBackupDate   string            `json:"last_cloud_backup_date,omitempty"` // IOS 8+
	AwaitingConfiguration bool              `json:"awaiting_configuration"`           // IOS 9+
	// iTunes
	ITunesStoreAccountIsActive bool   `json:"itunes_store_account_is_active"` // IOS 7+ OSX 10.9+
	ITunesStoreAccountHash     string `json:"itunes_store_account_hash"`      // IOS 8+ OSX 10.10+

	// Device
	DeviceName                    string  `json:"device_name"`
	OSVersion                     string  `json:"os_version"`
	BuildVersion                  string  `json:"build_version"`
	ModelName                     string  `json:"model_name"`
	Model                         string  `json:"model"`
	ProductName                   string  `json:"product_name"`
	SerialNumber                  string  `json:"serial_number"`
	DeviceCapacity                float32 `json:"device_capacity"`
	AvailableDeviceCapacity       float32 `json:"available_device_capacity"`
	BatteryLevel                  float32 `json:"battery_level,omitempty"`           // IOS 5+
	CellularTechnology            int     `json:"cellular_technology,omitempty"`     // IOS 4+
	IsSupervised                  bool    `json:"is_supervised"`                     // IOS 6+
	IsDeviceLocatorServiceEnabled bool    `json:"is_device_locator_service_enabled"` // IOS 7+
	IsActivationLockEnabled       bool    `json:"is_activation_lock_enabled"`        // IOS 7+ OSX 10.9+
	IsDoNotDisturbInEffect        bool    `json:"is_dnd_in_effect"`                  // IOS 7+
	EASDeviceIdentifier           string  `json:"eas_device_identifier"`             // IOS 7 OSX 10.9
	IsCloudBackupEnabled          bool    `json:"is_cloud_backup_enabled"`           // IOS 7.1

	// Network
	BluetoothMAC string   `json:"bluetooth_mac"`
	WiFiMAC      string   `json:"wifi_mac"`
	EthernetMACs []string `json:"ethernet_macs"` // Surprisingly works in IOS
}

// AtvQueryResponses contains AppleTV QueryResponses
type AtvQueryResponses struct {
}

// IosQueryResponses contains iOS QueryResponses
type IosQueryResponses struct {
	IMEI                 string `json:"imei"`
	MEID                 string `json:"meid"`
	ModemFirmwareVersion string `json:"modem_firmware_version"`
	IsMDMLostModeEnabled bool   `json:"is_mdm_lost_mode_enabled,omitempty"` // IOS 9.3
	MaximumResidentUsers int    `json:"maximum_resident_users"`             // IOS 9.3

	// Network
	ICCID                    string `json:"iccid,omitempty"` // IOS
	CurrentCarrierNetwork    string `json:"current_carrier_network,omitempty"`
	SIMCarrierNetwork        string `json:"sim_carrier_network,omitempty"`
	SubscriberCarrierNetwork string `json:"subscriber_carrier_network,omitempty"`
	CarrierSettingsVersion   string `json:"carrier_settings_version,omitempty"`
	PhoneNumber              string `json:"phone_number,omitempty"`
	VoiceRoamingEnabled      bool   `json:"voice_roaming_enabled,omitempty"`
	DataRoamingEnabled       bool   `json:"data_roaming_enabled,omitempty"`
	IsRoaming                bool   `json:"is_roaming,omitempty"`
	PersonalHotspotEnabled   bool   `json:"personal_hotspot_enabled,omitempty"`
	SubscriberMCC            string `json:"subscriber_mcc,omitempty"`
	SubscriberMNC            string `json:"subscriber_mnc,omitempty"`
	CurrentMCC               string `json:"current_mcc,omitempty"`
	CurrentMNC               string `json:"current_mnc,omitempty"`
}

// OSUpdateSettingsResponse contains information about macOS update settings.
type OSUpdateSettingsResponse struct {
	AutoCheckEnabled                bool   `json:"auto_check_enabled"`
	AutomaticAppInstallationEnabled bool   `json:"automatic_app_installation_enabled"`
	AutomaticOSInstallationEnabled  bool   `json:"automatic_os_installation_enabled"`
	AutomaticSecurityUpdatesEnabled bool   `json:"automatic_security_updates_enabled"`
	BackgroundDownloadEnabled       bool   `json:"background_download_enabled"`
	CatalogURL                      string `json:"catalog_url"`
	IsDefaultCatalog                bool   `json:"is_default_catalog"`
	PerformPeriodicCheck            bool   `json:"perform_periodic_check"`
	PreviousScanDate                string `json:"previous_scan_date"`
	PreviousScanResult              int    `json:"previous_scan_result"`
}

// MacosQueryResponses contains macOS queryResponses
type MacosQueryResponses struct {
	OSUpdateSettings   OSUpdateSettingsResponse // OSX 10.11+
	LocalHostName      string                   `json:"local_host_name,omitempty"` // OSX 10.11
	HostName           string                   `json:"host_name,omitempty"`       // OSX 10.11
	ActiveManagedUsers []string                 `json:"active_managed_users"`      // OSX 10.11
}

// QueryResponses is a DeviceInformation MDM Command Response
type QueryResponses struct {
	CommonQueryResponses
	MacosQueryResponses
	IosQueryResponses
	AtvQueryResponses
}
