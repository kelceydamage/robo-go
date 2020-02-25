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
	"sync"
	"time"
)

// BufferSensors is a go routine that continually loops through the Sensors and
// writes their data to a channel.
func BufferSensors(sp SensorPackage, channel chan SensorReading) {
	// Still need some safety around threading this.
	var wg sync.WaitGroup
	wg.Add(1)

	go bufferSensorsRoutine(&wg, sp, channel)

	// Will fail unless BufferSensors is set to infinite loop
	wg.Wait()
}

func bufferSensorsRoutine(wg *sync.WaitGroup, sp SensorPackage, channel chan SensorReading) {
	defer wg.Done()
	for {
		for sp.Next() {
			sensor := sp.GetSensor()
			fmt.Printf("Sending: %v\n", sensor.GetSerialCode())

			channel <- sensor.Read()

			time.Sleep(20 * time.Millisecond)
		}
	}
}
