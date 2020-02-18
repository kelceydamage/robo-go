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

// Copyright (c) 2020 Author Name. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package sensors

import (
	"encoding/binary"
	"fmt"
	"log"
	"sync"
	"time"
	"unsafe"
)

// SensorReading is a tagged value object.
type SensorReading struct {
	Port   byte
	Device byte
	Idx    byte
	Value  float32
}

// Sensor is the representation of a physical sensor on the controller.
// It contains the serial code to retrieve data from the ophysical sensor, along
// with the device, port, and id.
type Sensor struct {
	port       byte
	device     byte
	idx        byte
	Serialized []byte
}

func (s *Sensor) getReading(c *comm) SensorReading {
	var reading SensorReading
	reading.Port = s.port
	reading.Device = s.device
	reading.Idx = s.idx
	reading.Value = s.asFloat((*c).Result(CommRecv)[4:])
	return reading
}

func (s *Sensor) asFloat(bytes []byte) float32 {
	binrep := binary.LittleEndian.Uint32(bytes)
	floatrep := *(*float32)(unsafe.Pointer(&binrep))
	fmt.Printf("Converting: %v, %v -> %v\n", bytes, binrep, floatrep)
	return floatrep
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
func BufferSensors(wg *sync.WaitGroup, sensorPackage Sensors, c comm, channel chan SensorReading) {
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
			_, err = c.Read(tempBuff)
			if err != nil {
				log.Fatalf("port.Read: %v", err)
				break
			} else {
				channel <- sensor.getReading(&c)
			}
			time.Sleep(20 * time.Millisecond)
		}
	}
}

type comm interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Result(int) []byte
}
