package main

import (
	"time"
)

/* SAMPLE DATA
[{
    "setup": "stappenrobot",
    "timestamp": "2017-11-27T13:32:45.940Z",
    "ttl": "1Y",
    "data": [{
        "sensor": "sensor X",
        "data": 1.1,
    },
    {
        "sensor": "sensor Y",
        "data": 1.1,
    }]
}]
*/

// APIData is the data sent by an API call
type APIData struct {
	Setup     string        `json:"setup"`
	Timestamp time.Time     `json:"timestamp"`
	TTL       time.Duration `json:"ttl"`
	Data      []SensorData  `json:"data"`
}

// SensorData is the data of one sensor sent an an APIData
type SensorData struct {
	Sensor string  `json:"sensor"`
	Data   float64 `json:"data"`
}
