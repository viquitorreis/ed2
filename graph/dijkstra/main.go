package main

import (
	"fmt"
	"math"
)

func main() {
	graph := NewGraph().AddVertex("A").AddVertex("B").AddVertex("C").AddVertex("D")
	graph.AddEdge("A", "B", 10)
	graph.AddEdge("A", "C", 20)
	graph.AddEdge("B", "D", 5)
	graph.AddEdge("C", "D", 2)
	res := graph.Dijkstra("A", "D")
	fmt.Println("Dijkstra result: ", res) // 15 [A B D] true
}

type IGraph interface {
	AddVertex(val string) *Graph
	AddEdge(from, to string, weight int) bool
	HasEdge(from, to string) bool
	GetVertex(val string) *Vertex
	GetNeighbors(n string) []*Edge
	Dijkstra(source, target string) Result
	PrintGraph()
}

type Graph struct {
	Vertices map[string]*Vertex
	Len      int
	EdgesLen int
}

type Vertex struct {
	val   string
	edges map[string]*Edge
}

type Edge struct {
	weight int
	dest   *Vertex
}

type Result struct {
	Distance int
	Path     []string
	Found    bool
}

type DistanceTable struct {
	NodeKey            string
	ShortestPathWeight int
	PrevNodeKey        string
}

func NewGraph() IGraph {
	return &Graph{
		Vertices: make(map[string]*Vertex),
		Len:      0,
		EdgesLen: 0,
	}
}
func (g *Graph) GetVertex(val string) *Vertex {
	if vertex, exists := g.Vertices[val]; exists {
		return vertex
	}

	return nil
}

func (g *Graph) AddVertex(val string) *Graph {
	vertex := g.GetVertex(val)
	if vertex != nil {
		return g
	}

	vertex = &Vertex{
		val:   val,
		edges: make(map[string]*Edge),
	}

	g.Vertices[val] = vertex
	g.Len++

	return g
}

func (g *Graph) AddEdge(from, to string, weight int) bool {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	if fromVertex == nil || toVertex == nil {
		return false
	}

	var toEdge *Edge
	if oldEdge, exists := fromVertex.edges[to]; exists {
		toEdge = &Edge{
			weight: weight,
			dest:   oldEdge.dest,
		}
	} else {
		toEdge = &Edge{
			weight: weight,
			dest:   toVertex,
		}
	}

	fromEdge := &Edge{
		weight: weight,
		dest:   fromVertex,
	}

	fromVertex.edges[to] = toEdge
	toVertex.edges[from] = fromEdge
	g.EdgesLen++

	return true
}

func (g *Graph) GetNeighbors(n string) []*Edge {
	vertex := g.GetVertex(n)
	if vertex == nil {
		return []*Edge{}
	}

	var neighbors []*Edge
	for _, edge := range vertex.edges {
		neighbors = append(neighbors, edge)
	}

	return neighbors
}

func (g *Graph) GetLen() int {
	return g.Len
}

func (g *Graph) HasEdge(from, to string) bool {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	if fromVertex == nil || toVertex == nil {
		return false
	}

	if _, exists := fromVertex.edges[to]; exists {
		return true
	}

	return false
}

func (g *Graph) Dijkstra(source, target string) Result {
	// 1 - inicializa a tabela. Node, Shortest Node weight, Previous Node
	distancesTable := g.initDistancesTable()
	distancesTable[source].ShortestPathWeight = 0
	unvisited := make(map[string]bool)
	for key := range g.Vertices {
		unvisited[key] = true
	}

	// 2 - Range nos não visitados
	for len(unvisited) > 0 {
		// 3 - Escolher vertice não visitado com menor distancia
		minUnvisited := g.getMinFromUnvisited(distancesTable, unvisited)

		// 4 - Base case, se chegou no destino, para o loop
		if minUnvisited == target {
			break
		}

		// 5 - Se a distancia até o atual for infinita, nao tem solução
		if distancesTable[minUnvisited].ShortestPathWeight == math.MaxInt {
			break
		}

		// 6 - Atualiza distancia dos vizinhos, se for menor que a anterior, somando com o que já foi caminhado
		neighbors := g.GetNeighbors(minUnvisited)
		for _, edge := range neighbors {
			neighborKey := edge.dest.val
			newDistance := distancesTable[minUnvisited].ShortestPathWeight + edge.weight

			if newDistance < distancesTable[neighborKey].ShortestPathWeight {
				// atualiza o menor peso até chegar no VIZINHO e não no atual...
				distancesTable[neighborKey].ShortestPathWeight = newDistance
				distancesTable[neighborKey].PrevNodeKey = minUnvisited
			}
		}

		delete(unvisited, minUnvisited)
	}

	// 7 - Após o loop, retorna o resultado
	if distancesTable[target].ShortestPathWeight == math.MaxInt {
		return Result{Distance: -1, Path: []string{}, Found: false}
	}

	path := g.reconstructPath(distancesTable, target)
	return Result{
		Distance: distancesTable[target].ShortestPathWeight,
		Path:     path,
		Found:    true,
	}
}

func (g *Graph) getMinFromUnvisited(distancesTable map[string]*DistanceTable, unvisited map[string]bool) string {
	minKey := ""
	minDistance := math.MaxInt

	for nodeKey := range unvisited {
		if distancesTable[nodeKey].ShortestPathWeight < minDistance {
			minDistance = distancesTable[nodeKey].ShortestPathWeight
			minKey = nodeKey
		}
	}

	return minKey
}

func (g *Graph) initDistancesTable() map[string]*DistanceTable {
	table := make(map[string]*DistanceTable, g.Len)

	for key := range g.Vertices {
		table[key] = &DistanceTable{
			NodeKey:            key,
			ShortestPathWeight: math.MaxInt,
			PrevNodeKey:        "",
		}
	}

	return table
}

func (g *Graph) reconstructPath(distancesTable map[string]*DistanceTable, target string) []string {
	path := []string{}
	current := target

	for current != "" {
		path = append([]string{current}, path...)
		current = distancesTable[current].PrevNodeKey
	}

	return path
}

func (g *Graph) PrintGraph() {
	for k, v := range g.Vertices {
		fmt.Printf("key is: %s - Val: %v", k, v)

		if edge, exists := v.edges[k]; exists {
			fmt.Printf("connection edge: %#v ", edge)
		}

		fmt.Printf("\n")
	}
}
