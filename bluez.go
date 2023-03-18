package bluez

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/mark2b/bluez-connect/v2/internal/log"
	"github.com/pkg/errors"
	"strings"
)

func NewBLueZ() (bluez *BlueZ, e error) {
	if conn, err := dbus.SystemBus(); err == nil {
		if obj := conn.Object("org.bluez", "/"); obj != nil {
			bluez = &BlueZ{BlueZObject: &BlueZObject{conn, obj}}
		} else {
			conn.Close()
			err = errors.New("Can't create BlueZ Object")
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZ) Close() (e error) {
	e = self.Conn.Close()
	return
}

func (self *BlueZ) GetAdapter(hostId string) (a *BlueZAdapter, e error) {
	if adapters, err := self.GetAdapters(); err == nil {
		for _, adapter := range adapters {
			if strings.HasSuffix(string(adapter.Object.Path()), hostId) {
				a = adapter
				return
			}
		}
		if e == nil {
			e = errors.Errorf("Adapter %s not found", hostId)
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZ) GetAdapters() (adapters []*BlueZAdapter, e error) {
	if managedObjects, err := self.getManagedObjects(); err == nil {
		for path, o := range managedObjects {
			if adapterObject, exists := o["org.bluez.Adapter1"]; exists {
				adapter := &BlueZAdapter{BlueZObject: BlueZObject{self.Conn, self.Conn.Object("org.bluez", path)}, bluez: self, adapterObject: adapterObject}
				adapters = append(adapters, adapter)
			}
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZ) RegisterAgent(agent Agent, name string, path string, iface string) (e error) {
	agentPath := dbus.ObjectPath(path)
	if err := self.export(agent, name, agentPath, iface, Agent1Intro); err == nil {
		agentManager := self.Conn.Object("org.bluez", dbus.ObjectPath("/org/bluez"))
		if err := agentManager.Call("org.bluez.AgentManager1.RegisterAgent", 0, agentPath, agent.Capability()).Err; err == nil {

		} else {
			e = errors.WithMessage(err, "RegisterAgent failed.")
		}
	} else {
		e = errors.WithMessage(err, "export failed.")
	}
	return
}

func (self *BlueZ) UnregisterAgent(path string) (e error) {
	agentPath := dbus.ObjectPath(path)
	agentManager := self.Conn.Object("org.bluez", dbus.ObjectPath("/org/bluez"))
	if err := agentManager.Call("org.bluez.AgentManager1.UnregisterAgent", 0, agentPath).Err; err == nil {

	} else {
		e = errors.WithMessage(err, "UnregisterAgent failed.")
	}
	return
}

func (self *BlueZ) WaitForSignals(callback func(signal *dbus.Signal)) {
	self.SignalChannel = make(chan *dbus.Signal, 10)
	self.Conn.Signal(self.SignalChannel)
	go func() {
		for ch := range self.SignalChannel {
			callback(ch)
		}
	}()
}

func (self *BlueZ) AddInterfacesObserver() (e error) {
	return self.addSignalsObserver("org.freedesktop.DBus.ObjectManager")
}

func (self *BlueZ) RemoveInterfacesObserver() (e error) {
	return self.removeSignalsObserver("org.freedesktop.DBus.ObjectManager")
}

func (self *BlueZ) getManagedObject(path dbus.ObjectPath) (managedObject map[string]map[string]dbus.Variant, e error) {
	if managedObjects, err := self.getManagedObjects(); err == nil {
		managedObject = managedObjects[path]
	} else {
		e = err
	}
	return
}

func (self *BlueZ) getManagedObjects() (managedObjects map[dbus.ObjectPath]map[string]map[string]dbus.Variant, e error) {
	if call := self.Object.Call("org.freedesktop.DBus.ObjectManager.GetManagedObjects", 0); call.Err == nil {
		if err := call.Store(&managedObjects); err == nil {
		} else {
			e = err
		}
	} else {
		e = call.Err
	}
	return
}

func (self *BlueZ) export(instance interface{}, name string, path dbus.ObjectPath, iface string, ifaceIntrospectable string) (e error) {
	if reply, err := self.Conn.RequestName(name, dbus.NameFlagDoNotQueue&dbus.NameFlagReplaceExisting); err == nil {
		if reply == dbus.RequestNameReplyPrimaryOwner {
			if err := self.Conn.Export(instance, path, iface); err == nil {
				if err := self.Conn.Export(introspect.Introspectable(ifaceIntrospectable), path,
					DBusIntrospectableInterface); err == nil {
				} else {
					e = err
				}
			} else {
				e = err
			}
		} else {
			e = errors.Errorf("%s already registered.", iface)
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZ) exportWithProperties(instance interface{}, path dbus.ObjectPath, iface string, ifaceIntrospectable string) (e error) {
	log.Log.Debugf("Exporting %s on %s", path, self.Object.Destination())
	if err := self.Conn.Export(instance, path, iface); err == nil {
		if err := self.Conn.Export(instance, path, DBusPropertiesInterface); err == nil {
			if err := self.Conn.Export(introspect.Introspectable(ifaceIntrospectable), path,
				DBusIntrospectableInterface); err == nil {
			} else {
				e = err
			}
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func (self *BlueZ) exportSingletonWithProperties(instance interface{}, name string, path dbus.ObjectPath, iface string, ifaceIntrospectable string) (e error) {
	if reply, err := self.Conn.RequestName(name, dbus.NameFlagDoNotQueue&dbus.NameFlagReplaceExisting); err == nil {
		if reply == dbus.RequestNameReplyPrimaryOwner {
			if err := self.Conn.Export(instance, path, iface); err == nil {
				if err := self.Conn.Export(instance, path, DBusPropertiesInterface); err == nil {
					if err := self.Conn.Export(introspect.Introspectable(ifaceIntrospectable), path,
						DBusIntrospectableInterface); err == nil {
					} else {
						e = err
					}
				} else {
					e = err
				}
			} else {
				e = err
			}
		} else {
			e = errors.Errorf("%s already registered.", iface)
		}
	} else {
		e = err
	}
	return
}

type BlueZ struct {
	*BlueZObject
	SignalChannel chan *dbus.Signal
}

type BlueZObject struct {
	Conn   *dbus.Conn
	Object dbus.BusObject
}

func (self *BlueZObject) AddPropertiesObserver() (e error) {
	return self.addSignalsObserver("org.freedesktop.DBus.Properties")
}

func (self *BlueZObject) RemovePropertiesObserver() (e error) {
	return self.removeSignalsObserver("org.freedesktop.DBus.Properties")
}

func (self *BlueZObject) addSignalsObserver(iface string) (e error) {
	match := fmt.Sprintf("type='signal',interface='%s',path='%s'", iface, self.Object.Path())
	if call := self.Conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, match); call.Err == nil {
	} else {
		e = call.Err
	}
	return
}

func (self *BlueZObject) removeSignalsObserver(iface string) (e error) {
	match := fmt.Sprintf("type='signal',interface='%s',path='%s'", iface, self.Object.Path())
	if call := self.Conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, match); call.Err == nil {
	} else {
		e = call.Err
	}
	return
}

func MakeFailedError(err error) *dbus.Error {
	return &dbus.Error{
		"org.bluez.Error.Failed",
		[]interface{}{err.Error()},
	}
}

func hasPrefix(path1 dbus.ObjectPath, path2 dbus.ObjectPath) bool {
	return strings.HasPrefix(strings.ToLower(string(path1)), strings.ToLower(string(path2)))
}

func SetLogOutputMode(debug bool, verbose bool, systemd bool) {
	log.SetOutputMode(debug, verbose, systemd)
}
