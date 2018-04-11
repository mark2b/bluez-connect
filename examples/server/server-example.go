package main

import (
	"github.com/godbus/dbus"
	bluez "bluez-connect"
	"bluez-connect/examples/server/service"
	"log"
)

func main() {
	var echoService = service.NewService()
	var gattApplication = bluez.NewGattApplication("/example/server")
	gattApplication.AddService(echoService)

	if blueZ, err := bluez.NewBLueZ(); err == nil {
		if blueZAdapter, err := blueZ.GetAdapter("hci0"); err == nil {
			blueZ.WaitForSignals(onBlueZSignal)

			if err := blueZAdapter.StartAdvertise(gattApplication.Path, "ECHO", []string{echoService.UUID}); err == nil {
				if blueZGattManager, err := blueZAdapter.GetGattManager(); err == nil {
					if err := blueZGattManager.AddApplication(gattApplication); err == nil {
						log.Println("Bluetooth server started")
						select {}
					}
				} else {
					println("%s", err.Error())
				}
			} else {
				println("%s", err.Error())
			}
		} else {
			println("%s", err.Error())
		}
	} else {
		println("%s", err.Error())
	}
}

func onBlueZSignal(signal *dbus.Signal) {
	switch signal.Name {
	case "org.freedesktop.DBus.Properties.PropertiesChanged":
	case "org.freedesktop.DBus.ObjectManager.InterfacesAdded":
	case "org.freedesktop.DBus.ObjectManager.InterfacesRemoved":
	}
}
