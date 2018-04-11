package bluez

import (
	"fmt"
	"github.com/godbus/dbus"
)

func (self *BlueZService) GetCharacteristicByUUID(uuid string) (foundCharacteristic *BlueZCharacteristic, e error) {
	if characteristics, err := self.GetCharacteristics(); err == nil {
		for _, characteristic := range characteristics {
			if characteristic.UUID() == uuid {
				foundCharacteristic = &characteristic
				break
			}
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZService) GetCharacteristics() (characteristics []BlueZCharacteristic, e error) {
	if managedObjects, err := self.device.adapter.bluez.getManagedObjects(); err == nil {
		for path, o := range managedObjects {
			if HasPrefix(path, self.Object.Path()) {
				if data, exists := o["org.bluez.GattCharacteristic1"]; exists {
					characteristic := BlueZCharacteristic{BlueZObject: BlueZObject{self.Conn, self.Conn.Object("org.bluez", path)}, service: self, data: data}
					characteristics = append(characteristics, characteristic)
				}
			}
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZService) UUID() (value string) {
	if v, exists := self.data["UUID"]; exists {
		value = v.Value().(string)
	}
	return
}

func (self *BlueZService) ToDisplayString() (text string) {
	text += fmt.Sprintf("%v\n", self.Object.Path())
	for k, v := range self.data {
		text += fmt.Sprintf("\t%v %v\n", k, v)
	}
	return
}

type BlueZService struct {
	BlueZObject
	device *BlueZDevice
	data   map[string]dbus.Variant
}
