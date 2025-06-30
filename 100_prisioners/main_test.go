package main

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestPrisonersSmall(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	trials := 10000
	successes := 0

	quantity := 100

	for i := 0; i < trials; i++ {
		boxes := genIntArr(quantity)
		prisioners := genPrisioners(quantity)

		if find(prisioners, boxes) {
			successes++
		}
	}

	successRate := float64(successes) / float64(trials) * 100
	t.Logf("Small test: %d trials, %.2f%% success rate", trials, successRate)

	if successRate < 25 || successRate > 35 {
		t.Errorf("Success rate %.2f%% seems off, expected around 30.685%%", successRate)
	}
}

// go test -v -run TestPrisonersMedium
func TestPrisonersMedium(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	trials := 1_000_000
	successes := 0

	t.Logf("Running %d trials...", trials)

	quantity := 100

	for i := 0; i < trials; i++ {
		boxes := genIntArr(quantity)
		prisioners := genPrisioners(quantity)

		if find(prisioners, boxes) {
			successes++
		}

		// vai mostrar os resultados a cada 100k
		if (i+1)%100_000 == 0 {
			currentRate := float64(successes) / float64(i+1) * 100
			t.Logf("Progress: %d trials, %.3f%% success rate", i+1, currentRate)
		}
	}

	successRate := float64(successes) / float64(trials) * 100
	t.Logf("Final: %d trials, %.4f%% success rate", trials, successRate)
	t.Logf("Theoretical: 30.685%%")
	t.Logf("Difference: %.4f%%", successRate-30.685)

	if successRate < 30.0 || successRate > 31.5 {
		t.Errorf("Success rate %.4f%% is too far from expected 30.685%%", successRate)
	}
}

// go test -v -run TestPrisonersLarge -timeout 30m
func TestPrisonersLarge(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large test in short mode")
	}

	rand.Seed(time.Now().UnixNano())

	trials := 10_000_000
	successes := 0

	quantity := 100

	t.Logf("Running %d trials (this will take a while)...", trials)
	start := time.Now()

	for i := 0; i < trials; i++ {
		boxes := genIntArr(quantity)
		prisioners := genPrisioners(quantity)

		if find(prisioners, boxes) {
			successes++
		}

		// mostra o progresso a cada 500k
		if (i+1)%500_000 == 0 {
			currentRate := float64(successes) / float64(i+1) * 100
			elapsed := time.Since(start)
			t.Logf("Progress: %d trials, %.4f%% success rate (elapsed: %v)",
				i+1, currentRate, elapsed.Round(time.Second))
		}
	}

	elapsed := time.Since(start)
	successRate := float64(successes) / float64(trials) * 100

	t.Logf("=== FINAL RESULTS ===")
	t.Logf("Total de tentativas: %d", trials)
	t.Logf("Sucessos: %d", successes)
	t.Logf("Taxa de sucesso: %.6f%%", successRate)
	t.Logf("Tempo gasto: %v", elapsed.Round(time.Second))
	t.Logf("Tentativas/segundo: %.0f", float64(trials)/elapsed.Seconds())
	t.Logf("Teoricamente: 30.685%%")
	t.Logf("Diferenca: %.6f%%", successRate-30.685)

	if successRate < 30.4 || successRate > 31.0 {
		t.Errorf("Success rate %.6f%% is too far from expected 30.685%%", successRate)
	}
}

func TestPrisonersBillion(t *testing.T) {
	if testing.Short() {
		t.Skip("pulando teste de 1 Bi - short mode")
	}

	const totalTrials = 1_000_000_000
	numWorkers := runtime.NumCPU()
	trialsPerWorker := totalTrials / numWorkers
	const reportInterval = 50_000_000
	quantity := 100

	var successes int64
	var completed int64
	var wg sync.WaitGroup

	t.Logf("Rodando %d tentativas usando %d threads...", totalTrials, numWorkers)
	start := time.Now()

	for worker := 0; worker < numWorkers; worker++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			localSuccesses := 0

			for i := 0; i < trialsPerWorker; i++ {
				boxes := genIntArr(quantity)
				prisioners := genPrisioners(quantity)

				if find(prisioners, boxes) {
					localSuccesses++
				}

				if (i+1)%1_000_000 == 0 {
					atomic.AddInt64(&successes, int64(localSuccesses))
					localSuccesses = 0

					current := atomic.AddInt64(&completed, 1_000_000)
					if current%reportInterval == 0 {
						currentSuccesses := atomic.LoadInt64(&successes)
						rate := float64(currentSuccesses) / float64(current) * 100
						elapsed := time.Since(start)
						t.Logf("Progress: %d trials, %.4f%% success rate (elapsed: %v)",
							current, rate, elapsed.Round(time.Second))
					}
				}
			}

			atomic.AddInt64(&successes, int64(localSuccesses))
		}(worker)
	}

	wg.Wait()

	elapsed := time.Since(start)
	successRate := float64(successes) / float64(totalTrials) * 100

	t.Logf("=== FINAL RESULTS ===")
	t.Logf("Total de tentativas: %d", totalTrials)
	t.Logf("Sucessos: %d", successes)
	t.Logf("Taxa de sucesso: %.6f%%", successRate)
	t.Logf("Tempo gasto: %v", elapsed.Round(time.Second))
	t.Logf("Tentativas/segundo: %.0f", float64(totalTrials)/elapsed.Seconds())
	t.Logf("Teoricamente: 30.685%%")
	t.Logf("Diferenca: %.6f%%", successRate-30.685)

	if successRate < 30.0 || successRate > 31.5 {
		t.Errorf("Success rate %.6f%% is too far from expected 30.685%%", successRate)
	}
}
