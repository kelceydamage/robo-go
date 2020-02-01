package sensors

import (
	"time"
	"log"
	"fmt"
)

type Sensor struct {
	port byte
	device byte
	idx byte
	Serialized []byte
}

func (s *Sensor)generateId() {
	s.idx = ((s.port<<4) + s.device) & 0xff;
}

func (s *Sensor)Configure(device byte, port byte) {
	s.device = device
	s.port = port
	s.generateId()
	s.Serialized = []byte{StartByte1, StartByte2, 4, s.idx, 0x01, s.device, s.port}
}

type sensors struct {
	manifest map[int]Sensor
}

// Getter function
func (s *sensors)Get(id int) (sensor Sensor) {
	return s.manifest[id]
}

// Setter functiomn
func (s *sensors)Set(id int, sensor Sensor) {
	s.manifest[id] = sensor
}

// Constructor for Sensors
func SensorPackage(numberOfSensors int) (s sensors) {
	s.manifest = make(map[int]Sensor)
	return s
}

func BufferSensors(sensorPackage sensors, c comm, channel chan []byte) {
	tempBuff := make([]byte, 12)
	for _, sensor := range sensorPackage.manifest {
		fmt.Println("Sending: %v", sensor.Serialized)
		_, err := c.Write(sensor.Serialized)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
			break
		}
		_, err = c.Read(tempBuff)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
			break
		} else {
			channel <- c.Result(CommRecv - 1)
		}
		time.Sleep(2 * time.Millisecond)
	}
}

type comm interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Result(int) ([]byte)
}