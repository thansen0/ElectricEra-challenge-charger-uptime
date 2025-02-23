# Overview

This is the coding-challenge-charger-uptime for Electric Era. [Original Prompt](./problem_statement.md) and [original repo](https://gitlab.com/electric-era-public/coding-challenge-charger-uptime).

Although I write most of my code in C/C++/Python, I thought this might be a fun opportunity to practice my Go, especially since it's statically typed and listed as a language under Embedded Software Engineer`.

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
 - Time complexity for the number of chargers is not a significant concern. While We could have tens of thousands of stations, a station may only have 1-50 chargers. With this assumption, I store stations in a `map` and chargers in a `slice` (i.e. array). There is a reverse lookup map (charger ID to station ID), however this doesn't hold or point to charger data.
 - Everything is stored in memory, as I felt databases/etc were beyond the scope of this assignment.
 - The time values are not sequential or otherwise meaningful. No information is known about an epoch, and correct up/downtime calculations seem to ignore segments with no recorded information.
 - If no information is known about the charger status, we assume it's down.
