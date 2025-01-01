package protocol

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/LoneWolf38/registry/pkg/registry"
)

var (
	REG_EVENT        = []byte("01")
	STATUS_EVENT     = []byte("11")
	DEREGISTER_EVENT = []byte("10")
	SHUTDOWN_EVENT   = []byte("00")
)

// TCP Data is represented as follows in string format for version 1
// Version:OpCode:Len:Data
// Example: 1:RE:12:192.168.1.0
type HeartBeat struct {
	// version stores the version of the protocol
	// V1 is for text based TCP protocol
	// V2 will have binary based TCP protocol
	version byte
	// Opcode can be the
	//
	// 1. Register(RE)(01) - takes in the hostname, ip, mac-address, uptime, type of os, timestamp when its sending
	//
	// 2. Status(ST)(11) - takes in the hostname, ip, mac-address, uptime,type of os, timestamp when its sending
	//
	// 3. DeRegister(DR)(10) - takes in the hostname, ip, mac-address, uptime, type of os, timestamp when its sending
	//
	// 4. Shutdown(SH)(00) - takes in the hostname, ip, mac-address, uptime, type of os, timestamp when its sending
	opCode []byte // Opcode takes 2bytes
	len    byte   // Stores the len of the event
	// Events are completed by a ';' at the end
	event []byte // Stores the data of the events kind of like headers with metadata
}

func NewHB(version uint8, opCode []byte, event []byte) HeartBeat {
	return HeartBeat{
		version: version,
		opCode:  opCode,
		len:     uint8(len(event)),
		event:   event,
	}
}

// marshalls the heartbeat to a stream of bytes
func (r *HeartBeat) Marshal() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%d%s%d%s;", r.version, r.opCode, r.len, r.event)
	return b.Bytes(), nil
}

// unmarshalls the stream of bytes in to HeartBeat
func (r *HeartBeat) Unmarshal(data []byte) error {
	// Read the byte array and extract the information from the data
	r.version = data[0]
	v, err := strconv.Atoi(string(data[0]))
	if err != nil {
		return fmt.Errorf("cannot parse version, Reason: %s", err.Error())
	}
	r.version = byte(v)
	r.opCode = data[1:3]
	l, err := strconv.Atoi(string(data[3:5]))
	if err != nil {
		return fmt.Errorf("cannot parse len, Reason: %s", err.Error())
	}
	r.len = byte(l)

	data_len := r.len + 5

	r.event = data[5:data_len]
	return nil
}

// parses the heartbeat to an registry record
// TODO: Implement this!!!
func (r *HeartBeat) Parse() (registry.Record, error) {
	var rr registry.Record
	return rr, nil
}
