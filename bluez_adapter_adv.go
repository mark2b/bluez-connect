package bluez

import (
	"github.com/godbus/dbus"
	"github.com/pkg/errors"
)

func (self *BlueZAdapter) StartAdvertise(path string, localName string, serviceUUIDs []string) (e error) {
	advertisementObjectPath := dbus.ObjectPath(path)

	if obj := self.Conn.Object("org.bluez", self.Object.Path()); obj != nil {
		advertisement := &LEAdvertisement1{serviceUUIDs: serviceUUIDs, duration: 15, timeout: 5}
		if len(localName) == 0 {
			advertisement.includes = []string{"tx-power", "appearence", "local-name"}
		} else {
			advertisement.localName = localName
			advertisement.includes = []string{"tx-power", "appearence"}

		}
		if self.advertisements == nil {
			self.advertisements = make(map[dbus.ObjectPath]*LEAdvertisement1, 0)
		}
		self.advertisements[advertisementObjectPath] = advertisement

		if err := self.bluez.exportSingletonWithProperties(advertisement, advertisementObjectPath, LEAdvertisement1Interface, LEAdvertisement1Intro); err == nil {
			options := make(map[string]dbus.Variant)
			if call := obj.Call("org.bluez.LEAdvertisingManager1.RegisterAdvertisement", 0, advertisementObjectPath, options); call.Err == nil {

			} else {
				e = call.Err
			}
		} else {
			e = err
		}
	} else {
		e = errors.New("Can't create LEAdvertisingManager1 Object")
	}
	return
}

func (self *BlueZAdapter) StopAdvertise(path string) (e error) {
	advertisementObjectPath := dbus.ObjectPath(path)
	if obj := self.Conn.Object("org.bluez", self.Object.Path()); obj != nil {
		if call := obj.Call("org.bluez.LEAdvertisingManager1.UnregisterAdvertisement", 0, advertisementObjectPath); call.Err == nil {
			delete(self.advertisements, advertisementObjectPath)
		} else {
			e = call.Err
		}
	} else {
		e = errors.New("Can't create LEAdvertisingManager1 Object")
	}
	return
}

type LEAdvertisement1 struct {
	serviceUUIDs []string
	//ManufacturerData map[string]dbus.Variant
	//SolicitUUIDs     []string
	//ServiceData      map[string]dbus.Variant
	includes  []string
	localName string
	//Appearance       uint16
	duration uint16
	timeout  uint16
}

func (self *LEAdvertisement1) Type() (string, *dbus.Error) {
	return "peripheral", nil
}

func (self *LEAdvertisement1) ServiceUUIDs() ([]string, *dbus.Error) {
	return self.serviceUUIDs, nil
}

func (self *LEAdvertisement1) Includes() ([]string, *dbus.Error) {
	return self.includes, nil
}

func (self *LEAdvertisement1) LocalName() (string, *dbus.Error) {
	return self.localName, nil
}

func (self *LEAdvertisement1) Duration() (uint16, *dbus.Error) {
	return self.duration, nil
}

func (self *LEAdvertisement1) Timeout() (uint16, *dbus.Error) {
	return self.timeout, nil
}

func (self *LEAdvertisement1) GetAll(iface string) (properties map[string]dbus.Variant, e *dbus.Error) {
	names := []string{"Type", "ServiceUUIDs", "Includes", "Duration", "Timeout", "LocalName"}
	props := make(map[string]dbus.Variant, 0)
	for _, name := range names {
		if value, err := self.get(name); err == nil {
			props[name] = dbus.MakeVariant(value)
		} else {
			e = err
			break
		}
	}
	properties = props
	return
}

func (self *LEAdvertisement1) Get(iface string, name string) (variant dbus.Variant, e *dbus.Error) {
	if value, err := self.get(name); err == nil {
		variant = dbus.MakeVariant(value)
	} else {
		e = err
	}
	return
}

func (self *LEAdvertisement1) get(name string) (interface{}, *dbus.Error) {
	switch name {
	case "Type":
		return self.Type()
	case "ServiceUUIDs":
		return self.ServiceUUIDs()
	case "Includes":
		return self.Includes()
	case "LocalName":
		return self.LocalName()
	case "Duration":
		return self.Duration()
	case "Timeout":
		return self.Timeout()
	default:
		return nil, dbus.MakeFailedError(errors.Errorf("Property '%s' not found", name))
	}
}

func (self *LEAdvertisement1) Release() (dbusError *dbus.Error) {
	return
}
