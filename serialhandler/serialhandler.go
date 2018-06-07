/*
The Serial port requires permissions, make sure you add the group to the current user.
sudo usermod -a -G dialout/uucp USER
More info:
https://github.com/GoldenCheetah/GoldenCheetah/wiki/Allowing-your-linux-userid-permission-to-use-your-usb-device
DO NOT RUN A WEBSERVER AS ROOT!
*/

package serialhandler

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

type SerialPort interface {
	Open(string) error
	Close() error
	WriteBytes([]byte) (int, error)
}

type Tty struct {
	PortHandler io.ReadWriteCloser
}

func (internalHandler *Tty) Open(portName string) error {

	if internalHandler.PortHandler == nil {
		return errors.New("Port Already in use")
	}

	// Set up options.
	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
		return err
	}

	internalHandler.PortHandler = port
	fmt.Println("Port openned ", portName)

	return nil
}

func (internalHandler *Tty) Close() error {

	if internalHandler.PortHandler == nil {
		return errors.New("No port found")
	}

	// Make sure to close it later.
	defer internalHandler.PortHandler.Close()
	internalHandler.PortHandler = nil
	// Write 4 bytes to the port.
	//b := []byte{0x00, 0x01, 0x02, 0x03}
	//n, err := port.Write(b)
	//if err != nil {
	//log.Fatalf("port.Write: %v", err)
	//return err
	//}

	fmt.Println("Inside close")
	//fmt.Println("Wrote", n, "bytes.")
	return nil

}

// WriteBytes write the provided arry into the serial port
// Returns the written bytes, and the error message in case of failiure
func (internalHandler *Tty) WriteBytes(message []byte) (int, error) {

	if internalHandler.PortHandler == nil {
		return 0, errors.New("No port found")
	}

	n, err := internalHandler.PortHandler.Write(message)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
		return 0, err
	}
	return n, nil
}
