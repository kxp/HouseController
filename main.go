package main

import (
	"HouseController/serverhandler"
	"HouseController/tools"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// muxer for the hhtp server, its a blobal
var mux map[string]func(http.ResponseWriter, *http.Request)

func ConfigServer() (*http.Server, error) {

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = serverhandler.Hello
	mux["/alarm"] = serverhandler.Alarm
	mux["/light"] = serverhandler.Light

	return &server, nil
}

func OpenSerial() (*tools.Tty, error) {

	fmt.Println("opening...")
	var serialPort tools.SerialPort

	serialPort = &tools.Tty{}

	serialPort.Open("test")
	serialPort.Close()

	return nil, nil
}

func main() {

	var settings = tools.LoadConfiguration("configs.json")

	httpServer, err := ConfigServer()
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
