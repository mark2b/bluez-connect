package bluez

import (
	"github.com/godbus/dbus"
	"github.com/pkg/errors"
	"strings"
)

func (self *BlueZAdapter) GetDevices() (devices []*BlueZDevice, e error) {
	if managedObjects, err := self.bluez.getManagedObjects(); err == nil {
		for path, o := range managedObjects {
			if HasPrefix(path, self.Object.Path()) {
				if data, exists := o["org.bluez.Device1"]; exists {
					device := &BlueZDevice{BlueZObject: BlueZObject{self.Conn, self.Conn.Object("org.bluez", path)}, adapter: self, data: data}
					devices = append(devices, device)
				}
			}
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZAdapter) GetGattManager() (gattManager *BlueZGattManager, e error) {
	if managedObjects, err := self.bluez.getManagedObjects(); err == nil {
		for path, o := range managedObjects {
			if HasPrefix(path, self.Object.Path()) {
				if data, exists := o["org.bluez.GattManager1"]; exists {
					blueZGattManager = &BlueZGattManager{BlueZObject: BlueZObject{self.Conn, self.Conn.Object("org.bluez", path)}, adapter: self, data: data}
					gattManager = blueZGattManager
					return
				}
			}
		}
		if gattManager == nil {
			e = errors.New("GattManager not found")
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZAdapter) StartDiscovery() (e error) {
	if call := self.Object.Call("org.bluez.Adapter1.StartDiscovery", 0); call.Err == nil {

	} else {
		e = call.Err
	}
	return
}

func (self *BlueZAdapter) StopDiscovery() (e error) {
	if call := self.Object.Call("org.bluez.Adapter1.StopDiscovery", 0); call.Err == nil {

	} else {
		e = call.Err
	}
	return
}

func (self *BlueZAdapter) FindPeripheral(nameOrAddress string) (foundDevice *BlueZDevice) {
	tokens := strings.Split(nameOrAddress, ":")
	var findByAddress = len(tokens) == 6
	nameOrAddress = strings.ToLower(nameOrAddress)
	if devices, err := self.GetDevices(); err == nil {
		for _, device := range devices {
			var key string
			if findByAddress {
				key = device.Address()
			} else {
				key = device.Name()
			}
			if strings.ToLower(key) == nameOrAddress {
				foundDevice = device
				return
			}
		}
	}
	return
}

type BlueZAdapter struct {
	BlueZObject
	bluez          *BlueZ
	data           map[string]dbus.Variant
	advertisements map[dbus.ObjectPath]*LEAdvertisement1
}
