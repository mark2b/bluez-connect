package bluez

import (
	"github.com/godbus/dbus"
	"github.com/pkg/errors"
)

type BlueZGattService struct {
	gattService              *GattService
	blueZGattCharacteristics []*BlueZGattCharacteristic
	path                     dbus.ObjectPath
}

func (self *BlueZGattService) GetAll(iface string) (properties map[string]dbus.Variant, e *dbus.Error) {
	names := []string{"UUID", "Primary"}
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

func (self *BlueZGattService) Get(iface string, name string) (variant dbus.Variant, e *dbus.Error) {
	if value, err := self.get(name); err == nil {
		variant = dbus.MakeVariant(value)
	} else {
		e = dbus.MakeFailedError(err)
	}
	return
}

func (self *BlueZGattService) get(name string) (value interface{}, e error) {
	switch name {
	case "UUID":
		value = self.gattService.UUID
	case "Primary":
		value = true
	default:
		e = errors.Errorf("Property '%s' not found", name)
	}
	return
}
