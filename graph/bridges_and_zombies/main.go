package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Em uma ilha, voce (o estagi´ario) acidentalmente libera uma orda de zumbis.
// Na ilha est˜ao: vocˆe, um pesquisador, o professor e o zelador. Para fugir da
// ilha, existe um velha ponte de corda. O plano ´e simples: atravessar a ponte
// antes dos zumbis e cortar as cordas!!
// Os zumbis est˜ao a 17 minutos de distˆancia (calculos do professor). Para
// atravessar a ponte gasta-se:
// Vocˆe: 1 minuto
// Pesquisador: 2 minutos
// Zelador: 5
// Professor: 10 minutos
// A ponte tem capacidade para 2 pessoas...e est´a de noite e voces tem apenas
// uma velha lanterna1
// , necess´aria para cruzar a ponte.
// Pergunta: ´e poss´ıvel salvar a todos??
func init() {
	once.Do(func() {
		timer = time.NewTicker(time.Second * 17)
	})
}

func main() {
	graph := GenerateGraph()
	go initTimer()
	if graph.TraverseBridge(
		TraverseOrder{v1: Eu, v2: Pesquisador, doBack: Eu},
		TraverseOrder{v1: Professor, v2: Zelador, doBack: Pesquisador},
		TraverseOrder{v1: Eu, v2: Pesquisador, doBack: ""},
	) {
		// if graph.TraverseBridge(
		// 	TraverseOrder{v1: Eu, v2: Professor, doBack: Professor},
		// 	TraverseOrder{v1: Eu, v2: Professor, doBack: Professor},
		// 	TraverseOrder{v1: Eu, v2: Professor, doBack: Professor},
		// ) {
		log.Println("Sucesso! Conseguiram atravessar.")
	}
}

var (
	timer *time.Ticker
	once  sync.Once
)

func initTimer() {
	count := 0
	for {
		select {
		case <-timer.C:
			log.Fatalln("tempo esgotado!")
		default:
			time.Sleep(time.Second * 1)
			count++
			fmt.Printf("%dm passou\n", count)
		}
	}
}

type IGraph interface {
	AddVertex(val int, personType personType) *Graph
	GetVertex(val int) *Vertex
	GetVertexByType(personType personType) *Vertex
	TraverseBridge(traverseOrder ...TraverseOrder) bool
	CheckAllTraversed() bool
	InitializeSide()
	MovePeople(v1, v2 *Vertex)
	PrintGraph()
}

type personType string

const (
	Eu          personType = "eu"
	Pesquisador personType = "pesquisador"
	Zelador     personType = "zelador"
	Professor   personType = "professor"
)

type Graph struct {
	vertices    map[int]*Vertex
	initialSide *[]personType
	finalSide   *[]personType
}

type Vertex struct {
	val         int
	personType  personType
	didTraverse bool
}

type TraverseOrder struct {
	v1, v2 personType
	doBack personType
}

func NewGraph() IGraph {
	initialState := []personType{}
	finalState := []personType{}

	return &Graph{
		vertices:    make(map[int]*Vertex),
		initialSide: &initialState,
		finalSide:   &finalState,
	}
}

func GenerateGraph() IGraph {
	graph := NewGraph()
	graph.AddVertex(1, Eu).AddVertex(2, Pesquisador).AddVertex(5, Zelador).AddVertex(10, Professor)
	graph.InitializeSide()
	return graph
}

func (g *Graph) AddVertex(val int, personType personType) *Graph {
	if g.GetVertex(val) != nil {
		return g
	}

	vertex := &Vertex{
		val:        val,
		personType: personType,
	}

	g.vertices[val] = vertex

	return g
}

func (g *Graph) GetVertex(val int) *Vertex {
	vertex, exists := g.vertices[val]
	if !exists {
		return nil
	}

	return vertex
}

func (g *Graph) GetVertexByType(personType personType) *Vertex {
	for _, vertex := range g.vertices {
		if vertex.personType == personType {
			return vertex
		}
	}

	return nil
}

func (g *Graph) InitializeSide() {
	for _, vertex := range g.vertices {
		*g.initialSide = append(*g.initialSide, vertex.personType)
	}
}

func (g *Graph) TraverseBridge(traverseOrder ...TraverseOrder) bool {
	log.Println("Atravessando a ponte...")

	for _, travOrder := range traverseOrder {
		fmt.Printf("Atravessando -> %s e %s\n", travOrder.v1, travOrder.v2)

		v1 := g.GetVertexByType(travOrder.v1)
		v2 := g.GetVertexByType(travOrder.v2)
		if v1 == nil || v2 == nil {
			fmt.Printf("Vertice nao existe!")
			return false
		}

		g.MovePeople(v1, v2)
		time.Sleep(time.Duration(maxInt(v1.val, v2.val)) * time.Second)

		backVertex := g.GetVertexByType(travOrder.doBack)
		if backVertex != nil {
			g.MovePersonBack(backVertex)
			time.Sleep(time.Duration(backVertex.val) * time.Second)
		}

	}

	return len(*g.finalSide) == 4
}

func (g *Graph) MovePeople(v1, v2 *Vertex) {
	newInitialSide := []personType{}
	for _, person := range *g.initialSide {
		if person != v1.personType && person != v2.personType {
			newInitialSide = append(newInitialSide, person)
		}
	}

	*g.finalSide = append(*g.finalSide, v1.personType)
	*g.finalSide = append(*g.finalSide, v2.personType)

	*g.initialSide = newInitialSide
}

func (g *Graph) MovePersonBack(v *Vertex) {
	fmt.Printf("<- voltando %s\n", v.personType)
	if len(*g.finalSide) == 4 {
		return
	}

	newFinalSide := []personType{}
	for _, person := range *g.finalSide {
		if person != v.personType {
			newFinalSide = append(newFinalSide, person)
		}
	}

	*g.initialSide = append(*g.initialSide, v.personType)
	*g.finalSide = newFinalSide
}

func (g *Graph) CheckAllTraversed() bool {
	return true
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (g *Graph) PrintGraph() {
	for k, v := range g.vertices {
		fmt.Printf("graph type: %s - key is: %v - Val: %d", v.personType, k, v.val)

		fmt.Printf("\n")
	}
}
