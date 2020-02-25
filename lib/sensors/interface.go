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

// SensorReader is the interface for reading from a sensor.
type serialReader interface {
	Read() SensorReading
}

type portReader interface {
	GetPort() byte
}

type deviceReader interface {
	GetDevice() byte
}

type idReader interface {
	GetID() byte
}

type valueReader interface {
	GetValue() float32
}

type serialCodeReader interface {
	GetSerialCode() [SerialCodeLength]byte
}

type sensorReader interface {
	GetSensor() Sensor
}

type sensorWriter interface {
	SetSensor(int Sensor)
}

type nextIterator interface {
	Next() bool
}

type lengthWriter interface {
	WriteLength(int)
}

type initializer interface {
	Initialize()
}

// Sensor is the interface for physical sensors plugged into the board.
type Sensor interface {
	serialReader
	portReader
	deviceReader
	idReader
	serialCodeReader
	initializer
}

// SensorReading is the interface for anyobject returned by a physical sensor.
type SensorReading interface {
	valueReader
	sensorReader
}

// SensorPackage is a map of registered sensor objects.
type SensorPackage interface {
	initializer
	sensorReader
	sensorWriter
	nextIterator
	lengthWriter
}
