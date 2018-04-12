# bluez-connect = GATT of BlueZ
## Package provides a Bluetooth Low Energy GATT implementation over linux BlueZ API.


**bluez-connect** communicates with BlueZ over D-Bus (linux message bus system).


This package was developed as part of IoT project in order to add GATT capability to Raspberry Pi like devices. Package was tested against BlueZ 5.46

Because this package mostly makes sense for Linux / arm devices, build environment is adopted to such platform.
In other words ```GOOS=linux GOARCH=arm```  


## Usage
Please see [godoc.org](http://godoc.org/github.com/mark2b/bluez-connect) for documentation. (Not ready yet)

## Examples

### Peripheral example
This example creates and advertizes, **Echo** service with single **Echo** characteristic with Read / Write capabilities.
 
Build and run
```
GOOS=linux GOARCH=arm go build examples/server/server-example.go 
# Copy to target device
./server-example

```
### Central example
This example is only template. You need to replace Service and Characteristics by real ones you want to connect.

Build and run
```
GOOS=linux GOARCH=arm go build examples/client/client-example.go 
# Copy to target device
./client-example

```

Package release under a [MIT license](./LICENSE.md).