package main

import (
	"bytes"
	"fmt"
	"strconv"
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
type RegEvent struct {
	// Version stores the version of the protocol
	// V1 is for text based TCP protocol
	// V2 will have binary based TCP protocol
	Version byte
	// Opcode can be the
	//
	// 1. Register(RE)(01) - takes in the hostname, ip, mac-address, uptime, type of os, timestamp when its sending
	//
	// 2. Status(ST)(11) - takes in the hostname, ip, mac-address, uptime,type of os, timestamp when its sending
	//
	// 3. DeRegister(DR)(10) - takes in the hostname, ip, mac-address, uptime, type of os, timestamp when its sending
	//
	// 4. Shutdown(SH)(00) - takes in the hostname, ip, mac-address, uptime, type of os, timestamp when its sending
	OpCode []byte // Opcode takes 2bytes
	Len    byte   // Stores the len of the event
	// Events are completed by a ';' at the end
	Event []byte // Stores the data of the events kind of like headers with metadata
}

func (r *RegEvent) Marshal() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%d%s%d%s;", r.Version, r.OpCode, r.Len, r.Event)
	return b.Bytes(), nil
}

// Lets assume the last byte of the data contains the ;
func (r *RegEvent) Unmarshal(data []byte) error {
	// Read the byte array and extract the information from the data
	r.Version = data[0]
	v, err := strconv.Atoi(string(data[0]))
	if err != nil {
		return fmt.Errorf("cannot parse version, Reason: %s", err.Error())
	}
	r.Version = byte(v)
	r.OpCode = data[1:3]
	l, err := strconv.Atoi(string(data[3:5]))
	if err != nil {
		return fmt.Errorf("cannot parse len, Reason: %s", err.Error())
	}
	r.Len = byte(l)

	data_len := r.Len + 5

	r.Event = data[5:data_len]
	return nil
}
