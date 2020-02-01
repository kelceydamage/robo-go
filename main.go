package main

import (
	//"log"
	"fmt"
	//"time"

	"github.com/kelceydamage/robo-go/lib/serialDriver"
	"github.com/kelceydamage/robo-go/lib/sensors"
)

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
*/

// Global objects
var serial = serialDriver.SerialState
var sensorPackage = sensors.SensorPackage(1)
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

	// Configure sensors
	sensors.Ultrasonic.Configure(1, 8)
	sensorPackage.Set(0, sensors.Ultrasonic)
}

func main() {
	defer serial.Close()

	sensors.BufferSensors(sensorPackage, &serial, sensorFeed)

	result := <- sensorFeed
	fmt.Printf("%v\n", result)

}