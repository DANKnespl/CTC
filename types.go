package main

import (
	"container/list"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

// Define YAML structs to represent the configuration
type Config struct {
	Cars      CarConfig                `yaml:"cars"`
	Stations  map[string]StationConfig `yaml:"stations"`
	Registers RegisterConfig           `yaml:"registers"`
}

type CarConfig struct {
	Count          int64         `yaml:"count"`
	ArrivalTimeMin time.Duration `yaml:"ArrivalTimeMin"`
	ArrivalTimeMax time.Duration `yaml:"ArrivalTimeMax"`
}

type StationConfig struct {
	Count        int           `yaml:"count"`
	ServeTimeMin time.Duration `yaml:"ServeTimeMin"`
	ServeTimeMax time.Duration `yaml:"ServeTimeMax"`
}

type RegisterConfig struct {
	Count         int           `yaml:"count"`
	HandleTimeMin time.Duration `yaml:"HandleTimeMin"`
	HandleTimeMax time.Duration `yaml:"HandleTimeMax"`
}

// Structs of entities used in simulation
type Car struct {
	ID            int64
	Type          string
	time_to_serve time.Duration

	time_of_arrival  time.Time //timestamp of arrival 1-2 ms from last
	time_of_station  time.Time //timestamp of getting out of Q1
	time_of_register time.Time //timestamp of getting out of Q2 -> give to next in queue as TOR
	time_of_leaving  time.Time //timestamp of leaving -> give to next in Q1 as TOS
}

type Station struct {
	Type        string
	Queue       *list.List
	Substations []*Substation
	mu          sync.Mutex // Mutex to synchronize access to the Queue
}

type Substation struct {
	ServeTimeMin time.Duration
	ServeTimeMax time.Duration
	active       bool
	memory       time.Time
	mu           sync.RWMutex //is very much needed
}

type Register struct {
	HandleTimeMin time.Duration
	HandleTimeMax time.Duration
	active        bool
	memory        time.Time    //previous car time_of_leaving
	mu            sync.RWMutex //may not be needed, but I am not risking it //I may have been a fool there
}

// Car Heap (minheap) as a Queue for registers
type CarHeap []*Car

func (h CarHeap) Len() int { return len(h) }
func (h CarHeap) Less(i, j int) bool {
	return h[i].time_of_station.Add(h[i].time_to_serve).Before(h[j].time_of_station.Add(h[j].time_to_serve))
}

func (h CarHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *CarHeap) Push(x interface{}) {
	*h = append(*h, x.(*Car))
}
func (h *CarHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]  // Get the first element, which is the minimal element according to Less
	*h = old[0 : n-1] // Remove the minimal element from the heap
	return item
}

// register, car, station constructors based on contents of config
func construct_registers(count int, handleTimeMin, handleTimeMax time.Duration, maxCount int) []*Register {
	var reg []*Register
	if maxCount < count {

		fmt.Println("Actual number of registers: ", count)
		count = maxCount
		fmt.Println("Effective number of registers: ", count)
	}
	for i := 0; i < count; i++ {
		r := &Register{
			HandleTimeMin: handleTimeMin,
			HandleTimeMax: handleTimeMax,
			active:        false,
			memory:        time.Time{},
		}
		reg = append(reg, r)
	}
	return reg
}

func construct_cars(count int64, arrivalTimeMin, arrivalTimeMax time.Duration) []*Car {
	var cars []*Car
	lastArrival := time.Now()
	types := station_types
	if len(types) > 0 {
		for i := 0; int64(i) < count; i++ {
			car := &Car{
				ID:              int64(i),
				Type:            types[rand.Intn(len(types))],
				time_of_arrival: lastArrival.Add(get_random_duration(arrivalTimeMin, arrivalTimeMax)),
				time_of_leaving: time.Time{},
			}
			cars = append(cars, car)
			lastArrival = car.time_of_arrival
		}
	}
	return cars
}

func construct_stations() {
	stations = make(map[string]*Station)
	keys := make([]string, 0, len(config.Stations))
	for key := range config.Stations {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for key := 0; key < len(keys); key++ {
		station_type := keys[key]
		station_conf := config.Stations[keys[key]]
		if station_conf.Count > 0 {

			var substations []*Substation
			for i := 0; i < station_conf.Count; i++ {
				substation := &Substation{
					ServeTimeMin: station_conf.ServeTimeMin,
					ServeTimeMax: station_conf.ServeTimeMax,
				}
				substations = append(substations, substation)
			}
			station := &Station{
				Type:        station_type,
				Queue:       list.New(),
				Substations: substations,
			}
			stations[station_type] = station
			station_types = append(station_types, station_type)
		}
	}
}

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
