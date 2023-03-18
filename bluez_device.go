package bluez

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"
	"strings"
)

func (self *BlueZDevice) GetCharacteristic(uuid string, serviceUuid string) (foundCharacteristic *BlueZCharacteristic, foundService *BlueZService, e error) {
	uuid = strings.ToLower(uuid)
	serviceUuid = strings.ToLower(serviceUuid)
	if managedObjects, err := self.adapter.bluez.getManagedObjects(); err == nil {
		for path, o := range managedObjects {
			if hasPrefix(path, self.Object.Path()) {
				if data, exists := o["org.bluez.GattService1"]; exists {
					if strings.ToLower(data["UUID"].Value().(string)) == serviceUuid {
						foundService = &BlueZService{BlueZObject: BlueZObject{self.Conn, self.Conn.Object("org.bluez", path)}, device: self, data: data}
						for path, o := range managedObjects {
							if hasPrefix(path, self.Object.Path()) {
								if data, exists := o["org.bluez.GattCharacteristic1"]; exists {
									if strings.ToLower(data["UUID"].Value().(string)) == uuid {
										foundCharacteristic = &BlueZCharacteristic{BlueZObject: BlueZObject{self.Conn, self.Conn.Object("org.bluez", path)}, service: foundService, data: data}
										return
									}
								}
							}
						}
					}
				}
			}
		}
		if foundCharacteristic == nil {
			e = errors.Errorf("Characteristic %s not found", uuid)
		} else if foundService == nil {
			e = errors.Errorf("Service %s not found", serviceUuid)
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZDevice) GetServices() (services []BlueZService, e error) {
	if managedObjects, err := self.adapter.bluez.getManagedObjects(); err == nil {
		for path, o := range managedObjects {
			if hasPrefix(path, self.Object.Path()) {
				if data, exists := o["org.bluez.GattService1"]; exists {
					service := BlueZService{BlueZObject: BlueZObject{self.Conn, self.Conn.Object("org.bluez", path)}, device: self, data: data}
					services = append(services, service)
				}
			}
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZDevice) Connect() (e error) {
	if self.Object != nil {
		if call := self.Object.Call("org.bluez.Device1.Connect", 0); call.Err == nil {
			if err := self.Refresh(); err == nil {

			} else {
				e = err
			}
		} else {
			e = call.Err
		}
	} else {
		e = errors.New("Device not initialized")
	}
	e = errors.Wrap(e, "Connect with peripheral failed")
	return
}

func (self *BlueZDevice) Disconnect() (e error) {
	if self.Object != nil {
		if call := self.Object.Call("org.bluez.Device1.Disconnect", 0); call.Err == nil {
			if err := self.Refresh(); err == nil {

			} else {
				e = err
			}
		} else {
			e = call.Err
		}
	} else {
		e = errors.New("Device not initialized")
	}
	e = errors.Wrap(e, "Disconnect from peripheral failed")
	return
}

func (self *BlueZDevice) Name() (value string) {
	if v, exists := self.deviceObject["Name"]; exists {
		value = v.Value().(string)
	}
	return
}

func (self *BlueZDevice) Connected() (value bool) {
	if v, exists := self.deviceObject["Connected"]; exists {
		value = v.Value().(bool)
	}
	return
}

func (self *BlueZDevice) ServicesResolved() (value bool) {
	if v, exists := self.deviceObject["ServicesResolved"]; exists {
		value = v.Value().(bool)
	}
	return
}

func (self *BlueZDevice) Address() (value string) {
	if v, exists := self.deviceObject["Address"]; exists {
		value = v.Value().(string)
	}
	return
}

func (self *BlueZDevice) UUIDs() (value []string) {
	if v, exists := self.deviceObject["UUIDs"]; exists {
		value = v.Value().([]string)
	}
	return
}

func (self *BlueZDevice) Refresh() (e error) {
	if managedObject, err := self.adapter.bluez.getManagedObject(self.Object.Path()); err == nil && len(managedObject) != 0 {
		self.deviceObject = managedObject["org.bluez.Device1"]
	} else {
		e = err
	}
	return
}

func (self *BlueZDevice) ToDisplayString() (text string) {
	for k, v := range self.deviceObject {
		text += fmt.Sprintf("\t%v %v\n", k, v)
	}
	return
}

type BlueZDevice struct {
	BlueZObject
	adapter      *BlueZAdapter
	deviceObject map[string]dbus.Variant
}
