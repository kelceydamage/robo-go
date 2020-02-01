package main

import (
	"log"
	"fmt"
	//"time"
	//"bytes"
	//"os"
	//"strconv"

	//"go.bug.st/serial"
	"github.com/kelceydamage/robo-go/lib/serialDriver"
	"github.com/kelceydamage/robo-go/lib/sensors"
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



func bufferSensors() {
	
}

var serial = serialDriver.SerialState
var sensorPackage sensors.Sensors
var sensorFeed = make(chan []byte, 512)

func init() {
	// Configure communications
	options := serialDriver.OpenOptions{
		PortName: "/dev/ttyTHS1",
		BaudRate: 76800, // Best|stable option for using Jetson and Megapi
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,
		InterCharacterTimeout: 1,
	}
	serial.Open(options)
	defer serial.Close()

	// Configure sensors
	sensors.Ultrasonic.Configure(1, 8)
	sensorPackage.Set(0, sensors.Ultrasonic)
}

func main() {

	// Set up options.
	var tbuff = make([]byte, 16)

	n, err := serial.Write(sensorPackage.Get(0).Serialized)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}
	fmt.Printf("WRITE %v\n", n)
	
	for {
		n, err = serial.Read(tbuff)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
			break
		}
		fmt.Printf("READ %v\n", n)
		for _, b := range tbuff {
			fmt.Printf("%v ", b)
		}
		fmt.Printf("\n")
		err = serial.ParseIncomming(n, tbuff)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
			break
		}
		if serial.Complete == true {
			break
		}
	}
	fmt.Printf("\nReceived: ")
	for _, b := range serial.Buff {
		fmt.Printf("%v ", b)
	}
	fmt.Printf("\n")
}