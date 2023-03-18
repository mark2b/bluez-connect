package bluez

import (
	"github.com/godbus/dbus/v5"
	"github.com/mark2b/bluez-connect/v2/internal/log"
	"github.com/pkg/errors"
)

type BlueZGattCharacteristic struct {
	BlueZObject
	blueZGattService   *BlueZGattService
	gattCharacteristic *GattCharacteristic
	path               dbus.ObjectPath
	value              []byte
	notifying          bool
}

func (self *BlueZGattCharacteristic) ReadValue(options map[string]interface{}) (output []byte, e *dbus.Error) {
	log.Log.Debugf("ReadValue %v", options)
	if self.gattCharacteristic.OnReadFunc != nil {
		if v, err := self.gattCharacteristic.OnReadFunc(); err == nil {
			output = v
			if err := self.set(GattCharacteristic1Interface, "Value", dbus.MakeVariant(output), true); err == nil {
			} else {
				log.Log.Errorf("SetProperty error: %v", err)
				e = MakeFailedError(err)
			}
		} else {
			e = MakeFailedError(err)
		}
	} else if self.gattCharacteristic.OnReadValueFunc != nil {
		if v, err := self.get(GattCharacteristic1PropValue); err == nil {
			if variant, ok := v.(dbus.Variant); ok {
				if value, ok := variant.Value().([]byte); ok {
					if v, err := self.gattCharacteristic.OnReadValueFunc(value); err == nil {
						output = v
					} else {
						e = MakeFailedError(err)
					}
				}
			} else if value, ok := v.([]byte); ok {
				if v, err := self.gattCharacteristic.OnReadValueFunc(value); err == nil {
					output = v
				} else {
					e = MakeFailedError(err)
				}
			}
		} else {
			log.Log.Errorf("GetProperty error: %v", err)
			e = MakeFailedError(err)
		}
	}
	return
}

func (self *BlueZGattCharacteristic) WriteValue(input []byte, options map[string]interface{}) (e *dbus.Error) {
	log.Log.Debugf("WriteValue %v", options)
	if self.gattCharacteristic.OnWriteFunc != nil {
		if err := self.set(GattCharacteristic1Interface, "Value", dbus.MakeVariant([]byte{}), false); err == nil {
			if v, err := self.gattCharacteristic.OnWriteFunc(input); err == nil {
				if err := self.set(GattCharacteristic1Interface, "Value", dbus.MakeVariant(v), true); err == nil {
				} else {
					log.Log.Errorf("SetProperty error: %v", err)
					e = MakeFailedError(err)
				}
			} else {
				e = MakeFailedError(err)
			}
		}
	} else if self.gattCharacteristic.OnWriteAsyncFunc != nil {
		go func() {
			if err := self.set(GattCharacteristic1Interface, "Value", dbus.MakeVariant([]byte{}), false); err == nil {
				if v, err := self.gattCharacteristic.OnWriteAsyncFunc(input); err == nil {
					if err := self.set(GattCharacteristic1Interface, "Value", dbus.MakeVariant(v), true); err == nil {
					} else {
						log.Log.Errorf("SetProperty error: %v", err)
						e = MakeFailedError(err)
					}
				} else {
					log.Log.Errorf("OnWriteAsyncFunc error: %v", err)
					e = MakeFailedError(err)
				}
			}
		}()
	}
	return
}

func (self *BlueZGattCharacteristic) StartNotify() (e *dbus.Error) {
	log.Log.Debugf("StartNotify")
	if err := self.Object.SetProperty("org.bluez.GattCharacteristic1.Notifying", dbus.MakeVariant(true)); err == nil {
	} else {
		e = MakeFailedError(err)
	}
	return
}
func (self *BlueZGattCharacteristic) StopNotify() (e *dbus.Error) {
	log.Log.Debugf("StopNotify")
	if err := self.Object.SetProperty("org.bluez.GattCharacteristic1.Notifying", dbus.MakeVariant(false)); err == nil {
	} else {
		e = MakeFailedError(err)
	}
	return
}

func (self *BlueZGattCharacteristic) Confirm() (e *dbus.Error) {
	log.Log.Debugf("Confirm")
	return
}

func (self *BlueZGattCharacteristic) GetAll(iface string) (properties map[string]dbus.Variant, e *dbus.Error) {
	names := []string{GattCharacteristic1PropUUID, GattCharacteristic1PropService, GattCharacteristic1PropValue, GattCharacteristic1PropNotifying, GattCharacteristic1PropFlags}
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
	log.Log.Debugf("Get %s", name)
	switch name {
	case GattCharacteristic1PropUUID:
		value = self.gattCharacteristic.UUID
	case GattCharacteristic1PropService:
		value = self.blueZGattService.path
	case GattCharacteristic1PropValue:
		value = self.value
	case GattCharacteristic1PropNotifying:
		value = self.notifying
	case GattCharacteristic1PropFlags:
		value = self.gattCharacteristic.Flags
	default:
		e = errors.Errorf("Property(Get) '%s' not found", name)
	}
	return
}
func (self *BlueZGattCharacteristic) set(iface string, name string, variant dbus.Variant, emitPropertyChanged bool) (e *dbus.Error) {
	if err := self.Set(iface, name, variant); err == nil {
		if emitPropertyChanged {
			log.Log.Debugf("Emit %s of %s %s ", "org.freedesktop.DBus.Properties.PropertiesChanged", iface, name)
			if err := self.Conn.Emit(self.path, "org.freedesktop.DBus.Properties.PropertiesChanged", iface, map[string]dbus.Variant{name: variant}, []string{}); err == nil {

			} else {
				log.Log.Errorf("Emit error: %v", err)
				e = dbus.MakeFailedError(err)
			}
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZGattCharacteristic) Set(iface string, name string, variant dbus.Variant) (e *dbus.Error) {
	log.Log.Debugf("Set %s = %v", name, variant.Value())
	switch name {
	case GattCharacteristic1PropValue:
		self.value = variant.Value().([]byte)
	case GattCharacteristic1PropNotifying:
		self.notifying = variant.Value().(bool)
	default:
		e = dbus.MakeFailedError(errors.Errorf("Property(Set) '%s' not found", name))
	}
	return
}
