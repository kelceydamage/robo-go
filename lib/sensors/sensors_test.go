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
	"testing"
)

func TestConst(t *testing.T) {
	var c byte
	c = StartByte1
	if c != 0xff {
		t.Errorf("Start byte for serial code is incorrect. Expected: 0xff, Got: %x", c)
	}

	c = StartByte2
	if c != 0x55 {
		t.Errorf("Confirmation byte for serial code is incorrect. Expected: 0x55, Got: %x", c)
	}

	c = CommRecv
	if c != 8 {
		t.Errorf("Recv byte length for serial code is incorrect. Expected: 0x08, Got: %x", c)
	}

	c = CommSend
	if c != 7 {
		t.Errorf("Send byte length for serial code is incorrect. Expected: 0x07, Got: %x", c)
	}
}
