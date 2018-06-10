package light

import (
	"HouseController/serialport"
	"HouseController/tools"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var results []string

// singleton stuff

var instance *LightManager
var once sync.Once

func GetInstance() *LightManager {
	once.Do(func() {
		instance = &LightManager{}
	})
	return instance
}

// LightManager its the class that manages the light, and is a singleton because of serial port
type LightManager struct {
	device serialport.SerialPort
	eState int
}

func (manager *LightManager) Light(w http.ResponseWriter, r *http.Request) {
	if manager.eState != tools.Ready {
		log.Println("The serial port wasn't created in the request")
		manager.Initialize("/dev/ttyUSB0")
	}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		//does the JSON parsing internally
		results = append(results, string(body))

		fmt.Fprint(w, "POST done")
	} else {
		//http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		io.WriteString(w, "Hello world! from Light")
		fmt.Println("Light: ", r.Body)
	}
}

func (manager *LightManager) Initialize(deviceName string) {

	if manager.eState != 0 {
		return
	}

	manager.device = &serialport.Tty{}
	err := manager.device.Open(deviceName)
	if err != nil {
		log.Println("Fail to open the serial port: %s", err)
		manager.eState = tools.NotInitialized
		return
	}
	manager.eState = tools.Ready
}

func (manager *LightManager) ReadLightFromDevice() {

}

func (manager *LightManager) SetLightColor() {

}
