package bluez

import (
	"agtinternational.com/bluez-connect/log"
	"github.com/godbus/dbus"
	"github.com/pkg/errors"
)

type BlueZGattCharacteristic struct {
	blueZGattService   *BlueZGattService
	gattCharacteristic *GattCharacteristic
	path               dbus.ObjectPath
}

func (self *BlueZGattCharacteristic) ReadValue(options map[string]interface{}) (value []byte, e *dbus.Error) {
	log.Log.Debug("characteristic.ReadValue for %s %v", self.gattCharacteristic.UUID, options)
	if self.gattCharacteristic.OnReadFunc != nil {
		if v, err := self.gattCharacteristic.OnReadFunc(); err == nil {
			value = v
		} else {
			e = MakeFailedError(err)
		}
	} else {
		log.Log.Error("OnReadFunc is nil")
	}
	return
}

func (self *BlueZGattCharacteristic) WriteValue(value []byte, options map[string]interface{}) (e *dbus.Error) {
	log.Log.Debug("characteristic.WriteValue for %s %v", self.gattCharacteristic.UUID, options)
	if self.gattCharacteristic.OnWriteFunc != nil {
		if err := self.gattCharacteristic.OnWriteFunc(value); err == nil {
			log.Log.Debug("characteristic wrote value for %s", self.gattCharacteristic.UUID)

		} else {
			e = MakeFailedError(err)
		}
	} else {
		log.Log.Error("OnWriteFunc is nil")
	}
	return
}

func (self *BlueZGattCharacteristic) StartNotify() (e *dbus.Error) {
	log.Log.Debug("characteristic.StartNotify for %s", self.gattCharacteristic.UUID)
	return
}
func (self *BlueZGattCharacteristic) StopNotify() (e *dbus.Error) {
	log.Log.Debug("characteristic.StopNotify for %s", self.gattCharacteristic.UUID)
	return
}

func (self *BlueZGattCharacteristic) Confirm() (e *dbus.Error) {
	log.Log.Debug("characteristic.Confirm for %s", self.gattCharacteristic.UUID)
	return
}

func (self *BlueZGattCharacteristic) GetAll(iface string) (properties map[string]dbus.Variant, e *dbus.Error) {
	names := []string{"UUID", "Service", "Flags"}
	props := make(map[string]dbus.Variant, 0)
	for _, name := range names {
		if value, err := self.get(name); err == nil {
			props[name] = dbus.MakeVariant(value)
		} else {
			e = dbus.MakeFailedError(err)
			break
		}
	}
	properties = props
	return
}

func (self *BlueZGattCharacteristic) Get(iface string, name string) (variant dbus.Variant, e *dbus.Error) {
	if value, err := self.get(name); err == nil {
		variant = dbus.MakeVariant(value)
	} else {
		e = dbus.MakeFailedError(err)
	}
	return
}

func (self *BlueZGattCharacteristic) get(name string) (value interface{}, e error) {
	switch name {
	case "UUID":
		value = self.gattCharacteristic.UUID
	case "Service":
		value = self.blueZGattService.path
	case "Flags":
		value = self.gattCharacteristic.Flags
	default:
		e = errors.Errorf("Property '%s' not found", name)
	}
	return
}
