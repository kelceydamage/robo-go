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
