package sensors

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Sensor is the representation of a physical sensor on the controller.
// It contains the serial code to retrieve data from the ophysical sensor, along
// with the device, port, and id.
type Sensor struct {
	port       byte
	device     byte
	idx        byte
	Serialized []byte
}

func (s *Sensor) generateID() {
	s.idx = ((s.port << 4) + s.device) & 0xff
}

// Configure generates the serial code for calling the sensor.
func (s *Sensor) Configure(device byte, port byte) {
	s.device = device
	s.port = port
	s.generateID()
	s.Serialized = []byte{StartByte1, StartByte2, 4, s.idx, 0x01, s.device, s.port}
}

// Sensors is a map designed to store a Sensor at any given index.
type Sensors struct {
	manifest map[int]Sensor
}

// Get a particular sensor at a given index from the manifest.
func (s *Sensors) Get(id int) (sensor Sensor) {
	return s.manifest[id]
}

// Set a particular Sensor at a given index in the manifest.
func (s *Sensors) Set(id int, sensor Sensor) {
	s.manifest[id] = sensor
}

// SensorPackage constructor.
func SensorPackage(numberOfSensors int) (s Sensors) {
	s.manifest = make(map[int]Sensor)
	return s
}

// BufferSensors is a go routine that continually loops through the Sensors and
// writes their data to a channel.
func BufferSensors(wg *sync.WaitGroup, sensorPackage Sensors, c comm, channel chan []byte) {
	defer wg.Done()
	for {
		tempBuff := make([]byte, 12)
		for _, sensor := range sensorPackage.manifest {
			fmt.Printf("Sending: %v\n", sensor.Serialized)
			_, err := c.Write(sensor.Serialized)
			if err != nil {
				log.Fatalf("port.Read: %v", err)
				break
			}
			//fmt.Println("Receiving ...")
			_, err = c.Read(tempBuff)
			if err != nil {
				log.Fatalf("port.Read: %v", err)
				break
			} else {
				fmt.Printf("Adding to channel: %v\n", c.Result(CommRecv))
				channel <- c.Result(CommRecv)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
}

type comm interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Result(int) []byte
}
