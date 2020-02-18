package drivers

import (
	"fmt"
	//"errors"
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

type serialState struct {
	Buff     []byte
	start    bool
	discard  bool
	head     byte
	tail     byte
	counter  int
	length   int
	Complete bool
	err      error
	prevByte byte
	port     io.ReadWriteCloser
}

// SerialState is the driver object to be used for communicating with
// MegaPi control board.
var SerialState serialState

// OpenOptions are the options object for the serial interface.
type OpenOptions = serial.OpenOptions

func (s *serialState) Open(options serial.OpenOptions) {
	s.start = false
	s.prevByte = 0x00
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

func (s *serialState) Write(msg []byte) (bytesWritten int, err error) {
	return s.port.Write(msg)
}

func (s *serialState) Read(buff []byte) (bytesRead int, err error) {
	var n int
	for {
		//fmt.Println("Buffer cycle")
		n, err := s.port.Read(buff)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
			s.err = err
			fmt.Printf("error: %v\n", err)
		}
		if n == 0 {
			n = 12
		}
		fmt.Printf("tempbuff: %v, N: %v\n", buff, n)
		s.parseIncomming(n, buff)
		if s.Complete == true || s.discard == true {
			break
		}
	}
	return n, err
}

func (s *serialState) Result(n int) (buff []byte) {
	return s.Buff[0:n]
}

func (s *serialState) Close() {
	// not yet implemented
	// s.port.Close()
}

func (s *serialState) parseIncomming(n int, buff []byte) {
	for i := 0; i < n; i++ {
		s.parseSerialByte(buff[i])
		if s.discard == true {
			break
		}
	}
}

func (s *serialState) incrementAndStore(recvByte byte) {
	s.counter++
	fmt.Printf("I&C\n Count: %v\n", s.counter)
	s.Buff[s.counter] = recvByte
}

/**************************************************
    ff 55 len idx ... cr(0d) nl(0a)
    0  1  2   3   n   n+1    n+2
***************************************************/
func (s *serialState) parseSerialByte(recvByte byte) {
	var selected = true
	var err error
	switch {
	// confirm full start sequence
	case recvByte == 0x55 && s.prevByte == 0xff:
		s.counter = -1
		s.incrementAndStore(recvByte)
		s.Buff = make([]byte, 12)
		s.Complete = false
		s.counter = 0
	// All other bytes
	default:
		s.tail = 0
		s.head = 0
		// register length
		if s.counter == 2 {
			s.length = int(recvByte)
		} else if s.counter == 1 {
			// register id
		} else if s.length > 1 {
			s.length--
		} else if s.length == 1 {
			s.length = 0
			s.Complete = true
		} else {
			selected = false
			s.counter = -1
		}
	}
	s.prevByte = recvByte
	if selected == true {
		s.incrementAndStore(recvByte)
		if s.Complete == true {
			s.counter = -1
		}
	}
	s.err = err
}
