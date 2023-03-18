package agent

import (
	"github.com/godbus/dbus/v5"
	"github.com/mark2b/bluez-connect/v2"
	"github.com/mark2b/bluez-connect/v2/internal/log"
)

func (self *DefaultAgent) Release() (e *dbus.Error) {
	log.Log.Debug("Release")
	return
}

func (self *DefaultAgent) RequestPinCode(device dbus.Object) (pincode string, e *dbus.Error) {
	log.Log.Debug("RequestPinCode")
	return
}

func (self *DefaultAgent) DisplayPinCode(device dbus.Object, pincode string) (e *dbus.Error) {
	log.Log.Debug("DisplayPinCode")
	return
}

func (self *DefaultAgent) RequestPasskey(device dbus.Object) (passkey uint32, e *dbus.Error) {
	log.Log.Debug("RequestPasskey")
	return
}

func (self *DefaultAgent) DisplayPasskey(device dbus.Object, passkey uint32, entered int16) (e *dbus.Error) {
	log.Log.Debug("DisplayPasskey")
	return
}

func (self *DefaultAgent) RequestConfirmation(device dbus.Object, passkey uint32) (e *dbus.Error) {
	log.Log.Debug("RequestConfirmation")
	return
}

func (self *DefaultAgent) RequestAuthorization(device dbus.Object) (e *dbus.Error) {
	log.Log.Debug("RequestAuthorization")
	return
}

func (self *DefaultAgent) AuthorizeService(device dbus.Object, uuid string) (e *dbus.Error) {
	log.Log.Debug("AuthorizeService")
	return
}

func (self *DefaultAgent) Cancel() (e *dbus.Error) {
	log.Log.Debug("Cancel")
	return
}

func (self *DefaultAgent) Capability() bluez.AgentCapability {
	log.Log.Debug("Cancel")
	return bluez.NoInputNoOutput
}

func NewDefaultAgent() *DefaultAgent {
	return &DefaultAgent{}
}

type DefaultAgent struct {
}
