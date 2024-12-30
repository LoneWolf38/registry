package main

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	OPCODE_LEN = 2
)

// TCP Data is represented as follows in string format for version 1
// Version:OpCode:Len:Data
// Example: 1:RE:12:192.168.1.0
type RegEvent struct {
	Version byte // Version number cannot be greater than 8Bit
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
	Event  []byte // Stores the data of the events kind of like headers with metadata
}

// Bytes representation of the RegEvents
func (r *RegEvent) Bytes() []byte {
	return nil
}

func (r *RegEvent) ToString() string {
	return ""
}

func (r *RegEvent) Marshal() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%d%s%d%s", r.Version, r.OpCode, r.Len, r.Event)
	return b.Bytes(), nil
}

func (r *RegEvent) Unmarshal(data []byte) error {
	// Read the byte array and extract the information from the data
	fmt.Printf("%b", data)
	r.Version = data[0]
	v, err := strconv.Atoi(string(data[0]))
	if err != nil {
		return fmt.Errorf("cannot parse version, Reason: %s", err.Error())
	}
	r.Version = byte(v)
	r.OpCode = data[1:3]
	fmt.Printf("%b", r.OpCode)
	l, err := strconv.Atoi(string(data[3:4]))
	if err != nil {
		return fmt.Errorf("cannot parse len, Reason: %s", err.Error())
	}
	r.Len = byte(l)

	r.Event = data[:r.Len]
	return nil
}
