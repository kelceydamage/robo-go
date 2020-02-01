package sensors

type Sensor struct {
	port byte
	device byte
	idx byte
	Serialized []byte
}

func (s *Sensor)generateId() {
	s.idx = ((s.port<<4) + s.device) & 0xff;
}

func (s *Sensor)Configure(device byte, port byte) {
	s.device = device
	s.port = port
	s.generateId()
	s.Serialized = []byte{StartByte1, StartByte2, 4, s.idx, 0x01, s.device, s.port}
}

type Sensors struct {
	manifest map[int]Sensor
}

// Getter function
func (s *Sensors)Get(id int) (sensor Sensor) {
	return s.manifest[id]
}

// Setter functiomn
func (s *Sensors)Set(id int, sensor Sensor) {
	s.manifest[id] = sensor
}