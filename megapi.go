/*
This package is used to control motors that are connected to a MegaPi device. This work is largely based off of: https://github.com/Makeblock-official
*/
package megapi

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/huin/goserial"
	"io"
	"io/ioutil"
	"strings"
)

const Baud = 9600

// Send motor commands to the MegaPi
type MegaPi struct {
	serialPort io.ReadWriteCloser
}

// Create new MegaPi for sending commands
func NewMegaPi(device string) (*MegaPi, error) {
	c := &goserial.Config{Name: device, Baud: Baud}
	s, err := goserial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	return &MegaPi{
		serialPort: s,
	}, nil
}

func (p *MegaPi) MotorRun(port uint8, speed int16) error {
	bufOut := new(bytes.Buffer)
	bufOut.Write([]byte{0xff, 0x55, 0x6, 0x0, 0x2, 0xa, port})
	err := binary.Write(bufOut, binary.LittleEndian, speed)
	if err != nil {
		return err
	}
	_, err = p.serialPort.Write(bufOut.Bytes())
	return err
}

// Close the connection to the MegaPi
func (p *MegaPi) Close() error {
	return p.serialPort.Close()
}

// Finds a MegaPi device connected over usb
func Find_megapi_usb_device() (string, error) {

	// look for something that looks like "/dev/ttyUSB"
	contents, err := ioutil.ReadDir("/dev")
	if err != nil {
		return "", err
	}
	for _, f := range contents {
		if strings.Contains(f.Name(), "ttyUSB") {
			return "/dev/" + f.Name(), nil
		}
	}
	return "", errors.New("Could not find a device '/dev/ttyUSB*', ensure that the MegaPi is connected over USB")
}
