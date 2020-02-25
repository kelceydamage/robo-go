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
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/kelceydamage/robo-go/lib/drivers"
)

// PackageSensors constructor.
func PackageSensors(numberOfSensors int) (s Sensors) {
	s.manifest = make(map[int]Sensor)
	return s
}

// BufferSensors is a go routine that continually loops through the Sensors and
// writes their data to a channel.
func BufferSensors(sensorPackage Sensors, c drivers.Comm, channel chan SensorReading) {
	// Still need some safety around threading this.
	var wg sync.WaitGroup
	wg.Add(1)

	go bufferSensorsRoutine(&wg, sensorPackage, c, channel)

	// Will fail unless BufferSensors is set to infinite loop
	wg.Wait()
}

func bufferSensorsRoutine(wg *sync.WaitGroup, sensorPackage Sensors, c drivers.Comm, channel chan SensorReading) {
	defer wg.Done()
	for {
		tempBuff := [12]byte{11: 0x00}
		for _, sensor := range sensorPackage.manifest {
			fmt.Printf("Sending: %v\n", sensor.Serialized)

			_, err := c.Write(sensor.Serialized)
			if err != nil {
				log.Fatalf("port.Read: %v", err)
				break
			}
			_, err = c.Read(&tempBuff)
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
