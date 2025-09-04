package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// graph := NewGraph()
	// graph.AddVertex(10).AddVertex(30).AddEdge(10, 30, 5)
	// graph.AddVertex(20).AddVertex(40).AddEdge(20, 40, 15)
	// graph.AddEdge(20, 10, 10)
	// graph.AddEdge(30, 40, 20)
	// graph.AddVertex(1).AddEdge(30, 1, 25)
	// graph.AddVertex(2).AddEdge(30, 2, 30)
	// graph.AddVertex(3).AddEdge(30, 3, 35)
	// graph.AddVertex(4).AddEdge(30, 4, 40)

	// fmt.Println(graph.GetNeighbors(20))
	// graph.TraverseGraphSimple(20)
	// graph.PrintGraph()

	// graph.BacktrackColoring(3)
	// graph.GreedyColoring(3)
	// graph.GreedyColoringByDegree(3)

	// Criando triangulo (3 vertices todos conectados)
	// graph.AddVertex(1).AddVertex(2).AddVertex(3)
	// graph.AddEdge(1, 2, 1) // 1 ↔ 2
	// graph.AddEdge(2, 3, 1) // 2 ↔ 3
	// graph.AddEdge(1, 3, 1) // 1 ↔ 3

	// tentar colorir com apenas 2 cores
	// fmt.Println(graph.BacktrackColoring(2))
	// graph.GreedyColoring(2)
	// fmt.Println(GenerateGraph(20, mid))

	testSuite := &TestSuite{}

	// parametros de teste
	vertices := []int{5, 10, 20, 50, 100}
	densities := []graphDensity{sparse, mid, dense}
	algorithms := []algorithm{backtrack, greedy, greedyByDegree}
	repetitions := 5

	for _, v := range vertices {
		for _, d := range densities {
			for _, alg := range algorithms {
				results := runMultipleTests(v, d, alg, repetitions)

				for _, result := range results {
					testSuite.AddResult(result)
				}

				printAndWriteResults(results)
			}
		}
	}

	fmt.Printf("Total de testes completos: %d\n", len(testSuite.Results))
}

type IGraph interface {
	AddVertex(val int) *Graph
	AddEdge(from, to, weight int) bool
	HasEdge(from, to int) bool
	GetVertex(val int) *Vertex
	GetDensity() float32
	GetNeighbors(n int) []int
	GetLen() int
	TraverseGraphSimple(start int)
	TraverseGraphDFS()
	BacktrackColoring(n int) bool
	GreedyColoring(n int) bool
	GreedyColoringByDegree(n int) bool
	PrintGraph()
}

// valor do vertice -> cor
type coloredMap map[int]int

type graphDensity uint

const (
	sparse graphDensity = iota
	mid
	dense
)

type Graph struct {
	Vertices map[int]*Vertex
	Len      int
	EdgesLen int
}

type Vertex struct {
	val   int
	edges map[int]*Edge
}

type Edge struct {
	weight int
	dest   *Vertex
}

func NewGraph() IGraph {
	return &Graph{
		Vertices: make(map[int]*Vertex),
		Len:      0,
		EdgesLen: 0,
	}
}

func GenerateGraph(nVertices int, density float32) IGraph {
	graph := NewGraph()

	for n := range nVertices {
		graph.AddVertex(n + 1)
	}

	// número máximo de arestas de um grafo não direcionado é n(n - 1) / 2
	maxEdges := nVertices * (nVertices - 1) / 2
	targetEdges := int(float32(maxEdges) * density)

	edgesAdded := 0
	for edgesAdded < targetEdges {
		i := rand.Intn(nVertices) + 1
		j := rand.Intn(nVertices) + 1

		if i != j && !graph.HasEdge(i, j) {
			weight := rand.Intn(10) + 1
			if graph.AddEdge(i, j, weight) {
				edgesAdded++
			}
		}
	}

	return graph
}

func genDensity(density graphDensity) float32 {
	switch density {
	case sparse:
		return 0.1
	case mid:
		return 0.5
	case dense:
		return 1
	default:
		return 0.5
	}
}

func (g *Graph) GetVertex(val int) *Vertex {
	if vertex, exists := g.Vertices[val]; exists {
		return vertex
	}

	return nil
}

func (g *Graph) AddVertex(val int) *Graph {
	vertex := g.GetVertex(val)
	if vertex != nil {
		return g
	}

	vertex = &Vertex{
		val:   val,
		edges: make(map[int]*Edge),
	}

	g.Vertices[val] = vertex
	g.Len++

	return g
}

