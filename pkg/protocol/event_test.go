package protocol

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventMarshal(t *testing.T) {
	// Actual data
	actual_d := []byte(strings.TrimSuffix("10111hello world;", "\r\n"))
	le := uint8(len([]byte("hello world")))
	r := RegEvent{
		version: byte(1),
		opCode:  []byte("01"),
		len:     le,
		event:   []byte("hello world"),
	}
	d, err := r.Marshal()
	assert.Equal(t, err, nil, "they should be equal")
	assert.Equal(t, d, actual_d)
}

func TestEventUnMarshal(t *testing.T) {
	// Actual data
	actual_d := []byte(strings.TrimSuffix("10111hello world;", "\r\n"))
	le := uint8(len([]byte("hello world")))
	r := &RegEvent{
		version: byte(1),
		opCode:  []byte("01"),
		len:     le,
		event:   []byte("hello world"),
	}
	result_r := &RegEvent{}
	err := result_r.Unmarshal(actual_d)
	assert.Equal(t, err, nil)
	assert.Equal(t, r, result_r)
}
