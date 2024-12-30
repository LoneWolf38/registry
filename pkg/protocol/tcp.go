package protocol

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

// RegEventServer is the struct which stores the connections
type RegEventServer struct {
	listner  net.Listener
	conn     net.Conn
	protocol string
	port     uint16
}

type RegEventClient struct {
	conn     net.Conn
	protocol string
	port     uint16
}

func NewRegEventClient(proto string, port uint16) (RegEventClient, error) {
	r := RegEventClient{}
	c, err := net.Dial(proto, fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return r, err
	}
	r.conn = c
	return r, nil
}

func (rc *RegEventClient) Close() error {
	return rc.conn.Close()
}

func (rc *RegEventClient) Send(regEvent RegEvent) error {
	d, err := regEvent.Marshal()
	if err != nil {
		return err
	}
	_, err = rc.conn.Write(d)
	if err != nil {
		return err
	}
	return nil
}

func NewEventServer(proto string, port uint16) (RegEventServer, error) {
	r := RegEventServer{}
	l, err := net.Listen(proto, fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return r, err
	}
	r.listner = l
	r.port = port
	r.protocol = proto
	return r, nil
}

func (r RegEventServer) Start() {
	log.Println("started the event server")
	regEvents := make(chan RegEvent)
	go func() {
		for {
			conn, err := r.listner.Accept()
			if err != nil {
				log.Printf("E! cannot accept connection, Reason: %s", err.Error())
			}
			r.conn = conn
			go func() {
				err := readConn(r, regEvents)
				if err != nil {
					log.Printf("E! cannot parse data: Reason: %s", err)
				}
			}()
		}
	}()
	for regEvent := range regEvents {
		if bytes.Equal(regEvent.opCode, REG_EVENT) {
			fmt.Printf("REG EVENT %s\n", string(regEvent.event))
		} else if bytes.Equal(regEvent.opCode, DEREGISTER_EVENT) {
			fmt.Printf("DEREG EVENT %s\n", string(regEvent.event))
		} else if bytes.Equal(regEvent.opCode, STATUS_EVENT) {
			fmt.Printf("STATUS EVENTS %s\n", string(regEvent.event))
		} else if bytes.Equal(regEvent.opCode, SHUTDOWN_EVENT) {
			fmt.Printf("SHUTDOWN EVENT %s\n", string(regEvent.event))
		} else {
			fmt.Printf("ERROR EVENT\n")
		}
	}
}

func readConn(r RegEventServer, regEvents chan RegEvent) error {
	defer r.conn.Close()
	// read upto 1024 bytes
	data := make([]byte, 1024)
	regEvent := RegEvent{}
	buf := bytes.Buffer{}
	for {
		_, err := r.conn.Read(data)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		for _, d := range data {
			if d == ';' {
				if err := regEvent.Unmarshal(buf.Bytes()); err != nil {
					return err
				}
				regEvents <- regEvent
			}
			buf.WriteByte(d)
		}
	}
}
