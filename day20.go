package aoc2023

import (
	"regexp"
	"strings"
)

type Gate interface {
	ConnectIn(in string)
	ReceivePulse(p Pulse) []Pulse
	Reset()
}

type BroadcasterGate struct {
	outputs []string
}

func (g *BroadcasterGate) ConnectIn(in string) {
}

func createPulses(from string, tos []string, value bool) []Pulse {
	pulses := []Pulse{}
	for _, to := range tos {
		pulses = append(pulses, Pulse{from, to, value})
	}
	return pulses
}

func (g *BroadcasterGate) ReceivePulse(p Pulse) []Pulse {
	return createPulses(p.to, g.outputs, p.value)
}

func (g *BroadcasterGate) Reset() {
}

type FlipFlopGate struct {
	outputs []string
	state   bool
}

func (g *FlipFlopGate) ConnectIn(in string) {
}

func (g *FlipFlopGate) ReceivePulse(p Pulse) []Pulse {
	if p.value {
		return []Pulse{}
	}

	g.state = !g.state
	return createPulses(p.to, g.outputs, g.state)
}

func (g *FlipFlopGate) Reset() {
	g.state = false
}

type ConjGate struct {
	outputs     []string
	inputStates map[string]bool
}

func (g *ConjGate) ConnectIn(in string) {
	if g.inputStates == nil {
		g.inputStates = map[string]bool{}
	}
	g.inputStates[in] = false
}

func (g *ConjGate) isAllHigh() bool {
	for _, state := range g.inputStates {
		if !state {
			return false
		}
	}
	return true
}

func (g *ConjGate) ReceivePulse(p Pulse) []Pulse {
	g.inputStates[p.from] = p.value
	return createPulses(p.to, g.outputs, !g.isAllHigh())
}

func (g *ConjGate) Reset() {
	for key := range g.inputStates {
		g.inputStates[key] = false
	}
}

type OutputGate struct {
}

func (g *OutputGate) ConnectIn(in string) {
}

func (g *OutputGate) ReceivePulse(p Pulse) []Pulse {
	return []Pulse{}
}

func (g *OutputGate) Reset() {
}

type GateGraph map[string]*Gate

func (g *GateGraph) Reset() {
	for _, gate := range *g {
		(*gate).Reset()
	}
}

var gateRegex = regexp.MustCompile(`([%&]?)(.*) -> (.*)`)

func parseLine(str string) (string, string, []string) {
	matches := gateRegex.FindStringSubmatch(str)

	gateType := matches[1]
	name := matches[2]
	outputs := strings.Split(strings.ReplaceAll(matches[3], " ", ""), ",")

	return gateType, name, outputs

}

func parseGate(graph GateGraph, str string) (string, Gate) {
	gateType, name, outputs := parseLine(str)

	switch gateType {
	case "":
		if name == "broadcaster" {
			return name, &BroadcasterGate{outputs: outputs}
		} else {
			panic("unknown gate: " + name)
		}
	case "%":
		return name, &FlipFlopGate{outputs: outputs}
	case "&":
		return name, &ConjGate{outputs: outputs}
	default:
		panic("unknown instruction: " + str)
	}
}

func ReadNetwork(path string) GateGraph {
	lines := ReadFile(path)

	graph := GateGraph{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		name, node := parseGate(graph, line)
		graph[name] = &node
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		_, name, outputs := parseLine(line)
		for _, output := range outputs {
			gate := graph[output]
			if gate == nil {
				var og Gate = &OutputGate{}
				gate = &og
				graph[output] = gate
			}

			(*gate).ConnectIn(name)
		}
	}

	return graph
}

type Pulse struct {
	from  string
	to    string
	value bool
}

func ExecNetwork(graph GateGraph) (lowCount int, highCount int) {
	pulses := []Pulse{{"button", "broadcaster", false}}

	for len(pulses) > 0 {
		pulse := pulses[0]
		pulses = pulses[1:]

		gate := graph[pulse.to]
		// fmt.Println(pulse.from, pulse.value, "->", pulse.to)
		newPulses := (*gate).ReceivePulse(pulse)

		pulses = append(pulses, newPulses...)

		if pulse.value {
			highCount++
		} else {
			lowCount++
		}
	}

	return
}

func ExecNetworkUntil(graph GateGraph, to string, value bool) bool {
	pulses := []Pulse{{"button", "broadcaster", false}}
	lowRx := false

	for len(pulses) > 0 {
		pulse := pulses[0]
		pulses = pulses[1:]

		if pulse.to == to && pulse.value == value {
			lowRx = true
		}

		gate := graph[pulse.to]
		// fmt.Println(pulse.from, pulse.value, "->", pulse.to)
		newPulses := (*gate).ReceivePulse(pulse)

		pulses = append(pulses, newPulses...)
	}

	return lowRx
}

func PressUntilLowRx(graph GateGraph, to string, value bool) int {
	for count := 1; ; count++ {
		if ExecNetworkUntil(graph, to, value) {
			return count
		}
	}
}
