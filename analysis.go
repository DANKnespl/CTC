package main

import (
	"fmt"
	"math"
	"time"
)

func aggregate_data(cars []*Car, car_count int) int {
	if len(station_types) > 0 {
		start_time := time.Now()
		counts := make([]int64, len(station_types))
		queue_max := make([]time.Duration, len(station_types))
		time_min := make([]time.Time, len(station_types))
		time_max := make([]time.Time, len(station_types))
		var time_leaving_max time.Time
		queue_times := make([]time.Duration, len(station_types))
		var time_arive_to_leave time.Duration
		for i := 0; i < car_count; i++ {
			station_index := get_station_index(cars[i])
			if station_index != -1 {
				if counts[station_index] == 0 {
					time_min[station_index] = cars[i].time_of_station
				}
				counts[station_index] += 1
				time_max[station_index] = cars[i].time_of_leaving
				if cars[i].time_of_station.Sub(cars[i].time_of_arrival) > queue_max[station_index] {
					queue_max[station_index] = cars[i].time_of_station.Sub(cars[i].time_of_arrival)
				}

				if queue_times[station_index]+cars[i].time_of_station.Sub(cars[i].time_of_arrival) < queue_times[station_index] {
					queue_times[station_index] = time.Duration(math.MaxInt64)
				} else {
					queue_times[station_index] += cars[i].time_of_station.Sub(cars[i].time_of_arrival)
				}

				if cars[i].time_of_leaving.After(time_leaving_max) {
					time_leaving_max = cars[i].time_of_leaving
				}
			}
		}

		fmt.Println("Stations:")
		for station_index, station := range station_types {
			if counts[station_index] > 0 {
				if queue_times[station_index] != time.Duration(math.MaxInt64) {
					fmt.Printf("   %s:\n     Cars Count: %d\n     Cumulative time in Queue: %v\n     Average time in Queue: %v\n     Longest time in Queue: %v\n", station, counts[station_index], queue_times[station_index], time.Duration(int64(queue_times[station_index])/counts[station_index]), queue_max[station_index])
				} else {
					fmt.Printf("   %s:\n     Cars Count: %d\n     Cumulative time in Queue: %v\n     Average time in Queue: %v\n     Longest time in Queue: %v\n     Queue values are invalid, there was an int64 overflow!\n", station, counts[station_index], queue_times[station_index], time.Duration(int64(queue_times[station_index])/counts[station_index]), queue_max[station_index])
				}

			} else {
				fmt.Printf("   %s:\n     Cars Count: %d\n     Cumulative time in Queue: %v\n     Average time in Queue: %ds\n     Longest time in Queue: %v\n", station, counts[station_index], queue_times[station_index], 0, queue_max[station_index])
			}
		}
		fmt.Println("Registers")
		if car_count > 0 {
			time_arive_to_leave = time_leaving_max.Sub(cars[0].time_of_arrival)
		}
		fmt.Printf("  Cars Count: %d\n  Time from first Arrival to last Leaving: %v\n", sum_int_array(counts), time_arive_to_leave)
		fmt.Println("\nDuration of Analysis: ", time.Since(start_time))
		return 0
	}
	fmt.Println("No stations? No throughput!")
	return 1
}
func sum_int_array(arr []int64) int64 {
	var total int64
	total = 0
	for _, num := range arr {
		total += num
	}
	return total
}

func get_station_index(car *Car) int {
	for i, station_type := range station_types {
		if car.Type == station_type {
			return i
		}
	}
	return -1 // Return -1 if station type not found
}
