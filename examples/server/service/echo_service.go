package service

import (
	"bluez-connect"
	"log"
)

const (
	GattEchoServiceUUID        = "C37D52CE-3E08-420A-BBD7-2F7C656884C8"
	GattEchoCharacteristicUUID = "F1FE1C00-FD56-4BDB-A752-B36BB005B2B3"
)

func NewService() *bluez.GattService {

	thisService.gattService = bluez.NewGattService(GattEchoServiceUUID)

	thisService.echoCharacteristic = bluez.NewGattCharacteristic(GattEchoCharacteristicUUID, []string{"read", "write"})
	thisService.gattService.AddCharacteristic(thisService.echoCharacteristic)

	thisService.echoCharacteristic.OnWriteFunc = func(value []byte) (e error) {
		log.Printf("Write data with size: %d\f", len(value))
		thisService.echoData = value
		return nil
	}
	thisService.echoCharacteristic.OnReadFunc = func() (value []byte, e error) {
		return thisService.echoData, nil
	}

	thisService.echoData = []byte("Hello, World !!!")

	return thisService.gattService
}

type echoService struct {
	gattService        *bluez.GattService
	echoCharacteristic *bluez.GattCharacteristic
	echoData []byte
}

var thisService = new(echoService)
