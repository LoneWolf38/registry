package registry

import "time"

// This is the record of each event sent by an individual host
// The list of type of events are:
// 1. Hostname
// 2. IPV4 Address
// 3. MacAddress
// 4. OSType
// 5. TimeStamp when it was sent
// 6. UpTime of the server
type Record struct {
	hostname string
	ipAddr   string
	macAddr  string
	osType   string
	ts       time.Time
	upTime   time.Time
}
