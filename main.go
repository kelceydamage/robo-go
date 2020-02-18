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
var sensorFeed = make(chan sensors.SensorReading, 512)

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
		fmt.Printf("Receiving: %v\n", result.Value)
		counter++
		fmt.Printf("Receiving Count: %v Channel Occupancy: %v\n", counter, len(sensorFeed))
	}

	// Will fail unless BufferSensors is set to infinite loop
	wg.Wait()

	//fmt.Println("main finished")
}
