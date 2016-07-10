package mdm

import (
	"encoding/json"
	"fmt"
	"github.com/groob/plist"
	"testing"
)

// Basic tests will attempt to marshal and unmarshal mdm command structures to identify any naming or tag errors.

// Make sure a command can be marshalled to json
func testMarshalJSON(t *testing.T, cmd interface{}) {
	jsonCmd, err := json.Marshal(cmd)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonCmd))
}

// Make sure a command can be marshalled to plist
func testMarshalPlist(t *testing.T, cmd interface{}) {
	plistCmd, err := plist.MarshalIndent(cmd, "\t")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(plistCmd))
}

func TestInstallProfile(t *testing.T) {
	cmd := InstallProfile{Payload: []byte{00}}
	testMarshalJSON(t, cmd)
	testMarshalPlist(t, cmd)
}

func TestRemoveProfile(t *testing.T) {
	cmd := RemoveProfile{Identifier: "io.micromdm.test.profile"}
	testMarshalJSON(t, cmd)
	testMarshalPlist(t, cmd)
}

func TestInstallProvisioningProfile(t *testing.T) {
	cmd := InstallProvisioningProfile{ProvisioningProfile: []byte{00}}
	testMarshalJSON(t, cmd)
	testMarshalPlist(t, cmd)
}

func TestRemoveProvisioningProfile(t *testing.T) {
	cmd := RemoveProvisioningProfile{UUID: "1111-2222-3333"}
	testMarshalJSON(t, cmd)
	testMarshalPlist(t, cmd)
}

func TestInstalledApplicationList(t *testing.T) {
	cmd := InstalledApplicationList{Identifiers: []string{"io.micromdm.application"}, ManagedAppsOnly: true}
	testMarshalJSON(t, cmd)
	testMarshalPlist(t, cmd)
}

func TestDeviceInformation(t *testing.T) {
	cmd := DeviceInformation{Queries: []string{"SerialNumber"}}
	testMarshalJSON(t, cmd)
	testMarshalPlist(t, cmd)
}

func TestDeviceLock(t *testing.T) {
	cmd := DeviceLock{PIN: "123456", Message: "Locked", PhoneNumber: "123-4567"}
	testMarshalJSON(t, cmd)
	testMarshalPlist(t, cmd)
}

func TestClearPasscode(t *testing.T) {
	cmd := ClearPasscode{UnlockToken: "abcdefg"}
	testMarshalJSON(t, cmd)
	testMarshalPlist(t, cmd)
}
