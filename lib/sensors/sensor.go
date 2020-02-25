// Copyright (c) 2020 Author Name. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sensors

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/kelceydamage/robo-go/lib/drivers"
)

// PhysicalSensor is the representation of a physical sensor on the controller.
// It contains the serial code to retrieve data from the ophysical sensor, along
// with the device, port, and id.
type PhysicalSensor struct {
	port       byte
	device     byte
	idx        byte
	serialCode [SerialCodeLength]byte
	driver     drivers.Driver
}

// GetPort returns the physical sensor port.
func (s *PhysicalSensor) GetPort() byte {
	return s.port
}

// GetDevice returns the physical sensor device type.
func (s *PhysicalSensor) GetDevice() byte {
	return s.device
}

// GetID returns the physical sensor ID.
func (s *PhysicalSensor) GetID() byte {
	return s.idx
}

// GetSerialCode returns the physical sensor serial code.
func (s *PhysicalSensor) GetSerialCode() [SerialCodeLength]byte {
	return s.serialCode
}

func (s *PhysicalSensor) Read() SensorReading {
	var reading sensorReading
	_, err := s.driver.Write(s.GetSerialCode())
	if err != nil {
		fmt.Println("Error Writing Driver")
	}
	_, err = s.driver.Read()
	if err != nil {
		fmt.Println("Error Reading Driver")
	}
	reading.sensor = s
	reading.value = s.asFloat(s.driver.Result(CommRecv)[4:])
	return &reading
}

func (s *PhysicalSensor) asFloat(bytes []byte) float32 {
	binrep := binary.LittleEndian.Uint32(bytes)
	floatrep := *(*float32)(unsafe.Pointer(&binrep))
	fmt.Printf("Converting: %v, %v -> %v\n", bytes, binrep, floatrep)
	return floatrep
}

func (s *PhysicalSensor) generateID() {
	s.idx = ((s.GetPort() << 4) + s.GetDevice()) & 0xff
}

// Initialize generates the serial code for calling the sensor.
func (s *PhysicalSensor) Initialize() {
	s.generateID()
	s.serialCode = [SerialCodeLength]byte{StartByte1, StartByte2, 4, s.idx, 0x01, s.device, s.port}
	s.driver = &drivers.SerialState
}

type sensorReading struct {
	value  float32
	sensor Sensor
}

// GetValue returns the stored sensor value.
func (s *sensorReading) GetValue() float32 {
	return s.value
}

// GetSensor returns the stored sensor value.
func (s *sensorReading) GetSensor() Sensor {
	return s.sensor
}

// Sensors is a map designed to store a Sensor at any given index.
type sensorPackage struct {
	manifest     map[int]Sensor
	currentIndex int
	length       int
}

func (s *sensorPackage) Initialize() {
	s.manifest = make(map[int]Sensor)
}

// GetSensor a particular sensor at a given index from the manifest.
func (s *sensorPackage) GetSensor() Sensor {
	return s.manifest[s.currentIndex-1]
}

// SetSensor a particular Sensor at a given index in the manifest.
func (s *sensorPackage) SetSensor(sensor Sensor) {
	s.Next()
	s.manifest[s.currentIndex-1] = sensor
}

func (s *sensorPackage) WriteLength(length int) {
	s.length = length
}

// Next moves the iteration forward.
func (s *sensorPackage) Next() bool {
	s.currentIndex++
	if s.currentIndex < s.length {
		return true
	}
	s.currentIndex = 0
	return false
}
