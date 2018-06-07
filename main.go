package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
	//"encoding/json"
	"helper"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	fmt.Println("opening...")
	var serialPort helper.SerialPort

	serialPort = &helper.Tty{}

	serialPort.Open("test")
	serialPort.Close()

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello
	mux["/alarm"] = Alarm
	mux["/light"] = Light
	server.ListenAndServe()
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
