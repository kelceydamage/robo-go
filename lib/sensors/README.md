# Package Sensors

### Object Hierarchy

#### Sensors
A struct containing a manifest map of all the registered **Sensor** objects.

#### Sensor
A struct containing the serial code to communicate with the physical sensor.

#### SensorReading
A struct containing the response from a physical sensor.

#### SensorStack
A LIFO stack for intended for storing **SensorReading** objects pulled from the channel used in the **BufferSensors** goroutine.

### Helper Functions

#### PackageSensors
Creates a object of **Sensors** type with allocated memory for [n] sensors.

#### BufferSensors
Launces a goroutine that loops over all the registered **Sensor** objects in **Sensors** and queries them for a reading. It then submits those readings into a channel.

### Sensor Types
These are mainly pre established aliases for supported physical sensors, and map to **Sensor** objects.

#### Ultrasonic