func (g *Graph) AddEdge(from, to, weight int) bool {
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

func (g *Graph) GetNeighbors(n int) []int {
	vertex := g.GetVertex(n)
	if vertex == nil {
		return []int{}
	}

	var neighbors []int
	for dest := range vertex.edges {
		neighbors = append(neighbors, dest)
	}

	return neighbors
}

func (g *Graph) GetLen() int {
	return g.Len
}

func (g *Graph) GetDensity() float32 {
	maxEdges := g.Len * (g.Len - 1) / 2

	if maxEdges == 0 {
		return 0
	}

	return float32(g.EdgesLen) / float32(maxEdges)
}

func (g *Graph) TraverseGraphSimple(start int) {
	visited := make(map[int]bool, g.Len)

	for val := range g.Vertices {
		if exists := visited[val]; !exists {
			visited[val] = true
		}
	}

}

func (g *Graph) HasEdge(from, to int) bool {
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

func (g *Graph) TraverseGraphDFS() {
	visited := make(map[int]bool, g.Len)

	for dest, vertex := range g.Vertices {
		if exists := visited[dest]; !exists {
			g.DFS(vertex, &visited)
		}
	}
}

func (g *Graph) DFS(vertex *Vertex, visited *map[int]bool) {
	if _, exists := (*visited)[vertex.val]; exists {
		fmt.Printf("vertex already visited: %#v\n", vertex)
	}

	// pode fazer processamento aqui - coloracao de grafos e etc...

	(*visited)[vertex.val] = true

	neighbors := g.GetNeighbors(vertex.val)
	for _, neighbor := range neighbors {
		if _, exists := (*visited)[neighbor]; !exists {
			g.DFS(g.GetVertex(neighbor), visited)
		}
	}
}

// Backtrack O(k^V)
// Estratégia: tenta todas possibilitades, se nao funcionar faz o backtrack
func (g *Graph) BacktrackColoring(n int) bool {
	coloredMap, verticesList := g.makeColorMapAndVerticesList()
	if g.doColor(verticesList, 0, n, coloredMap) {
		g.printColors(coloredMap)
		return true
	} else {
		fmt.Printf("nenhuma solução com %d cores\n", n)
		return false
	}
}

// GREEDY O(V²)
// Estratégia: processar vertices um por um, e a cada vertice, assinar a menor cor que nao fique em conflito com as já existentes nos vizinhos
// 1: Escolhe uma ordenacao dos vertices (random, peso, alfabetico, etc..) - pode melhorar muito o resultado
// 2: Para cada vertice da ordem:
//
//	a. Olha para seus vizinhos já coloridos.
//	b. Acha quais cores estao proibidas (as usadas pelos vizinhos).
//	c. Pega a menenor cor disponivel que nao é proibida
//
// 3: Assina uma cor e continua. Colore o vertice na cor escolhida e move para o próximo, e nunca muda decieso anteriores.
// 4: Termina com todos vertices coloridos ou nao
func (g *Graph) GreedyColoring(n int) bool {
	coloredMap, verticesList := g.makeColorMapAndVerticesList()
	sort.Ints(verticesList)

	for _, v := range verticesList {
		availableColor := g.availableColor(n, v, coloredMap)
		if availableColor == -1 {
			return false
		}

		(*coloredMap)[v] = availableColor
	}

	g.printColors(coloredMap)
	return true
}

func (g *Graph) GreedyColoringByDegree(n int) bool {
	coloredMap, verticesList := g.makeColorMapAndVerticesList()
	ordered := g.sortByDegree(verticesList)

	for _, v := range ordered {
		availableColor := g.availableColor(n, v, coloredMap)
		if availableColor == -1 {
			return false
		}

		(*coloredMap)[v] = availableColor
	}

	g.printColors(coloredMap)

	return true
}

func (g *Graph) printColors(colors *coloredMap) {
	for v, c := range *colors {
		fmt.Printf("[vertex: %d - cor: %d], ", v, c)
	}

	fmt.Printf("\n")
}

func (g *Graph) sortByDegree(verticesList []int) []int {
	sort.Slice(verticesList, func(i, j int) bool {
		degreeI := len(g.GetNeighbors(verticesList[i]))
		degreeJ := len(g.GetNeighbors(verticesList[j]))
		return degreeI > degreeJ // ordem descendente
	})

	return verticesList
}

func (g *Graph) makeColorMapAndVerticesList() (*coloredMap, []int) {
	coloredMap := make(coloredMap, g.Len)
	var vertices []int

	for _, v := range g.Vertices {
		coloredMap[v.val] = -1
		vertices = append(vertices, v.val)
	}

	return &coloredMap, vertices
}

func (g *Graph) doColor(vertices []int, vertexIdx, maxColors int, coloredMap *coloredMap) bool {
	// base case: todos vertices coloridos
	if vertexIdx >= len(vertices) {
		return true
	}

	currentVertex := vertices[vertexIdx]
	// brute force -> tenta todas as cores nesse vertice especifico
	for color := 0; color < maxColors; color++ {
		if g.canColor(currentVertex, color, coloredMap) {
			(*coloredMap)[currentVertex] = color

			// recursao para o proximo vertice
			if g.doColor(vertices, vertexIdx+1, maxColors, coloredMap) {
				return true // achou a solucao completa
			}

			// backtrack: desfaz a escolha
			(*coloredMap)[currentVertex] = -1
		}
	}

	return false
}

func (g *Graph) canColor(currentVertex, color int, coloredMap *coloredMap) bool {
	neighbors := g.GetNeighbors(currentVertex)
	for _, neighbor := range neighbors {
		if (*coloredMap)[neighbor] == color {
			return false
		}
	}

	return true
}

func (g *Graph) availableColor(numColors, vertexVal int, coloredMap *coloredMap) int {
	neighbors := g.GetNeighbors(vertexVal)
	forbiddenColors := make(map[int]bool)

	for _, neighbor := range neighbors {
		color, exists := (*coloredMap)[neighbor]
		if exists && color != -1 {
			forbiddenColors[color] = true
		}
	}

	for color := 0; color < numColors; color++ {
		if !forbiddenColors[color] {
			return color
		}
	}

	return -1
}

func (g *Graph) PrintGraph() {
	for k, v := range g.Vertices {
		fmt.Printf("key is: %v - Val: %v", k, v)

		if edge, exists := v.edges[k]; exists {
			fmt.Printf("connection edge: %#v ", edge)
		}

		fmt.Printf("\n")
	}
}
