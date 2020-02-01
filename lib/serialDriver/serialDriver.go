package serialDriver

import (
	//"fmt"
	//"errors"
	"log"
	"io"

	"github.com/jacobsa/go-serial/serial"
)

type serialState struct {
	Buff []byte
	start bool
	discard bool
	head byte
	tail byte
	counter int
	length int
	Complete bool
	err error
	port io.ReadWriteCloser
}

var SerialState serialState
type OpenOptions = serial.OpenOptions

func (s *serialState)Open(options serial.OpenOptions) {
	s.start = false
	s.discard = false
	s.head = 0x00
	s.tail = 0x00
	s.counter = -1
	s.length = 0
	s.Complete = false
	s.port, s.err = serial.Open(options)
	if s.err != nil {
		log.Fatalf("port.Open: %v", s.err)
	}
}

func (s *serialState)Write(msg []byte) (bytesWritten int, err error) {
	return s.port.Write(msg)
}

func (s *serialState)Read(buff []byte) (bytesRead int, err error) {
	return s.port.Read(buff)
}

func (s *serialState)Close() {
	s.port.Close()
}

func (s *serialState)ParseIncomming(n int, buff []byte) (err error) {	
	for i := 0; i < n; i++ {
		s.parseSerialByte(buff[i])
	}
	return err
}

func (s *serialState)incrementAndStore(recvByte byte) {
	s.counter++
	s.Buff[s.counter] = recvByte
}

/**************************************************
    ff 55 len idx ... cr(0d) nl(0a) 
    0  1  2   3   n   n+1    n+2     
***************************************************/
func (s *serialState)parseSerialByte(recvByte byte) (err error) {
	var selected bool
	err = nil
	switch {
	// register start byte
	case recvByte == 0xff:
		s.Buff = make([]byte, 12)
		s.Complete = false
		s.counter = -1
		s.head = 0xff
		selected = true
	// confirm full start sequence
	case recvByte == 0x55 && s.head == 0xff:
		s.start = true
		s.head = 0
		selected = true
	// All other bytes
	default:
		s.tail = 0
		s.head = 0
		// register length
		if s.counter == 2 {
			s.length = int(recvByte)
			selected = true
		} else if s.counter == 1 {
			// register id
			selected = true
		} else if s.length > 1 {
			selected = true
			s.length--
		} else if s.length == 1 {
			selected = true
			s.length = 0
			s.Complete = true
		} else {
			selected = false
		}
	}
	if selected == true {
		s.incrementAndStore(recvByte)
		selected = false
		if s.Complete == true {
			s.counter = -1
		}
	}
	return err
}