package bluez

import (
	"github.com/godbus/dbus/v5"
)

type AgentCapability string

const (
	DisplayOnly     AgentCapability = "DisplayOnly"
	DisplayYesNo    AgentCapability = "DisplayYesNo"
	KeyboardOnly    AgentCapability = "KeyboardOnly"
	NoInputNoOutput AgentCapability = "NoInputNoOutput"
	KeyboardDisplay AgentCapability = "KeyboardDisplay"
)

type Agent interface {
	Release() (e *dbus.Error)
	RequestPinCode(device dbus.Object) (pincode string, e *dbus.Error)
	DisplayPinCode(device dbus.Object, pincode string) (e *dbus.Error)
	RequestPasskey(device dbus.Object) (passkey uint32, e *dbus.Error)
	DisplayPasskey(device dbus.Object, passkey uint32, entered int16) (e *dbus.Error)
	RequestConfirmation(device dbus.Object, passkey uint32) (e *dbus.Error)
	RequestAuthorization(device dbus.Object) (e *dbus.Error)
	AuthorizeService(device dbus.Object, uuid string) (e *dbus.Error)
	Cancel() (e *dbus.Error)
	Capability() (capability AgentCapability)
}
