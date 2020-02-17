package sensors

import (
	"testing"
)

func TestConst(t *testing.T) {
	var c byte
	c = StartByte1
	if c != 0xff {
		t.Errorf("Start byte for serial code is incorrect. Expected: 0xff, Got: %x", c)
	}

	c = StartByte2
	if c != 0x55 {
		t.Errorf("Confirmation byte for serial code is incorrect. Expected: 0x55, Got: %x", c)
	}

	c = CommRecv
	if c != 8 {
		t.Errorf("Recv byte length for serial code is incorrect. Expected: 0x08, Got: %x", c)
	}

	c = CommSend
	if c != 7 {
		t.Errorf("Send byte length for serial code is incorrect. Expected: 0x07, Got: %x", c)
	}
}
