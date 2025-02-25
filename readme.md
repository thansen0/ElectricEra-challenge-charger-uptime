# Overview

This is the coding-challenge-charger-uptime for Electric Era. [Original Prompt](./problem_statement.md) and [original repo](https://gitlab.com/electric-era-public/coding-challenge-charger-uptime).

Although I write most of my code in C/C++/Python, I thought this might be a fun opportunity to practice my Go, especially since it's statically typed and listed as a language under Embedded Software Engineer.

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

The program will not finish if it encounters any errors, and will instead print `ERROR` followed by the error message.

## Discussion

When designing the data layout, I operated under the assumption that while there may be thousands of unique stations and chargers, any station will only have a relatively small number of associated chargers. Under this assumption, I created a hash map where you can look up chargers using a station ID, and a station ID's using a charger ID, however to access the actual charger data you must iterate through the slice stored in the station struct. Slices are effectively dynamic arrays, and while you normally wouldn't store a dynamic array in a struct in this case it's simply a pointer to an array held in memory, and we don't have to worry about struct padding or reallocation.

There are a few inefficiencies with this design. In particular, I don't like how getting a list of all stations with `ListStations()` is an O(n) operation, and then sorting them is O(nlog(n)). I could store an array of them, however then I have to store them for the duration of the program, as well as write an insertion sort. I think for the context of this program it's fine, however maintaining a slice of StationID's in `main.go` would also have been a valid design choice in my opinion.

I made two assumptions in regards to maintaining uptime and downtime. The first is that the epoch starts at the first available reading, and goes until the last available reading. The second is that the time comes in sequential order, and if it doesn't I report it as an error and exit the program.

## Assumptions

 - ChargerID's are unique, even on different stations.
 - Time complexity for the number of chargers is not a significant concern. While we could have tens of thousands of stations, a station is only expected to have a few chargers associated with it. With this assumption, I store stations in a `map` and chargers in a `slice` (i.e. array). There is a reverse lookup map (charger ID to station ID), however this doesn't hold or point to charger data, just the ID.
 - Everything is stored in memory, as I felt databases/etc were beyond the scope of the challenge prompt.
 - The time values must come in sequential order. The epoch is the first time given.
 - If no information is known about the charger, we assume it's down.
 - Uptime percents are truncated to integers
 - Although I believe unit tests are important for any project of importance, I felt they were beyond the scope of a take-home assignment.
 - While the actual prompt only called for printing to stdout, I felt it would also be appropriate to output to a file, in this case `out.txt`.
