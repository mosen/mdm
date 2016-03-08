package mdm

// Response is an MDM Command Response
type Response struct {
	UDID           string
	Status         string
	CommandUUID    string
	RequestType    string         `json:"request_type,omitempty" plist:",omitempty"`
	QueryResponses QueryResponses `json:"query_responses,omitempty" plist:",omitempty"`
}

// QueryResponses is a DeviceInformation MDM Command Response
type QueryResponses struct {
	Model       string
	ProductName string
}
