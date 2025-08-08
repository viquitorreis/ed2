package main

import "fmt"

func main() {
	graph := NewGraph()
	graph.AddVertex(1).AddVertex(20).AddVertex(14)
	graph.AddEdge(1, 14, 5)
	graph.AddEdge(1, 20, 10)
	graph.PrintGraph()

	neighbors := graph.GetNeighbors(1)
	fmt.Println("neighbors", neighbors)
}

type IGraph interface {
	AddVertex(d int) *Graph
	AddEdge(from, to, weight int) *Edge
	GetVertex(d int) *Vertex
	GetNeighbors(vertex int) []int
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

	fromVertex := g.GetVertex(from)

	var edge *Edge
	// 2 - verificar se a aresta ja existe (apenas nesse vértice) - precisamos fazer apenas no vértice FROM (de destino)
	// se queremos adicionar na aresta 1 -> 10, precisamos olhar apenas nas arestas do vértice 1
	if oldEdge, exists := fromVertex.edges[to]; exists {
		edge = &Edge{
			weight: weight,
			dest:   oldEdge.dest,
		}

	} else {
		edge = &Edge{
			weight: weight,
			dest:   g.GetVertex(to),
		}
	}

	// 3 - Armazena o valor no map de arestas do vertice FROM
	fromVertex.edges[to] = edge

	return edge
}

func (g *Graph) GetVertex(d int) *Vertex {
	if vertex, exists := g.Vertices[d]; exists {
		return vertex
	}

	return nil
}

func (g *Graph) GetNeighbors(vertex int) []int {
	// 1 - verificar se o vertex existe
	targetVertex := g.GetVertex(vertex)
	if targetVertex == nil {
		return []int{}
	}

	// 2 - verificar o map de arestas e ver se tem algo
	var neighbors []int
	for k := range targetVertex.edges {
		neighbors = append(neighbors, int((*g.GetVertex(k)).val))
	}

	return neighbors
}

func (g *Graph) PrintGraph() {
	if g == nil {
		return
	}

	for k, v := range g.Vertices {
		fmt.Printf("Vertex - Chave: %d Valor: %v", k, *v)

		for _, val := range v.edges {
			if val != nil {
				fmt.Printf(" -> conecta ao vértice %v com peso %d", val.dest, val.weight)
			} else {
				fmt.Printf(" sem conexões")
			}
		}

		fmt.Printf("\n")
	}
}
