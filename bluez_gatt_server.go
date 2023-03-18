package bluez

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"strings"
)

func (self *BlueZGattManager) AddApplication(gattApplication *GattApplication) (e error) {
	applicationPath := dbus.ObjectPath(gattApplication.Path)
	if err := self.prepareBlueZGattApplication(gattApplication.Name, applicationPath, gattApplication); err == nil {
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

func (self *BlueZGattManager) prepareBlueZGattApplication(applicationName string, applicationPath dbus.ObjectPath, gattApplication *GattApplication) (e error) {
	bluezGattApplication := &BlueZGattApplication{blueZGattServices: make([]*BlueZGattService, 0)}

	for _, service := range gattApplication.Services {
		parentPath := normalizedObjectPath(applicationPath)

		bluezGattService := &BlueZGattService{gattService: service}
		bluezGattService.path = dbus.ObjectPath(fmt.Sprintf("%s/service%02d", parentPath, len(bluezGattApplication.blueZGattServices)+1))

		bluezGattApplication.blueZGattServices = append(bluezGattApplication.blueZGattServices, bluezGattService)

		for _, characteristic := range service.Characteristics {
			bluezGattCharacteristic := &BlueZGattCharacteristic{blueZGattService: bluezGattService, gattCharacteristic: characteristic}
			bluezGattCharacteristic.path = dbus.ObjectPath(fmt.Sprintf("%s/char%03d", bluezGattService.path, len(bluezGattService.blueZGattCharacteristics)+1))
			bluezGattService.blueZGattCharacteristics = append(bluezGattService.blueZGattCharacteristics, bluezGattCharacteristic)
		}
	}

	managedObjects := make(map[dbus.ObjectPath]map[string]map[string]dbus.Variant, 0)

	if err := self.adapter.bluez.export(bluezGattApplication, applicationName, applicationPath, ObjectManagerInterface, ObjectManagerIntro); err == nil {
		managedObjects[dbus.ObjectPath(applicationPath)] = map[string]map[string]dbus.Variant{ObjectManagerInterface: {}}

		self.application = bluezGattApplication

		for _, bluezGattService := range bluezGattApplication.blueZGattServices {

			props, _ := bluezGattService.GetAll("")
			if err := self.adapter.bluez.exportWithProperties(bluezGattService, bluezGattService.path, GattService1Interface, GattService1Intro); err == nil {
				managedObjects[dbus.ObjectPath(bluezGattService.path)] = map[string]map[string]dbus.Variant{
					GattService1Interface:       props,
					DBusIntrospectableInterface: {},
					DBusPropertiesInterface:     {},
				}

				for _, blueZGattCharacteristic := range bluezGattService.blueZGattCharacteristics {
					props, _ := blueZGattCharacteristic.GetAll("")

					blueZGattCharacteristic.BlueZObject = BlueZObject{Conn: self.adapter.Conn, Object: self.adapter.Conn.Object("com.white.connect", blueZGattCharacteristic.path)}
					if err := self.adapter.bluez.exportWithProperties(blueZGattCharacteristic, blueZGattCharacteristic.path, GattCharacteristic1Interface, GattCharacteristic1Intro); err == nil {
						managedObjects[dbus.ObjectPath(blueZGattCharacteristic.path)] = map[string]map[string]dbus.Variant{
							GattCharacteristic1Interface: props,
							DBusIntrospectableInterface:  {},
							DBusPropertiesInterface:      {},
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
	adapter           *BlueZAdapter
	application       *BlueZGattApplication
	gattManagerObject map[string]dbus.Variant
}

var blueZGattManager *BlueZGattManager

func (self *BlueZGattApplication) GetManagedObjects() (managedObjects map[dbus.ObjectPath]map[string]map[string]dbus.Variant, e *dbus.Error) {
	return self.managedObjects, nil
}

type BlueZGattApplication struct {
	blueZGattServices []*BlueZGattService
	managedObjects    map[dbus.ObjectPath]map[string]map[string]dbus.Variant
}

func normalizedObjectPath(path dbus.ObjectPath) dbus.ObjectPath {
	if strings.HasSuffix(string(path), "/") {
		return path[0 : len(path)-1]
	}
	return path
}
