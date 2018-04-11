package main

import (
	dbus "github.com/godbus/dbus"
	"bluez-connect"
	"log"
)

const (
	GattHeartRateServiceUUID                   = "0000180D-0000-1000-8000-00805F9B34FB" // 180D Heart Rate service
	GattHeartRateMeasurementCharacteristicUUID = "00002A37-0000-1000-8000-00805F9B34FB" // 2A37 Heart Rate Measurement
)

func main() {

	peripheralName := "POLAR-NAME" // name or address

	if blueZ, err := bluez.NewBLueZ(); err == nil {
		blueZ.WaitForSignals(onBlueZSignal)

		if blueZAdapter, err := blueZ.GetAdapter("hci0"); err == nil {
			if err := blueZAdapter.StartDiscovery(); err == nil {

				if device := blueZAdapter.FindPeripheral(peripheralName); device != nil {
					device.AddPropertiesObserver()
					log.Println("Connecting to", peripheralName)
					if device.Connected() && device.ServicesResolved() {
						log.Println("Already connected")
					}
					if err := device.Connect(); err == nil {
						log.Println("Connected")

						// Subscribe
						if c, _, err := device.GetCharacteristic(GattHeartRateMeasurementCharacteristicUUID, GattHeartRateServiceUUID); err == nil {
							c.AddPropertiesObserver()
							if err := c.StartNotify(); err == nil {

							} else {
								log.Println("StartNotify failed", err.Error())
							}
						} else {
							log.Println("Characteristic not found")
						}

						// Read
						if c, _, err := device.GetCharacteristic(GattHeartRateMeasurementCharacteristicUUID, GattHeartRateServiceUUID); err == nil {
							if data, err := c.ReadValue(); err == nil {
								log.Println("Received data from characteristic", data)
							} else {
								log.Println("ReadValue failed", err.Error())
							}
						} else {
							log.Println("Characteristic not found")
						}

						// Write
						if c, _, err := device.GetCharacteristic(GattHeartRateMeasurementCharacteristicUUID, GattHeartRateServiceUUID); err == nil {
							if err := c.WriteValue([]byte("Hello, World !!!")); err == nil {
								log.Println("Wrote value to characteristic")
							} else {
								log.Println("ReadValue failed", err.Error())
							}
						} else {
							log.Println("Characteristic not found")
						}

					} else {
						log.Println("Connection failed", err.Error())
					}
				} else {
					log.Println("Peripheral not found")
				}
			}
		}
	}
}

func onPeripheralPropertyChanged(signal *dbus.Signal) {
	log.Println("onPeripheralPropertyChanged")
	body := signal.Body
	if len(body) >= 2 {
		if body[0] == "org.bluez.GattCharacteristic1" {
			if properties, ok := body[1].(map[string]dbus.Variant); ok {
				if variant, exists := properties["Value"]; exists {
					log.Println("Received data from characteristic", signal.Path, variant.Value().([]byte))
				}
			}
		}
	}
}

func onBlueZSignal(signal *dbus.Signal) {
	switch signal.Name {
	case "org.freedesktop.DBus.Properties.PropertiesChanged":
		onPeripheralPropertyChanged(signal)
	case "org.freedesktop.DBus.ObjectManager.InterfacesAdded":
	case "org.freedesktop.DBus.ObjectManager.InterfacesRemoved":
	}
}
