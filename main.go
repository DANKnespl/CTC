package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

func load_config(path string) {
	// Read YAML file
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	// Unmarshal YAML data into Config struct
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("error unmarshaling YAML: %v", err)
	}
}

func main() {
	load_config("./config.yaml")

	// Initialize stations
	carHeap := &CarHeap{}
	heapMutex := sync.Mutex{}
	construct_stations()
	var stationQueues []*list.List
	for _, station := range stations {
		stationQueues = append(stationQueues, station.Queue)
	}

	// Initialize registers
	maxCount := 0

	for i := 0; i < len(station_types); i++ {
		maxCount += len(stations[station_types[i]].Substations)
	}
	registers := construct_registers(config.Registers.Count, config.Registers.HandleTimeMin, config.Registers.HandleTimeMax, maxCount)

	// Generate cars
	fmt.Println("1. Generating Cars")
	cars := construct_cars(config.Cars.Count, config.Cars.ArrivalTimeMin, config.Cars.ArrivalTimeMax)

	// Simulation
	fmt.Println("2. Cars sorting to queues")
	for _, car := range cars {
		for _, station := range stations {
			if station.Type == car.Type {
				station.mu.Lock()
				station.Queue.PushBack(car)
				station.mu.Unlock()
				break
			}
		}
	}

	// Spawn goroutines for managing cars at each station

	fmt.Println("3. Creating goroutines to manage cars")
	for _, station := range stations {
		for i := 0; i < len(station.Substations); i++ {
			go manage_cars(station, carHeap, &heapMutex)
		}
	}
	for i := 0; i < len(registers); i++ {
		go manage_heap(registers, carHeap, &heapMutex)
	}

	// Simulation progress
	fmt.Println("4. Managing Cars")
	startTime := time.Now()
	cars_in_Qs := 1
	for cars_in_Qs > 0 {
		time.Sleep(1000 * time.Millisecond)
		cars_in_Qs = 0
		for i := 0; i < len(station_types); i++ {
			cars_in_Qs += stations[station_types[i]].Queue.Len()
		}
		fmt.Println("       cars in Queues: ", cars_in_Qs)
	}
	for carHeap.Len() > 0 {
		time.Sleep(1000 * time.Millisecond)
	}
	// Aggregate data
	aggregate_data(cars, int(config.Cars.Count))
	fmt.Println("Duration of Simulation: ", time.Since(startTime))

}
