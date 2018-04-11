package bluez

import (
	"fmt"
	"github.com/godbus/dbus"
)

func (self *BlueZCharacteristic) ReadValue() (data []byte, e error) {
	options := make(map[string]dbus.Variant)
	if call := self.Object.Call("org.bluez.GattCharacteristic1.ReadValue", 0, options); call.Err == nil {
		if err := call.Store(&data); err == nil {
		} else {
			e = err
		}
	} else {
		e = call.Err
	}
	return
}

func (self *BlueZCharacteristic) WriteValue(value []byte) (e error) {
	options := make(map[string]dbus.Variant)
	if call := self.Object.Call("org.bluez.GattCharacteristic1.WriteValue", 0, value, options); call.Err == nil {
	} else {
		e = call.Err
	}
	return
}

func (self *BlueZCharacteristic) StartNotify() (e error) {
	if call := self.Object.Call("org.bluez.GattCharacteristic1.StartNotify", 0); call.Err == nil {
	} else {
		e = call.Err
	}
	return
}

func (self *BlueZCharacteristic) StopNotify() (e error) {
	if call := self.Object.Call("org.bluez.GattCharacteristic1.StopNotify", 0); call.Err == nil {
	} else {
		e = call.Err
	}
	return
}

func (self *BlueZCharacteristic) UUID() (value string) {
	if v, exists := self.data["UUID"]; exists {
		value = v.Value().(string)
	}
	return
}

func (self *BlueZCharacteristic) ToDisplayString() (text string) {
	for k, v := range self.data {
		text += fmt.Sprintf("\t%v %v\n", k, v)
	}
	return
}

type BlueZCharacteristic struct {
	BlueZObject
	service *BlueZService
	data    map[string]dbus.Variant
}
