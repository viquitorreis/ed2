package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	fName := "coloring/test_suites.txt"
	results := parseTestFile(fName)
	generateCleanTable(results)
}

type ResultSummary struct {
	Vertices    int
	Density     float64
	Algorithm   string
	AvgTime     time.Duration
	MinColors   int
	SuccessRate float64
	AvgColors   float64
}

func parseTestFile(filename string) []ResultSummary {
	file, _ := os.Open(filename)
	defer file.Close()

	var results []ResultSummary
	scanner := bufio.NewScanner(file)

	var current ResultSummary
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "Resultados para") {
			fmt.Sscanf(line, "Resultados para %d vertices", &current.Vertices)
		} else if strings.HasPrefix(line, "Densidade") {
			fmt.Sscanf(line, "Densidade %f", &current.Density)
		} else if strings.HasPrefix(line, "Algoritmo") {
			current.Algorithm = strings.TrimSuffix(strings.Replace(line, "Algoritmo ", "", 1), ":")
		} else if strings.HasPrefix(line, "Tempo médio:") {
			timeStr := strings.Split(line, ": ")[1]
			current.AvgTime, _ = time.ParseDuration(timeStr)
		} else if strings.HasPrefix(line, "Cores médias usadas:") {
			fmt.Sscanf(line, "Cores médias usadas: %f", &current.AvgColors)
		} else if line == "---" {
			if current.Vertices > 0 && current.AvgColors > 0 {
				results = append(results, current)
			}
			current = ResultSummary{}
		}
	}

	return results
}

func generateCleanTable(results []ResultSummary) {
	scenarios := make(map[string]map[string]ResultSummary)

	for _, r := range results {
		key := fmt.Sprintf("%dv-%.1fd", r.Vertices, r.Density)

		if scenarios[key] == nil {
			scenarios[key] = make(map[string]ResultSummary)
		}

		if existing, exists := scenarios[key][r.Algorithm]; exists {
			if r.AvgColors > existing.AvgColors {
				scenarios[key][r.Algorithm] = r
			}
		} else {
			scenarios[key][r.Algorithm] = r
		}
	}

	fmt.Println("| Cenário | Backtrack | Greedy | Greedy-Degree | Cores Mínimas (BT/G/GD) |")
	fmt.Println("|---------|-----------|--------|---------------|-----------------|")

	keys := make([]string, 0, len(scenarios))
	for k := range scenarios {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		algos := scenarios[key]

		bt := algos["backtrack"]
		g := algos["greedy"]
		gd := algos["greedyByDegree"]

		btTime := formatTime(bt.AvgTime)
		gTime := formatTime(g.AvgTime)
		gdTime := formatTime(gd.AvgTime)

		fmt.Printf("| %s | %s | %s | %s | %.0f/%.0f/%.0f |\n",
			key, btTime, gTime, gdTime,
			bt.AvgColors, g.AvgColors, gd.AvgColors)
	}
}

func formatTime(d time.Duration) string {
	if d >= time.Second {
		return fmt.Sprintf("%.2fs", d.Seconds())
	} else if d >= time.Millisecond {
		return fmt.Sprintf("%.1fms", float64(d.Nanoseconds())/1e6)
	} else {
		return fmt.Sprintf("%.0fµs", float64(d.Nanoseconds())/1e3)
	}
}
