# Statistiky
### config.yaml
cars:
  count: 1500000
  ArrivalTimeMin: 1h
  ArrivalTimeMax: 3h
stations:
  gas:
    count: 10000
    ServeTimeMin: 2ms
    ServeTimeMax: 5ms
  lpg:
    count: 10000
    ServeTimeMin: 4ms
    ServeTimeMax: 7ms
  diesel:
    count: 10000
    ServeTimeMin: 3ms
    ServeTimeMax: 6ms
  electric:
    count: 10000
    ServeTimeMin: 5ms
    ServeTimeMax: 10ms
registers:
  count: 50000
  HandleTimeMin: 1ms
  HandleTimeMax: 3ms

### výstup
Actual number of registers:  50000
Effective number of registers:  40000
1. Generating Cars
2. Cars sorting to queues
3. Creating goroutines to manage cars
4. Managing Cars
(... # of cars remaining to be serviced)
5. Analyzing data
Stations:
   diesel:
     Cars Count: 375682
     Cumulative time in Queue: 0s
     Average time in Queue: 0s
     Longest time in Queue: 0s
   electric:
     Cars Count: 374303
     Cumulative time in Queue: 0s
     Average time in Queue: 0s
     Longest time in Queue: 0s
   gas:
     Cars Count: 374914
     Cumulative time in Queue: 0s
     Average time in Queue: 0s
     Longest time in Queue: 0s
   lpg:
     Cars Count: 375101
     Cumulative time in Queue: 0s
     Average time in Queue: 0s
     Longest time in Queue: 0s
Registers
  Cars Count: 1500000
  Time from first Arrival to last Leaving: 2562047h47m16.854775807s

Duration of Analysis:  18.002013803s
Duration of Simulation:  6m55.212141629s
---
### config.yaml
cars:
  count: 15000
  ArrivalTimeMin: 1ms
  ArrivalTimeMax: 3ms
stations:
  gas:
    count: 200
    ServeTimeMin: 2ms
    ServeTimeMax: 5ms
  lpg:
    count: 100
    ServeTimeMin: 4ms
    ServeTimeMax: 7ms
  diesel:
    count: 200
    ServeTimeMin: 3ms
    ServeTimeMax: 6ms
  electric:
    count: 100
    ServeTimeMin: 5ms
    ServeTimeMax: 10ms
registers:
  count: 50000
  HandleTimeMin: 1ms
  HandleTimeMax: 3ms

### výstup
Actual number of registers:  50000
Effective number of registers:  600
1. Generating Cars
2. Cars sorting to queues
3. Creating goroutines to manage cars
4. Managing Cars
(... # of cars remaining to be serviced)
5. Analyzing data
Stations:
   diesel:
     Cars Count: 3750
     Cumulative time in Queue: 0s
     Average time in Queue: 0s
     Longest time in Queue: 0s
   electric:
     Cars Count: 3674
     Cumulative time in Queue: 2h49m11.887623327s
     Average time in Queue: 2.763170283s
     Longest time in Queue: 7.939790283s
   gas:
     Cars Count: 3905
     Cumulative time in Queue: 0s
     Average time in Queue: 0s
     Longest time in Queue: 0s
   lpg:
     Cars Count: 3671
     Cumulative time in Queue: 2h43m6.473054095s
     Average time in Queue: 2.665887511s
     Longest time in Queue: 7.96497517s
Registers
  Cars Count: 15000
  Time from first Arrival to last Leaving: 29.939486655s

Duration of Analysis:  3.321718ms
Duration of Simulation:  5.898450348s




