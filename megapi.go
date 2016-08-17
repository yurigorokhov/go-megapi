/*
This package is used to control motors that are connected to a MegaPi device. This work is largely based off of: https://github.com/Makeblock-official
*/
package megapi

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/tarm/serial"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

const Baud = 115200

// Send motor commands to the MegaPi
type MegaPi struct {
	serialPort io.ReadWriteCloser
}

// Create new MegaPi for sending commands
func NewMegaPi(device string) (*MegaPi, error) {
	c := &serial.Config{Name: device, Baud: Baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	return &MegaPi{
		serialPort: s,
	}, nil
}

// Run a dc motor
func (p *MegaPi) DcMotorRun(port byte, speed int16) error {

	//NOTE: executing this twice because of: https://github.com/Makeblock-official/PythonForMegaPi/issues/1
	if err := p.dcMotorRun_Helper(port, speed+1); err != nil {
		return err
	}
	if err := p.dcMotorRun_Helper(port, speed); err != nil {
		return err
	}
	return nil
}

// Stop a dc motor
func (p *MegaPi) DcMotorStop(port byte) error {

	//NOTE: executing this twice because of: https://github.com/Makeblock-official/PythonForMegaPi/issues/1
	if err := p.dcMotorRun_Helper(port, -1); err != nil {
		return err
	}
	if err := p.dcMotorRun_Helper(port, 0); err != nil {
		return err
	}
	return nil

}

// Close the connection to the MegaPi
func (p *MegaPi) Close() error {
	return p.serialPort.Close()
}

func (p *MegaPi) dcMotorRun_Helper(port byte, speed int16) error {
	bufOut := new(bytes.Buffer)

	// byte sequence: 0xff, 0x55, id, action, device, port
	bufOut.Write([]byte{0xff, 0x55, 0x6, 0x0, 0x2, 0xa, port})
	binary.Write(bufOut, binary.LittleEndian, speed)
	bufOut.Write([]byte{0xa})
	_, err := bufOut.WriteTo(p.serialPort)
	time.Sleep(5 * time.Millisecond)
	return err
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
