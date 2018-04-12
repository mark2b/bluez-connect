package bluez

import (
	"github.com/godbus/dbus"
	"github.com/pkg/errors"
)

type BlueZGattCharacteristic struct {
	blueZGattService   *BlueZGattService
	gattCharacteristic *GattCharacteristic
	path               dbus.ObjectPath
}

func (self *BlueZGattCharacteristic) ReadValue(options map[string]interface{}) (value []byte, e *dbus.Error) {
	if self.gattCharacteristic.OnReadFunc != nil {
		if v, err := self.gattCharacteristic.OnReadFunc(); err == nil {
			value = v
		} else {
			e = MakeFailedError(err)
		}
	}
	return
}

func (self *BlueZGattCharacteristic) WriteValue(value []byte, options map[string]interface{}) (e *dbus.Error) {
	if self.gattCharacteristic.OnWriteFunc != nil {
		if err := self.gattCharacteristic.OnWriteFunc(value); err == nil {

		} else {
			e = MakeFailedError(err)
		}
	}
	return
}

func (self *BlueZGattCharacteristic) StartNotify() (e *dbus.Error) {
	return
}
func (self *BlueZGattCharacteristic) StopNotify() (e *dbus.Error) {
	return
}

func (self *BlueZGattCharacteristic) Confirm() (e *dbus.Error) {
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
