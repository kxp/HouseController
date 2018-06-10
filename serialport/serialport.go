/*
The Serial port requires permissions, make sure you add the group to the current user.
sudo usermod -a -G dialout/uucp USER
More info:
https://github.com/GoldenCheetah/GoldenCheetah/wiki/Allowing-your-linux-userid-permission-to-use-your-usb-device
DO NOT RUN A WEBSERVER AS ROOT!
*/

package serialport

import (
	"errors"
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

// this class should be a singleton

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
		PortName:        portName,
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("Port name: %s. Error while openning: %v", portName, err)
		return err
	}
	internalHandler.PortHandler = port

	return nil
}

func (internalHandler *Tty) Close() error {

	if internalHandler.PortHandler == nil {
		return errors.New("No port found")
	}

	// Make sure to close it later.
	defer internalHandler.PortHandler.Close()
	internalHandler.PortHandler = nil
	log.Println("Serial port Closed")

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
