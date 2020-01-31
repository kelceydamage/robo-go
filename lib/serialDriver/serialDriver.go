package serialDriver

import (
	"fmt"
	"errors"
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
		err := s.parseSerialByte(buff[i])
		if err != nil {
			break
		}
		if s.Complete == true {
			break
		}
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
	var counter int
	err = nil
	//test
	fmt.Printf("raw byte: %v\n, counter: %v, length: %v, sel: %t", recvByte, s.counter, s.length, selected)
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
	/*
	case recvByte == 0x0d:
		// register end byte
		s.tail = 0x0d
		selected = true
	case recvByte == 0x0a && s.tail == 0x0d:
		// confirm full end sequence
		selected = true
		s.start = false
		fmt.Printf("Confirmed CR/NL sequence\n")
		if s.counter - 4 < s.length {
			err = errors.New("Corrupted: package: Too short\n")
			fmt.Printf(err.Error())
		} else {
			fmt.Print("Successful package built\n")
			s.Complete = true
		}
		*/
	default:
		s.tail = 0
		s.head = 0
		// register length
		if s.counter == 2 {
			s.length = int(recvByte)
			counter = 0
			selected = true
		} else if s.counter == 1 {
			// register id
			selected = true
		} else if s.start == true && counter < s.length {
			selected = true
			counter++
		} else if counter >= s.length {
			// set fail flag if too many data bytes
			err = errors.New("Corrupted: package Too long\n")
			fmt.Printf(err.Error())
			s.counter = -1
		} else if counter == s.length - 1 {
			selected = true
			s.Complete = true
			fmt.Print("Successful package built\n")
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