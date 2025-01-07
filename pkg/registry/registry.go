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

// hash creates a hash from the record which will be used as index in a go map[type]type
//
// the hash function works as follows
//
// Hostname will be turned to xor func hash
// ipAddr will be
func (r *Record) hash() []byte {
	return nil
}
