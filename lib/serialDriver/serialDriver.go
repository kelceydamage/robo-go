package serialDriver

import (
	"fmt"
	//"errors"
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
}

var SerialState serialState

func (s *serialState)Init() {
	s.start = false
	s.discard = false
	s.head = 0x00
	s.tail = 0x00
	s.counter = -1
	s.length = 0
	s.Complete = false
}

func (s *serialState)ParseIncomming(n int, buff []byte) (err error) {	
	fmt.Printf("New recv buffer, length: %v\n", n)
	fmt.Printf("bufflen: %v\n", len(buff))
	for i := 0; i < n; i++ {
		s.parseSerialByte(buff[i])
	}
	return err
}

func (s *serialState)incrementAndStore(recvByte byte) {
	s.counter++
	fmt.Printf("Counter: %v\n", s.counter)
	s.Buff[s.counter] = recvByte
	fmt.Printf("Storing: %v at index %v\n", recvByte, s.counter)
}

/**************************************************
    ff 55 len idx ... cr(0d) nl(0a) 
    0  1  2   3   n   n+1    n+2     
***************************************************/
func (s *serialState)parseSerialByte(recvByte byte) (err error) {
	//fmt.Printf("parsing byte: %v\n", recvByte)
	var selected bool
	err = nil
	//test
	fmt.Printf("raw byte: %v, counter: %v, length: %v\n", recvByte, s.counter, s.length)
	switch {
	case recvByte == 0xff:
		s.Buff = make([]byte, 12)
		s.Complete = false
		// register start byte
		s.counter = -1
		s.head = 0xff
		selected = true
	case recvByte == 0x55 && s.head == 0xff:
		// confirm full start sequence
		s.start = true
		s.head = 0
		selected = true
		fmt.Printf("Confirmed start sequence\n")
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
			fmt.Printf("FAIL BUCKET: %v\n", recvByte)
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