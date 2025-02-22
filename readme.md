# Overview

This is the coding-challenge-charger-uptime for Electric Era.

Although I write most of my code in C/C++/Python, I thought this might be a fun opportunity to practice my Go, especially since it's statically typed and listed as a language under `Embedded Software Engineer`.

## Build and run

This program is written in Go, and I ran it with 1.22 using standard modules. It's broken up into `main.go`, which contains the file parsing and high-level logic, and `ChargingStation.go`, which manages the station and charger structs. All of the data is housed within a `ChargingMonitor` struct.

```
go run main.go ChargingStation.go "data/input_1.txt"
```

Alternatively, you may run

```
go build main.go ChargingStation.go
./main "data/input_2.txt"
```

The program will output a result file to the folder you run the program in, titled `out.txt`. It will also print to `stdout`.

The program will not finish if it encounters any errors, and will instead print `ERROR` followed by an error message.

## Assumptions

 - ChargerID's are unique, even on different stations.
 - The time values are not sequential or otherwise meaningful. No information is known about an epoch, and correct up/downtime calculations seem to ignore segments with no recorded information.
 - If no information is known about the charger status, we assume it's down.
