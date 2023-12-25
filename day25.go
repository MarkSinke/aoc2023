package aoc2023

import (
	"math/rand"
	"slices"
	"strings"

	"github.com/fzipp/astar"
	"golang.org/x/exp/maps"
)

type NodeMap map[string][]string

func (m NodeMap) Neighbours(name string) []string {
	return m[name]
}

func insertNames(nodes NodeMap, s0, s1 string) {
	list, found := nodes[s0]
	if !found {
		list = []string{s1}
	} else {
		list = append(list, s1)
	}
	nodes[s0] = list
}

func parseNodeLine(nodes NodeMap, str string) {
	parts := strings.Split(str, " ")

	name0, _ := strings.CutSuffix(parts[0], ":")
	for _, part := range parts[1:] {
		insertNames(nodes, name0, part)
		insertNames(nodes, part, name0)
	}
}

func ReadComponentGraph(path string) NodeMap {
	lines := ReadFile(path)

	nodes := NodeMap{}
	for _, line := range lines {
		parseNodeLine(nodes, line)
	}

	return nodes
}

func cost(name0, name1 string) float64 {
	return 1
}

func estimate(name0, name1 string) float64 {
	return 0
}

func randomNode(nodeNames []string) string {
	i := rand.Intn(len(nodeNames))
	return nodeNames[i]
}

func removeEdge(nodes NodeMap, e Edge) {
	list := nodes[e.from]
	slices.DeleteFunc(list, func(str string) bool { return str == e.to })
	nodes[e.from] = list

	list = nodes[e.to]
	slices.DeleteFunc(list, func(str string) bool { return str == e.from })
	nodes[e.to] = list
}

func PartitionGraph(nodes NodeMap) (string, string) {
	// pick any two nodes and find the path between them
	// repeat for many pairs
	// the edges most traveled will be the bridges that we need to remove
	edge0, found := findMaxEdge(nodes)
	if !found {
		panic("should find edge0")
	}
	removeEdge(nodes, edge0)
	edge1, found := findMaxEdge(nodes)
	if !found {
		panic("should find edge1")
	}
	removeEdge(nodes, edge1)
	edge2, found := findMaxEdge(nodes)
	if !found {
		panic("should find edge2")
	}
	removeEdge(nodes, edge2)
	edge3, found := findMaxEdge(nodes)
	if found {
		panic("graph should be two components now")
	}

	return edge3.from, edge3.to
}

type Edge struct {
	from string
	to   string
}

// Return either
// (a) edge with max traversal count,  or
// (b) "edge" where from and to are in separate components of the graph
func findMaxEdge(nodes NodeMap) (edge Edge, found bool) {
	nodeNames := maps.Keys(nodes)
	counts := map[Edge]int{}
	for i := 0; i < 1000; i++ {
		from := randomNode(nodeNames)
		to := randomNode(nodeNames)

		path := astar.FindPath(nodes, from, to, cost, estimate)
		if path == nil {
			// graph is partioned between those nodes
			return Edge{from, to}, false
		}
		for i := 1; i < len(path); i++ {
			from := path[i-1]
			to := path[i]
			if from > to {
				from, to = to, from
			}
			elem := Edge{from, to}
			count, found := counts[elem]
			if !found {
				count = 1
			} else {
				count++
			}
			counts[elem] = count
		}
	}

	maxCount := 0
	var maxEdge Edge
	for edge, count := range counts {
		if count > maxCount {
			maxCount = count
			maxEdge = edge
		}
	}
	return maxEdge, true
}

func CountComponentSize(nodes NodeMap, nodeName string) int {
	visited := map[string]bool{nodeName: true}
	queue := []string{nodeName}
	count := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		count++
		visited[cur] = true

		for _, name := range nodes[cur] {
			if !visited[name] && !slices.Contains(queue, name) {
				queue = append(queue, name)
			}
		}
	}
	return count
}
