package main

import (
	"github.com/godbus/dbus/v5"
	bluez "github.com/mark2b/bluez-connect/v2"
	"github.com/mark2b/bluez-connect/v2/examples/server/agent"
	"github.com/mark2b/bluez-connect/v2/examples/server/service"
)

func main() {
	var echoService = service.NewService()
	var gattApplication = bluez.NewGattApplication("example.server", "/example/server")
	gattApplication.AddService(echoService)

	if blueZ, err := bluez.NewBLueZ(); err == nil {
		if blueZAdapter, err := blueZ.GetAdapter("hci0"); err == nil {
			blueZ.WaitForSignals(onBlueZSignal)

			if err := blueZAdapter.StartAdvertise("example.server", gattApplication.Path, "ECHO", []string{echoService.UUID}); err == nil {
				if blueZGattManager, err := blueZAdapter.GetGattManager(); err == nil {
					if err := blueZGattManager.AddApplication(gattApplication); err == nil {
						if err := blueZ.RegisterAgent(agent.NewDefaultAgent(), "example.server", "/example/server/agent", "com.white.connect"); err == nil {
							println("Bluetooth server started")
							select {}
						} else {
							println("RegisterAgent failed", err.Error())
						}
					} else {
						println("AddApplication failed", err.Error())
					}
				} else {
					println("GetGattManager failed", err.Error())
				}
			} else {
				println("StartAdvertise failed", err.Error())
			}
		} else {
			println("GetAdapter failed", err.Error())
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
