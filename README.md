# Package bluez-connect provides a Bluetooth Low Energy GATT implementation over linux BlueZ API.

BlueZ - Official Linux Bluetooth protocol stack

GATT protocol allows creation BLE peripherals and centrals.

bluez-connect communicates with BlueZ over D-Bus (linux message bus system).


Create blueZ object (handles D-Bus connection)

*blueZ, err := bluez.NewBLueZ()


*Register signals callback

blueZ.WaitForSignals(onBlueZSignal)

## Peripheral Role - (Server)

Peripheral defines GATT Services and Characteristics allowing Read, Write and Send notifications


## Central Role - (Client)

Central discovers peripheral by name, address or services sets.