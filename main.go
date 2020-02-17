package main

import (
	//"log"
	"fmt"
	"sync"
	"time"

	"github.com/kelceydamage/robo-go/lib/drivers"
	"github.com/kelceydamage/robo-go/lib/sensors"
)

/**************************************************
    ff 55 len idx action device port slot data a
    0  1  2   3   4      5      6    7    8
**************************************************
[]byte{255, 85, 6, 0, 2, 10, port, speed}
*/

// Global objects
var serial = drivers.SerialState
var sensorPackage = sensors.SensorPackage(1)
var sensorFeed = make(chan []byte, 512)

func init() {
	// Configure communications
	options := drivers.OpenOptions{
		PortName:              "/dev/ttyTHS1",
		BaudRate:              76800, // Best|stable option for using Jetson and Megapi
		DataBits:              8,
		StopBits:              1,
		MinimumReadSize:       4,
		InterCharacterTimeout: 1,
	}
	serial.Open(options)

	// Configure sensors
	sensors.Ultrasonic.Configure(1, 8)
	sensorPackage.Set(0, sensors.Ultrasonic)
}

func main() {
	defer serial.Close()
	var wg sync.WaitGroup

	// Still need some safety around threading this.
	wg.Add(1)
	go sensors.BufferSensors(&wg, sensorPackage, &serial, sensorFeed)

	time.Sleep(1 * time.Second)

	// Main loop
	counter := 0
	for {
		result := <-sensorFeed
		fmt.Printf("Receiving: %v\n", result)
		counter++
		fmt.Printf("Receiving Count: %v Channel Occupancy: %v\n", counter, len(sensorFeed))
	}

	// Will fail unless BufferSensors is set to infinite loop
	wg.Wait()

	//fmt.Println("main finished")
}
