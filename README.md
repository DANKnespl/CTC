# Statistiky
## minimální čas příjezdu auta > maximální doba zpracování
-> žádný čas ve frontách

### config.yaml
```yaml
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
```
### výstup
```
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
```

---
## různé doby zpracování a počty stanic 
-> někdy vznikne fronta, někdy ne
### config.yaml
```yaml
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
```
### výstup
```
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
```

---
## dlouhá doba u kasy 
-> všechny fronty se plní -> dlouhé fronty
### config.yaml
```yaml
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
  HandleTimeMin: 1h
  HandleTimeMax: 3h
```
### výstup
```
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
     Cars Count: 3767
     Cumulative time in Queue: 66728h54m2.763621747s
     Average time in Queue: 17h42m50.651118561s
     Longest time in Queue: 41h23m29.77151033s
   electric:
     Cars Count: 3842
     Cumulative time in Queue: 160890h11m41.254001607s
     Average time in Queue: 41h52m36.038847996s
     Longest time in Queue: 106h48m55.804484988s
   gas:
     Cars Count: 3674
     Cumulative time in Queue: 63652h55m13.028974387s
     Average time in Queue: 17h19m30.852756933s
     Longest time in Queue: 38h57m18.532447752s
   lpg:
     Cars Count: 3717
     Cumulative time in Queue: 154045h41m31.976487604s
     Average time in Queue: 41h26m36.796334809s
     Longest time in Queue: 104h41m55.093206187s
Registers
  Cars Count: 15000
  Time from first Arrival to last Leaving: 107h51m38.219965059s

Duration of Analysis:  2.829101ms
Duration of Simulation:  5.99137594s
```

---
## gas station bottleneck (2 stanice oproti 100/200)
-> malá paralelizace -> dlouhé fronty -> přetečení datového typu -> zavádějící výsledky
### config.yaml
```yaml
cars:
  count: 15000
  ArrivalTimeMin: 1ms
  ArrivalTimeMax: 3ms
stations:
  gas:
    count: 2
    ServeTimeMin: 2h
    ServeTimeMax: 5h
  lpg:
    count: 100
    ServeTimeMin: 4h
    ServeTimeMax: 7h
  diesel:
    count: 200
    ServeTimeMin: 3h
    ServeTimeMax: 6h
  electric:
    count: 100
    ServeTimeMin: 5h
    ServeTimeMax: 10h
registers:
  count: 50000
  HandleTimeMin: 1h
  HandleTimeMax: 3h
```
### výstup
```
Actual number of registers:  50000
Effective number of registers:  402
1. Generating Cars
2. Cars sorting to queues
3. Creating goroutines to manage cars
4. Managing Cars
(... # of cars remaining to be serviced)
5. Analyzing data
Stations:
   diesel:
     Cars Count: 3784
     Cumulative time in Queue: 221505h8m44.544004513s
     Average time in Queue: 58h32m14.282384779s
     Longest time in Queue: 134h26m11.813993358s
   electric:
     Cars Count: 3673
     Cumulative time in Queue: 643552h14m27.502919549s
     Average time in Queue: 175h12m41.793493852s
     Longest time in Queue: 468h7m8.973899652s
   gas:
     Cars Count: 3821
     Cumulative time in Queue: 2562047h47m16.854775807s
     Average time in Queue: 670h31m3.396193346s
     Longest time in Queue: 10593h48m44.47195704s
     Queue values might be invalid, there was an int64 overflow!
   lpg:
     Cars Count: 3722
     Cumulative time in Queue: 552745h22m18.275372812s
     Average time in Queue: 148h30m27.441771996s
     Longest time in Queue: 472h30m0.692719162s
Registers
  Cars Count: 15000
  Time from first Arrival to last Leaving: 10599h33m51.264556838s

Duration of Analysis:  1.836049ms
Duration of Simulation:  11m14.411859925s
```
