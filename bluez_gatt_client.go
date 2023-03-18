package bluez

type GattApplication struct {
	Name     string
	Path     string
	Services map[string]*GattService
}

type GattService struct {
	UUID            string
	Characteristics map[string]*GattCharacteristic
}

type GattCharacteristic struct {
	UUID             string
	Flags            []string
	OnReadFunc       func() ([]byte, error)
	OnWriteFunc      func([]byte) ([]byte, error)
	OnWriteAsyncFunc func([]byte) ([]byte, error)
	OnReadValueFunc  func([]byte) ([]byte, error)
}

func (self *GattApplication) AddService(gattService *GattService) {
	self.Services[gattService.UUID] = gattService
}

func (self *GattApplication) RemoveService(gattService *GattService) {
	delete(self.Services, gattService.UUID)
}

func (self *GattService) AddCharacteristic(gattCharacteristic *GattCharacteristic) {
	self.Characteristics[gattCharacteristic.UUID] = gattCharacteristic
}

func (self *GattService) RemoveCharacteristic(gattCharacteristic *GattCharacteristic) {
	delete(self.Characteristics, gattCharacteristic.UUID)
}

func NewGattApplication(name string, path string) (gattApplication *GattApplication) {
	gattApplication = &GattApplication{Name: name, Path: path}
	gattApplication.Services = make(map[string]*GattService, 0)
	return
}

func NewGattService(uuid string) (gattService *GattService) {
	gattService = &GattService{UUID: uuid}
	gattService.Characteristics = make(map[string]*GattCharacteristic, 0)
	return
}

func NewGattCharacteristic(uuid string, flags []string) (gattCharacteristic *GattCharacteristic) {
	gattCharacteristic = &GattCharacteristic{UUID: uuid, Flags: flags}
	return
}
