package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventMarshal(t *testing.T) {
	// Actual data
	// [49 49 48 48 48 49 58 48 49 58 49 49 58 104 101 108 108 111 32 119 111 114 108 100 10]
	actual_d := []byte(strings.TrimSuffix("10111hello world", "\r\n"))
	le := uint8(len([]byte("hello world")))
	r := RegEvent{
		Version: byte(1),
		OpCode:  []byte("01"),
		Len:     le,
		Event:   []byte("hello world"),
	}
	d, err := r.Marshal()
	assert.Equal(t, err, nil, "they should be equal")
	assert.Equal(t, d, actual_d)
}

func TestEventUnMarshal(t *testing.T) {
	// Actual data
	actual_d := []byte("10111hello world")
	le := uint8(len([]byte("hello world")))
	r := &RegEvent{
		Version: byte(1),
		OpCode:  []byte("01"),
		Len:     le,
		Event:   []byte("hello world"),
	}
	result_r := &RegEvent{}
	err := result_r.Unmarshal(actual_d)
	assert.Equal(t, err, nil)
	assert.Equal(t, r, result_r)
}
