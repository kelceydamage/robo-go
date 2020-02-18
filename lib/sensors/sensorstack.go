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
	"errors"
)

// Dummy return object for errors to avoid memory panic for non-existant named return.
var errorReading SensorReading

// SensorStack is a LIFO storage object to ensure the controller get
// the most recent sensor data when checking state.
type SensorStack []*SensorReading

// Push appends an object onto the stack.
func (q *SensorStack) Push(n *SensorReading) {
	*q = append(*q, n)
}

// Pop returns the top item and removes it from the stack.
func (q *SensorStack) Pop() (top *SensorReading, err error) {
	top = &errorReading
	x := q.Len() - 1
	if x < 0 {
		err = errors.New("Empty Stack")
		return
	}
	top = (*q)[x]
	*q = (*q)[:x]
	return
}

// Len returns the length of the stack.
func (q *SensorStack) Len() int {
	return len(*q)
}

// Peek returns the top item from the stack without removing it.
func (q *SensorStack) Peek() (top *SensorReading, err error) {
	top = &errorReading
	x := q.Len() - 1
	if x < 0 {
		err = errors.New("Empty Stack")
		return
	}
	top = (*q)[x]
	return
}

// Bottom returns the bottom item from the stack without removing it.
func (q *SensorStack) Bottom() (top *SensorReading, err error) {
	top = &errorReading
	if q.Len() == 0 {
		err = errors.New("Empty Stack")
		return
	}
	top = (*q)[0]
	return
}

// Slice returns a slice from the stack without removing it.
func (q *SensorStack) Slice(start int, end int) (top []*SensorReading, err error) {
	top = []*SensorReading{&errorReading}
	if q.Len() == 0 {
		err = errors.New("Empty Stack")
		return
	}
	if end > q.Len()-1 || start > q.Len()-1 {
		err = errors.New("Slice Range Out Of Bounds")
		return
	}
	top = (*q)[start:end]
	return
}
