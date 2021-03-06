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

// Ultrasonic is a predefined alias object for a sensor.
var Ultrasonic PhysicalSensor

// LoadSensorTypes makes all the known sensor types available and setup.
func LoadSensorTypes() {
	Ultrasonic.port = UltrasonicPort
	Ultrasonic.device = UltrasonicDevice
}
