package agent

import (
	"github.com/godbus/dbus"
	"github.com/mark2b/bluez-connect"
)

func (self *DefaultAgent) Release() (e *dbus.Error) {
	println("Release")
	return
}

func (self *DefaultAgent) RequestPinCode(device dbus.Object) (pincode string, e *dbus.Error) {
	println("RequestPinCode")
	return
}

func (self *DefaultAgent) DisplayPinCode(device dbus.Object, pincode string) (e *dbus.Error) {
	println("DisplayPinCode")
	return
}

func (self *DefaultAgent) RequestPasskey(device dbus.Object) (passkey uint32, e *dbus.Error) {
	println("RequestPasskey")
	return
}

func (self *DefaultAgent) DisplayPasskey(device dbus.Object, passkey uint32, entered int16) (e *dbus.Error) {
	println("DisplayPasskey")
	return
}

func (self *DefaultAgent) RequestConfirmation(device dbus.Object, passkey uint32) (e *dbus.Error) {
	println("RequestConfirmation")
	return
}

func (self *DefaultAgent) RequestAuthorization(device dbus.Object) (e *dbus.Error) {
	println("RequestAuthorization")
	return
}

func (self *DefaultAgent) AuthorizeService(device dbus.Object, uuid string) (e *dbus.Error) {
	println("AuthorizeService")
	return
}

func (self *DefaultAgent) Cancel() (e *dbus.Error) {
	println("Cancel")
	return
}

func (self *DefaultAgent) Capability() (bluez.AgentCapability) {
	println("Cancel")
	return bluez.NoInputNoOutput
}

func NewDefaultAgent() *DefaultAgent {
	return &DefaultAgent{}
}

type DefaultAgent struct {
}
