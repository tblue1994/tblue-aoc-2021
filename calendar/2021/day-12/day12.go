package main

import (
	"strings"
	"tblue-aoc-2021/utils/files"
	"unicode"
)

type Vertex struct {
	// Key is the unique identifier of the vertex
	Key string
	// Vertices will describe vertices connected to this one
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[string]*Vertex

	MultipleVisit bool
}

// We then create a constructor function for the Vertex
func NewVertex(key string, multipleVisit bool) *Vertex {
	return &Vertex{
		Key:           key,
		MultipleVisit: multipleVisit,
		Vertices:      map[string]*Vertex{},
	}
}

type Graph struct {
	// Vertices describes all vertices contained in the graph
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[string]*Vertex
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

func (g *Graph) AddVertex(key string, multipleVisit bool) {
	v := NewVertex(key, multipleVisit)
	g.Vertices[key] = v
}

// The AddEdge method adds an edge between two vertices in the graph
func (g *Graph) AddEdge(k1, k2 string) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	// return an error if one of the vertices doesn't exist
	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	// do nothing if the vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	// Add a directed edge between v1 and v2
	// If the graph is undirected, add a corresponding
	// edge back from v2 to v1, effectively making the
	// edge between v1 and v2 bidirectional
	v1.Vertices[v2.Key] = v2
	if v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}

	// Add the vertices to the graph's vertex map
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

func main() {
	input := files.ReadFile(12, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	g := parseInput(input)
	return countRoutes(g)
}

func solvePart2(input []string) int {
	g := parseInput(input)
	return countRoutes2(g)
}

func parseInput(input []string) *Graph {
	g := NewUndirectedGraph()
	for _, val := range input {
		nodes := strings.Split(val, "-")
		for _, node := range nodes {
			_, found := g.Vertices[node]
			if !found {
				g.AddVertex(node, IsUpper(node))
			}
		}
		g.AddEdge(nodes[0], nodes[1])
	}
	return g
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func countRoutes(g *Graph) int {
	count := 0
	start := g.Vertices["start"]
	visited := map[string]bool{}
	count = DFS(g, start, visited)

	return count
}

func countRoutes2(g *Graph) int {
	count := 0
	start := g.Vertices["start"]
	visited := map[string]int{}
	count = DFS2(g, start, visited)

	return count
}

func DFS(g *Graph, c *Vertex, visited map[string]bool) int {
	if c.Key == "end" {
		// this counts as one full path
		return 1
	}
	currentVisited := map[string]bool{}

	for k, v := range visited {
		currentVisited[k] = v
	}

	if !c.MultipleVisit {
		currentVisited[c.Key] = true
	}

	count := 0

	for _, v := range c.Vertices {
		if currentVisited[v.Key] {
			continue
		}
		count += DFS(g, v, currentVisited)
	}
	return count
}

func DFS2(g *Graph, c *Vertex, visited map[string]int) int {
	if c.Key == "end" {
		// this counts as one full path
		return 1
	}
	currentVisited := map[string]int{}

	for k, v := range visited {
		currentVisited[k] = v
	}

	if !c.MultipleVisit {
		currentVisited[c.Key] += 1
	}

	count := 0

	for _, v := range c.Vertices {
		if !canVisitNode(v, currentVisited) {
			continue
		}
		count += DFS2(g, v, currentVisited)
	}
	return count
}

func canVisitNode(c *Vertex, visited map[string]int) bool {
	value := visited[c.Key]
	canVisitSmallTwice := true
	for _, v := range visited {
		if v >= 2 {
			canVisitSmallTwice = false
			break
		}
	}
	return c.Key != "start" && (value == 0 || canVisitSmallTwice)
}
