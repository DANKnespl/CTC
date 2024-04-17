package main

import (
	"container/heap"
	"math/rand"
	"sync"
	"time"
)

func manage_cars(station *Station, carHeap *CarHeap, mu *sync.Mutex) {
	for {
		minMemory := time.Time{}
		var minSubstation *Substation
		for _, substation := range station.Substations {
			substation.mu.RLock()
			if !substation.active && (minMemory.IsZero() || substation.memory.Before(minMemory)) {
				minMemory = substation.memory
				minSubstation = substation
			}
			substation.mu.RUnlock()
		}
		if minSubstation != nil {
			minSubstation.mu.Lock()
			if !minSubstation.active {
				minSubstation.active = true
				minSubstation.mu.Unlock()
				station.mu.Lock()
				if station.Queue.Len() > 0 {
					carElement := station.Queue.Front()
					if carElement != nil {
						car := carElement.Value.(*Car)
						station.Queue.Remove(carElement)
						station.mu.Unlock()

						if car.time_of_arrival.After(minSubstation.memory) {
							car.time_of_station = car.time_of_arrival
						} else {
							car.time_of_station = minSubstation.memory
						}
						serve_car(car, minSubstation, carHeap, mu)

						minSubstation.mu.Lock()
						minSubstation.memory = car.time_of_leaving
						minSubstation.active = false
						minSubstation.mu.Unlock()
					} else {
						station.mu.Unlock()
						time.Sleep(100 * time.Millisecond) // Check again after a short delay
					}
				} else {
					station.mu.Unlock()
					time.Sleep(100 * time.Millisecond) // Check again after a short delay
				}
			} else {
				minSubstation.mu.Unlock()
			}
		}
	}
}

func serve_car(car *Car, station *Substation, car_heap *CarHeap, mu *sync.Mutex) time.Time {
	serveTime := time.Duration(rand.Intn(int(station.ServeTimeMax-station.ServeTimeMin))) + station.ServeTimeMin

	car.time_to_serve = serveTime

	mu.Lock()
	heap.Push(car_heap, car)
	mu.Unlock()

	for (car.time_of_leaving == time.Time{}) {
		time.Sleep(10 * time.Millisecond)
	}
	//if car.Type == "lpg" && car.time_of_arrival.Before(car.time_of_station) {
	//fmt.Println(car.ID, "TOA", car.time_of_arrival, "TOS", car.time_of_station, "TOR", car.time_of_register, "TOL", car.time_of_leaving, "\n\n ")
	//}
	return car.time_of_leaving
}

func manage_heap(registers []*Register, car_heap *CarHeap, mu *sync.Mutex) {
	for {
		minRegisterIndex := -1
		minMemory := time.Time{} // Initialize with zero time
		for i, reg := range registers {
			reg.mu.RLock()
			if !reg.active && minRegisterIndex == -1 || reg.memory.Before(minMemory) {
				if minRegisterIndex != -1 {
					registers[minRegisterIndex].active = false
				}
				minMemory = reg.memory
				minRegisterIndex = i
				registers[minRegisterIndex].active = true
			}
			reg.mu.RUnlock() // Unlock each register after accessing its memory
		}
		if minRegisterIndex != -1 {
			mu.Lock()
			if car_heap.Len() > 0 {

				//registers[minRegisterIndex].mu.Unlock()
				car := heap.Pop(car_heap).(*Car)
				mu.Unlock()

				// Find the register with the lowest memory time

				// Update car's TimeOfRegister and TimeOfLeaving based on the selected register

				if (car.time_of_station.Add(car.time_to_serve)).After(minMemory) {
					car.time_of_register = car.time_of_station.Add(car.time_to_serve)
				} else {
					car.time_of_register = minMemory
				}
				car.time_of_leaving = car.time_of_register.Add(get_random_duration(registers[minRegisterIndex].HandleTimeMin, registers[minRegisterIndex].HandleTimeMax))

				// Update memory of the selected register
				//registers[minRegisterIndex].mu.Lock() // Lock the selected register before updating its memory
				registers[minRegisterIndex].mu.Lock()
				registers[minRegisterIndex].memory = car.time_of_leaving
				registers[minRegisterIndex].active = false
				registers[minRegisterIndex].mu.Unlock() // Unlock the selected register after updating its memory
			} else {
				//registers[minRegisterIndex].mu.Unlock()
				mu.Unlock()
				time.Sleep(100 * time.Millisecond) // Check again after a short delay
			}
		}
	}
}

func get_random_duration(min, max time.Duration) time.Duration {
	return min + time.Duration(rand.Int63n(int64(max-min)))
}
