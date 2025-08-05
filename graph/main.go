package main

import "fmt"

func main() {
	graph := NewGraph()
	graph.AddVertex(1).AddVertex(20).AddVertex(14)
	graph.PrintGraph()
	graph.AddEdge(1, 14, 5)
	graph.PrintGraph()
}

type IGraph interface {
	AddVertex(d int) *Graph
	AddEdge(from, to, weight int) *Edge
	GetVertex(d int) *Vertex
	PrintGraph()
}

type Graph struct {
	// vertices vai conter um map das chaves de todos os vértices
	Vertices map[int]*Vertex
}

type Edge struct {
	weight int
	dest   *Vertex
}

type Vertex struct {
	val   int
	edges map[int]*Edge
}

func NewGraph() IGraph {
	return &Graph{
		Vertices: make(map[int]*Vertex),
	}
}

func (g *Graph) AddVertex(d int) *Graph {
	// 1 - se já existir, não cria, retorna o grafo para ficar idempotente
	if _, exists := g.Vertices[d]; exists {
		return g
	}

	// 2 - vertex novo, sem nenhuma conexao ainda...
	vertex := &Vertex{
		val:   d,
		edges: make(map[int]*Edge),
	}

	// 3 - registrar no grafo, adiciona o novo vertice nos vertices
	g.Vertices[d] = vertex

	return g
}

func (g *Graph) AddEdge(from, to, weight int) *Edge {
	// 1 - verificar se ambos os vertices existem
	if g.Vertices[from] == nil || g.Vertices[to] == nil {
		fmt.Printf("From or to doesnt exist")
		return nil
	}

	edge := Edge{
		weight: weight,
		dest:   g.GetVertex(to),
	}

	count := 0
	// 2 - verificar se a aresta ja existe
	for _, vertex := range g.Vertices {
		count++
		// se ja existir, atualiza o peso
		if ed, exists := vertex.edges[to]; exists {
			fmt.Println("existe, adicionando")
			*ed = edge

			break
		}

		// se nao existe cria
		if count == len(g.Vertices)-1 && vertex.edges[to] == nil {
			fmt.Println("num existe, criando")
			vertex.edges = map[int]*Edge{
				to: &edge,
			}
		}
	}

	// 3 - Armazena o valor no map de arestas do vertice
	if fr, exists := g.Vertices[from]; exists {
		fmt.Println("hello world")
		*fr = Vertex{
			val: from,
			edges: map[int]*Edge{
				from: &edge,
			},
		}
	}

	return &edge
}

func (g *Graph) GetVertex(d int) *Vertex {
	if vertex, exists := g.Vertices[d]; exists {
		return vertex
	}

	return nil
}

func (g *Graph) PrintGraph() {
	if g == nil {
		return
	}

	for k, v := range g.Vertices {
		if edge, exists := v.edges[k]; exists {
			fmt.Printf("key: %d Vertex Val: %d, Vertex Edge Dest: %d Vertex Edge Weight: %+v\n", k, v.val, edge.dest.val, edge.weight)
		} else {
			fmt.Printf("key: %d Vertex Val: %d, Edge address: %v\n", k, v.val, edge)
		}
	}
}
