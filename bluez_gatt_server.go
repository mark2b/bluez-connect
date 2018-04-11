package bluez

import (
	"fmt"
	"github.com/godbus/dbus"
)

func (self *BlueZGattManager) AddApplication(path string, gattApplication *GattApplication) (e error) {
	applicationPath := dbus.ObjectPath(path)
	if err := self.prepareBlueZGattApplication(applicationPath, gattApplication); err == nil {
		options := make(map[string]dbus.Variant)
		if call := self.Object.Call("org.bluez.GattManager1.RegisterApplication", 0, dbus.ObjectPath(applicationPath), options); call.Err == nil {

		} else {
			e = call.Err
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZGattManager) prepareBlueZGattApplication(applicationPath dbus.ObjectPath, gattApplication *GattApplication) (e error) {
	bluezGattApplication := &BlueZGattApplication{blueZGattServices: make([]*BlueZGattService, 0)}

	for _, service := range gattApplication.Services {
		bluezGattService := &BlueZGattService{gattService: service}
		bluezGattApplication.blueZGattServices = append(bluezGattApplication.blueZGattServices, bluezGattService)

		for _, characteristic := range service.Characteristics {
			bluezGattCharacteristic := &BlueZGattCharacteristic{blueZGattService: bluezGattService, gattCharacteristic: characteristic}
			bluezGattService.blueZGattCharacteristics = append(bluezGattService.blueZGattCharacteristics, bluezGattCharacteristic)
		}
	}

	managedObjects := make(map[dbus.ObjectPath]map[string]map[string]dbus.Variant, 0)

	if err := self.adapter.bluez.export(bluezGattApplication, applicationPath, ObjectManagerInterface, ObjectManagerIntro); err == nil {
		managedObjects[dbus.ObjectPath(applicationPath)] = map[string]map[string]dbus.Variant{ObjectManagerInterface: map[string]dbus.Variant{}}

		self.application = bluezGattApplication

		for index, bluezGattService := range bluezGattApplication.blueZGattServices {

			props, _ := bluezGattService.GetAll("")

			bluezGattService.path = dbus.ObjectPath(fmt.Sprintf("%s/service%d", applicationPath, index))
			if err := self.adapter.bluez.exportWithProperties(bluezGattService, bluezGattService.path, GattService1Interface, GattService1Intro); err == nil {
				managedObjects[dbus.ObjectPath(bluezGattService.path)] = map[string]map[string]dbus.Variant{
					GattService1Interface:       props,
					DBusIntrospectableInterface: map[string]dbus.Variant{},
					DBusPropertiesInterface:     map[string]dbus.Variant{},
				}

				for index, blueZGattCharacteristic := range bluezGattService.blueZGattCharacteristics {
					props, _ := blueZGattCharacteristic.GetAll("")

					blueZGattCharacteristic.path = dbus.ObjectPath(fmt.Sprintf("%s/char%d", bluezGattService.path, index))
					if err := self.adapter.bluez.exportWithProperties(blueZGattCharacteristic, blueZGattCharacteristic.path, GattCharacteristic1Interface, GattCharacteristic1Intro); err == nil {
						managedObjects[dbus.ObjectPath(blueZGattCharacteristic.path)] = map[string]map[string]dbus.Variant{
							GattCharacteristic1Interface: props,
							DBusIntrospectableInterface:  map[string]dbus.Variant{},
							DBusPropertiesInterface:      map[string]dbus.Variant{},
						}
					} else {
						e = err
						break
					}
				}
			} else {
				e = err
				break
			}
		}
	} else {
		e = err
	}
	bluezGattApplication.managedObjects = managedObjects
	return
}

type BlueZGattManager struct {
	BlueZObject
	adapter     *BlueZAdapter
	application *BlueZGattApplication
	data        map[string]dbus.Variant
}

var blueZGattManager *BlueZGattManager

func (self *BlueZGattApplication) GetManagedObjects() (managedObjects map[dbus.ObjectPath]map[string]map[string]dbus.Variant, e *dbus.Error) {
	return self.managedObjects, nil
}

type BlueZGattApplication struct {
	blueZGattServices []*BlueZGattService
	managedObjects    map[dbus.ObjectPath]map[string]map[string]dbus.Variant
}
