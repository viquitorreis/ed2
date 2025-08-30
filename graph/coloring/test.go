package main

import (
	"fmt"
	"os"
	"time"
)

type algorithm string

const (
	backtrack      algorithm = "backtrack"
	greedy         algorithm = "greedy"
	greedyByDegree algorithm = "greedyByDegree"
)

type TestResult struct {
	Vertices      int
	Density       float32
	Algorithm     algorithm
	ExecutionTime time.Duration
	MinColors     int
	Success       bool
	MaxKTested    int
}

type TestSuite struct {
	Results []TestResult
}

func (ts *TestSuite) AddResult(result TestResult) {
	ts.Results = append(ts.Results, result)
}

func runSingleTest(nVertices int, density float32, algorithm algorithm) TestResult {
	graph := GenerateGraph(nVertices, density)

	start := time.Now()
	minColors, success := findMinimumColors(graph, algorithm, 5) // Limite de 5 cores
	duration := time.Since(start)
	actualDensity := graph.GetDensity()

	return TestResult{
		Vertices:      nVertices,
		Density:       actualDensity,
		Algorithm:     algorithm,
		ExecutionTime: duration,
		MinColors:     minColors,
		Success:       success,
		MaxKTested:    20,
	}
}

func findMinimumColors(graph IGraph, algorithm algorithm, maxK int) (int, bool) {
	for k := 1; k <= maxK; k++ {
		var success bool
		switch algorithm {
		case backtrack:
			success = graph.BacktrackColoring(k)
		case greedy:
			success = graph.GreedyColoring(k)
		case greedyByDegree:
			success = graph.GreedyColoringByDegree(k)
		}

		if success {
			return k, true // k = numero minimo de cores encontrado
		}
	}
	return -1, false // nao conseguiu colorir com maxK cores
}

func runMultipleTests(nVertices int, density graphDensity, alg algorithm, repetitions int) []TestResult {
	var results []TestResult

	dens := genDensity(density)

	fmt.Printf("Testando: %d vertices, densidade %v, algoritmo %s, %d repetições\n",
		nVertices, dens, alg, repetitions)

	for i := 0; i < repetitions; i++ {
		fmt.Printf("  Execução %d/%d...\n", i+1, repetitions)
		result := runSingleTest(nVertices, dens, alg)
		results = append(results, result)
	}

	return results
}

func printAndWriteResults(results []TestResult) {
	if len(results) == 0 {
		return
	}

	var totalTime time.Duration
	var totalColors, successCount int

	for _, r := range results {
		totalTime += r.ExecutionTime
		if r.Success {
			totalColors += r.MinColors
			successCount++
		}
	}

	avgTime := totalTime / time.Duration(len(results))
	var avgColors float64
	if successCount > 0 {
		avgColors = float64(totalColors) / float64(successCount)
	}

	successRate := float64(successCount) / float64(len(results)) * 100

	testData := fmt.Sprintf(`Resultados para %d vertices, densidade %v, algoritmo %s:
		Tempo de execução: %v
		Tempo médio: %v
		Cores mínimas: %d
		Taxa de Sucesso: %.1f%% (%d/%d)
		Cores médias usadas: %.2f
		Sucesso: %v
		MaxK testados: %d
---
`, results[0].Vertices, results[0].Density, results[0].Algorithm,
		results[0].ExecutionTime, avgTime, results[0].MinColors, successRate,
		successCount, len(results), avgColors, results[0].Success, results[0].MaxKTested)

	// Fixed: Proper file appending
	file, err := os.OpenFile("test_suites.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(testData)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
