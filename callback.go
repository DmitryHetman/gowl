
package gowl

import (
	"bytes"
)

type Callback struct {
//	*WlObject
	id int32
	listeners map[int16][]chan interface{}
	events []func (c *Callback, msg []byte)
}

//// Requests
//// Events
func (c *Callback) HandleEvent(opcode int16, msg []byte) {
	if c.events[opcode] != nil {
		c.events[opcode](c, msg)
	}
}

type CallbackDone struct {
	serial uint32
}

func (c *Callback) AddDoneListener(channel chan interface{}) {
	c.addListener(0, channel)
}

func callback_done(c *Callback, msg []byte) {
	printEvent("done", msg)
	var data CallbackDone
	buf := bytes.NewBuffer(msg)

	serial,err := readUint32(buf)
	if err != nil {
		// XXX Error handling
	}
	data.serial = serial

	for _,channel := range c.listeners[0] {
		channel <- data
	}
}

func NewCallback() (c *Callback) {
	c = new(Callback)
	c.listeners = make(map[int16][]chan interface{}, 0)

	c.events = append(c.events, callback_done)
	return
}