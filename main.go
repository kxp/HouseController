package main

import (
	"HouseController/alarm"
	"HouseController/light"
	"HouseController/tools"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// muxer for the hhtp server, its a blobal
var mux map[string]func(http.ResponseWriter, *http.Request)

// Hello function only exists for debug
func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
	fmt.Println("we arrive on hello")
}

func ConfigServer(settings tools.Config) (*http.Server, error) {

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	server := http.Server{
		Addr:    settings.HttpPort,
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = Hello
	mux["/alarm"] = alarm.Alarm
	mux["/light"] = light.GetInstance().Light //light.Light

	return &server, nil
}

func main() {

	instanceTools := tools.GetInstance()
	settings, err := instanceTools.LoadConfiguration(tools.SettingsFilename)
	if err != nil || settings == nil {
		log.Fatalf("Fail to open settings file: %v", err)
		return
	}

	//nitializes the  serial port
	light.GetInstance().Initialize(settings.SerialPort)

	httpServer, err := ConfigServer(*settings)
	if err != nil {
		log.Fatalf("Fail to open the HTTP server: %v", err)
		return
	}

	//Blocking call
	httpServer.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h, ok := mux[r.URL.String()]
	if ok != false {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
