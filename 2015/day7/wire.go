package day7

import "strings"

type wire struct {
	name   string
	signal signal
}

var wires []*wire

func createWire(name string, signal signal) *wire {
	return &wire{name: name, signal: signal}
}

func getWire(name string) *wire {
	for _, wire := range wires {
		if strings.EqualFold(wire.name, name) {
			return wire
		}
	}

	newWire := createWire(name, signal{value: 0, valid: false})
	wires = append(wires, newWire)
	return newWire
}

func resetWires() {
	wires = nil
}
