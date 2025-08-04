package main

import "fmt"

func main() {
	graph := NewGraph()
	graph.AddVertex(1).AddVertex(20)
	graph.PrintGraph()
}

type IGraph interface {
	AddVertex(d int) *Graph
	PrintGraph()
}

type Graph struct {
	// vertices vai conter um map das chaves de todos os vértices
	Vertices map[int]*Vertex
}

type Edges struct {
	weight int
	dest   *Vertex
}

type Vertex struct {
	val   int
	edges map[int]*Edges
}

func NewGraph() IGraph {
	return &Graph{
		Vertices: make(map[int]*Vertex),
	}
}

func (g *Graph) AddVertex(d int) *Graph {
	// 1 - se já existir, não cria, retorna o grafico para ficar idempotente
	if _, exists := g.Vertices[d]; exists {
		return g
	}

	// 2 - vertex novo, sem nenhuma conexao ainda...
	vertex := &Vertex{
		val:   d,
		edges: make(map[int]*Edges),
	}

	// 3 - registrar no grafo, adiciona o novo vertice nos vertices
	g.Vertices[d] = vertex

	return g
}

func (g *Graph) PrintGraph() {
	if g == nil {
		return
	}

	for k, v := range g.Vertices {
		fmt.Printf("key: %d Vertex: %+v\n", k, v)
	}
}
