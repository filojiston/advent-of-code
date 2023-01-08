package day9

import "log"

type graph struct {
	nodes map[string]struct{}
	edges map[string]map[string]int
}

func newGraph() *graph {
	return &graph{
		nodes: make(map[string]struct{}),
		edges: make(map[string]map[string]int),
	}
}

func (g *graph) addNode(id string) bool {
	if _, ok := g.nodes[id]; ok {
		return false
	}

	g.nodes[id] = struct{}{}
	return true
}

func (g *graph) addEdge(u, v string, weight int) {
	if _, ok := g.nodes[u]; !ok {
		log.Fatalf("[addEdge] node %s not found on graph.", u)
		return
	}
	if _, ok := g.nodes[v]; !ok {
		log.Fatalf("[addEdge] node %s not found on graph.", v)
		return
	}

	if _, ok := g.edges[u]; !ok {
		g.edges[u] = make(map[string]int)
	}

	g.edges[u][v] = weight
	if _, ok := g.edges[v]; !ok {
		g.edges[v] = make(map[string]int)
	}

	g.edges[v][u] = weight
}
