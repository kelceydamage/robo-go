package main

import (
	"log"
	"fmt"
	"time"
	//"bytes"
	//"os"
	//"strconv"

	//"go.bug.st/serial"
	"github.com/jacobsa/go-serial/serial"
	"github.com/kelceydamage/robo-go/lib/serialDriver"
)
/*
func getActiveSerialPorts() ([]string) {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
	return ports
}
*/

/* Mode options:
 * BaudRate: 57600,
 * Parity: serial.EvenParity,
 * DataBits: 7,
 * StopBits: serial.OneStopBit,
*/
/*
func openSerialPort(port string) (serial.Port) {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	fmt.Printf("Opening port: %v\n", port)
	openPort, err := serial.Open(port, mode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Opened port: %v\n", openPort)
	
	return openPort
}
*/

/**************************************************
    ff 55 len idx action device port
    0  1  2   3   4      5      6   
***************************************************/
/*
* package: {ff, 55, len, idx, action, device, port}
*/
/*
func getSensor(port byte, device byte, action byte) {
	var idx byte = ((port<<4) + device) & 0xff;
	//var msg = fmt.Sprintf("255,85,4,%v %v %v %v", idx, action, device, port)

	var data = []byte{255, 85, 4, idx, action, device, port, 0x0d, 0x0a}
	fmt.Printf("mystr:\t %v \n", data)

	//mystr := "\n\r"
	//data = append(data, mystr...)
	fmt.Printf("mystr:\t %v \n", data)

	writeSerial(data)
}
*/
/**************************************************
    ff 55 len idx action device port slot data a
    0  1  2   3   4      5      6    7    8
***************************************************/
/* 
* package: {ff, 55, len, idx, action, device, port, slot, data, a}
* 7 & 8 are read as a 2 byte short int signed.
*/
/*
func getMotor(port byte, speed byte) {
	//var msg = fmt.Sprintf("255,85,4,%v %v %v %v", idx, action, device, port)

	var data = []byte{255, 85, 6, 0, 2, 10, port, speed}
	fmt.Printf("mystr:\t %v \n", data)

	//mystr := "\n\r"
	//data = append(data, mystr...)
	fmt.Printf("mystr:\t %v \n", data)

	writeSerial(data)
}

func writeSerial(command []byte) {
	n, err := openPort.Write(command)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)
}
*/
//func readSerial() (data []byte) {
//
//	//return 
//}


//var ports = getActiveSerialPorts()
//var openPort = openSerialPort(ports[0])

//var serialBuffer = make(chan byte, 512)

func main() {
	// Set up options.
	options := serial.OpenOptions{
		PortName: "/dev/ttyAMA0",
		BaudRate: 38400,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,
	}

	var tbuff = make([]byte, 16)
	var idx byte = ((8<<4) + 1) & 0xff;
	var datatest0 = []byte{255, 85, 4, idx, 1, 1, 8, 0x0d, 0x0a}
	
	//var datatest1 = []byte{5, 6, 1, 255, 85, 4, idx, 1, 1, 8, 0x0d, 0x0a, 5, 3 ,55}
	//var datatest2 = []byte{5, 6, 1, 255, 85, idx}
	//var datatest3 = []byte{4, 1, 1, 8, 0x0d, 0x0a, 5, 3 ,55}
	//var datatest4 = []byte{255, 85, 4, idx, 1, 1, 8, 0x0d, 0x00, 0x0a, 5, 3 ,55}
	//var datatest5 = []byte{255, 85, 4, idx, 1, 8, 0x0d, 0x0a, 5, 3 ,55}
	
	//var buff = make([]byte, 32)

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()
	//getSensor(8, 1, 1)


	_serial := serialDriver.SerialState
	_serial.Init()

	for i := 0; i < 10; i++ {
		time.Sleep(16 * time.Millisecond)
		n, err := port.Write(datatest0)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}

		fmt.Printf("WRITE %v\n", n)
		// Write 4 bytes to the port.
		for {
			n, err = port.Read(tbuff)
			if err != nil {
				log.Fatalf("port.Read: %v", err)
				break
			}
			fmt.Printf("READ %v\n", n)
			for _, b := range tbuff {
				fmt.Printf("%v ", b)
			}
			fmt.Printf("\n")
			err = _serial.ParseIncomming(n, tbuff)
			if err != nil {
				log.Fatalf("port.Read: %v", err)
				break
			}
			if _serial.Complete == true {
				break
			}
		}
		fmt.Printf("\nReceived: ")
		for _, b := range _serial.Buff {
			fmt.Printf("%v ", b)
		}
		fmt.Printf("\n")
	}

	
	/*
	tbuff[n] = 0x0d
	tbuff[n+1] = 0x0a
	fmt.Printf("MOD %v\n", n+2)
	for _, b := range tbuff {
		fmt.Printf("%v ", b)
	}
	fmt.Printf("\n")
	*/
	/*
	_serial := serialDriver.SerialState
	_serial.Init()

	//_serial.ParseIncomming(13, datatest4)
	//_serial.ParseIncomming(11, datatest5)
	_serial.ParseIncomming(6, datatest2)
	_serial.ParseIncomming(9, datatest3)

	fmt.Printf("\nReceived: ")
	for _, b := range _serial.Buff {
		fmt.Printf("%v ", b)
	}
	fmt.Printf("\n")
	*/

	
	  
}