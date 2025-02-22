package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World!")

    cm := NewChargingMonitor()

    // Add some stations and chargers
    cm.AddStation(1)
    cm.AddCharger(1, 1, 10, 0)
    cm.AddCharger(1, 2, 10, 0)
    cm.AddCharger(1, 2, 20, 0)
    cm.AddCharger(1, 3, 0, 10)
    cm.AddCharger(1, 3, 0, 20)
    cm.AddStation(2)
    cm.AddCharger(2, 3, 0, 10)

    // List chargers for station 1
    if chargers, ok := cm.ListChargers(1); ok {
        fmt.Println("Station 1 chargers:")
        for _, c := range chargers {
            fmt.Printf("ChargerID: %d, Status up: %d, down:%d \n", c.ChargerID, c.UpTime, c.DownTime)
        }
    }
}
