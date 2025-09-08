package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Measurement struct {
	Name  string
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

type Partial struct {
	Location string
	Temp     float64
}

func main() {
	start := time.Now()
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	dados := make(map[string]Measurement)
	var mu sync.Mutex

	lines := make(chan string, 1000)
	partials := make(chan Partial, 1000)
	var wg sync.WaitGroup


	go func() {
		scanner := bufio.NewScanner(measurements)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()


	numWorkers := 8 
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for rawData := range lines {
				semicolon := strings.Index(rawData, ";")
				if semicolon == -1 {
					continue
				}
				location := strings.TrimSpace(rawData[:semicolon])
				rawtemp := strings.TrimSpace(rawData[semicolon+1:])
				temp, err := strconv.ParseFloat(rawtemp, 64)
				if err != nil {
					continue
				}
				partials <- Partial{Location: location, Temp: temp}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(partials)
	}()
	
	for p := range partials {
		mu.Lock()
		measurement, ok := dados[p.Location]
		if !ok {
			measurement = Measurement{
				Name:  p.Location,
				Min:   p.Temp,
				Max:   p.Temp,
				Sum:   p.Temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, p.Temp)
			measurement.Max = max(measurement.Max, p.Temp)
			measurement.Sum += p.Temp
			measurement.Count++
		}
		dados[p.Location] = measurement
		mu.Unlock()
	}

	locations := make([]string, 0, len(dados))
	for name := range dados {
		locations = append(locations, name)
	}
	sort.Strings(locations)

	fmt.Printf("{")
	for i, name := range locations {
		measurement := dados[name]
		fmt.Printf("%s=%.1f/%.1f/%.1f", name, measurement.Min, measurement.Sum/float64(measurement.Count), measurement.Max)
		if i != len(locations)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("}\n")

	fmt.Println("Tempo de execução:", time.Since(start))
}
