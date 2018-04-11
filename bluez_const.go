package bluez

import "github.com/godbus/dbus/introspect"

const (
	DBusIntrospectableInterface = "org.freedesktop.DBus.Introspectable"
	DBusPropertiesInterface     = "org.freedesktop.DBus.Properties"
	DBusPropertiesIntro         = `
   <interface name="org.freedesktop.DBus.Properties">
      <method name="Get">
         <arg name="interface" type="s" direction="in" />
         <arg name="name" type="s" direction="in" />
         <arg name="value" type="v" direction="out" />
      </method>
      <method name="Set">
         <arg name="interface" type="s" direction="in" />
         <arg name="name" type="s" direction="in" />
         <arg name="value" type="v" direction="in" />
      </method>
      <method name="GetAll">
         <arg name="interface" type="s" direction="in" />
         <arg name="properties" type="a{sv}" direction="out" />
      </method>
      <signal name="PropertiesChanged">
         <arg name="interface" type="s" />
         <arg name="changed_properties" type="a{sv}" />
         <arg name="invalidated_properties" type="as" />
      </signal>
   </interface>
`
	GattService1Interface = "org.bluez.GattService1"
	GattService1Intro     = `<node>
<interface name="org.bluez.GattService1">
      <property name="UUID" type="s" access="read" />
      <property name="Primary" type="b" access="read" />
</interface>` + DBusPropertiesIntro + introspect.IntrospectDataString + `</node>`

	GattCharacteristic1Interface = "org.bluez.GattCharacteristic1"
	GattCharacteristic1Intro     = `<node>
<interface name="org.bluez.GattCharacteristic1">
      <property name="UUID" type="s" access="read" />
</interface>` + DBusPropertiesIntro + introspect.IntrospectDataString + `</node>`

	ObjectManagerInterface = "org.freedesktop.DBus.ObjectManager"
	ObjectManagerIntro     = `<node>
<interface name="org.freedesktop.DBus.ObjectManager">
	<method name="GetManagedObjects">
		<arg name="objects" type="a{oa{sa{sv}}}" direction="out"/>
	</method>
	<signal name="InterfacesAdded">
		<arg name="Object" type="o"/>
		<arg name="interfaces" type="a{sa{sv}}"/>
	</signal>
	<signal name="InterfacesRemoved">
		<arg name="Object" type="o"/>
		<arg name="interfaces" type="as"/>
	</signal>
</interface>` + DBusPropertiesIntro + introspect.IntrospectDataString + `</node>`

	LEAdvertisement1Interface = "org.bluez.LEAdvertisement1"
	LEAdvertisement1Intro     = `
<node>
   <interface name="org.freedesktop.DBus.Introspectable">
      <method name="Introspect">
         <arg name="xml" type="s" direction="out" />
      </method>
   </interface>
   <interface name="org.bluez.LEAdvertisement1">
      <method name="Release" />
      <property name="Type" type="s" access="read"  />
      <property name="ServiceUUIDs" type="as" access="read"  />
      <property name="Includes" type="as" access="read"  />
      <property name="LocalName" type="s" access="read"  />
      <property name="Diration" type="u" access="read"  />
      <property name="Timeout" type="u" access="read"  />
   </interface>
   <interface name="org.freedesktop.DBus.Properties">
      <method name="Get">
         <arg name="interface" type="s" direction="in" />
         <arg name="name" type="s" direction="in" />
         <arg name="value" type="v" direction="out" />
      </method>
      <method name="Set">
         <arg name="interface" type="s" direction="in" />
         <arg name="name" type="s" direction="in" />
         <arg name="value" type="v" direction="in" />
      </method>
      <method name="GetAll">
         <arg name="interface" type="s" direction="in" />
         <arg name="properties" type="a{sv}" direction="out" />
      </method>
      <signal name="PropertiesChanged">
         <arg name="interface" type="s" />
         <arg name="changed_properties" type="a{sv}" />
         <arg name="invalidated_properties" type="as" />
      </signal>
   </interface>
</node>`
)
