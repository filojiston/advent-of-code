// https://adventofcode.com/2015/day/9
// solution of advent of code 2015, day9

package day9

import (
	"fmt"
	"math"

	"github.com/filojiston/advent-of-code/2015/util"
)

func getMinMaxPathLengths() (int, int) {
	graph := createGraphFromDistances()
	minDistance := math.MaxInt
	maxDistance := math.MinInt

	for node := range graph.nodes {
		shortestPath := findShortestPathLengthForNode(graph, node)
		longestPath := findLongestPathLengthForNode(graph, node)
		if shortestPath < minDistance {
			minDistance = shortestPath
		}
		if longestPath > maxDistance {
			maxDistance = longestPath
		}
	}
	return minDistance, maxDistance
}

func createGraphFromDistances() *graph {
	lines := util.ReadInputFile("input.txt")
	graph := newGraph()
	var from, to string
	var distance int

	for _, line := range lines {
		fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance)
		graph.addNode(from)
		graph.addNode(to)
		graph.addEdge(from, to, distance)
	}

	return graph
}

func findShortestPathLengthForNode(graph *graph, name string) int {
	shortestPathLength := 0
	visited := make(map[string]bool)
	visited[name] = true
	currentNode := name
	for len(visited) < len(graph.nodes) {
		min := math.MaxInt
		nextNode := ""
		for node, distance := range graph.edges[currentNode] {
			if distance < min && !visited[node] {
				min = distance
				nextNode = node
			}
		}
		visited[nextNode] = true
		shortestPathLength += min
		currentNode = nextNode
	}
	return shortestPathLength
}

func findLongestPathLengthForNode(graph *graph, name string) int {
	longestPathLength := 0
	visited := make(map[string]bool)
	visited[name] = true
	currentNode := name
	for len(visited) < len(graph.nodes) {
		max := math.MinInt
		nextNode := ""
		for node, distance := range graph.edges[currentNode] {
			if distance > max && !visited[node] {
				max = distance
				nextNode = node
			}
		}
		visited[nextNode] = true
		longestPathLength += max
		currentNode = nextNode
	}
	return longestPathLength
}
