/*   _____ _           _        _        _____          
 *  |  ___| |         | |      (_)      |  ___|         
 *  | |__ | | ___  ___| |_ _ __ _  ___  | |__ _ __ __ _ 
 *  |  __|| |/ _ \/ __| __| '__| |/ __| |  __| '__/ _` |
 *  | |___| |  __/ (__| |_| |  | | (__  | |__| | | (_| |
 *  \____/|_|\___|\___|\__|_|  |_|\___| \____/_|  \__,_|
 *
 *  Author: Thomas Hansen
 *  Version: 0.0.1
 *  URL: https://gitlab.com/electric-era-public/coding-challenge-charger-uptime
**/
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Electric Era Interview Problem")

    cm := NewChargingMonitor()

    // Add some stations and chargers
    cm.AddStation(1)
    cm.AddCharger(1, 1, 10, 0)
    cm.AddCharger(1, 2, 10, 0)
    cm.AddCharger(1, 2, 20, 0)
    cm.AddCharger(1, 2, 0, 10)
    cm.AddCharger(1, 3, 0, 10)
    cm.AddCharger(1, 3, 0, 40)
    cm.AddCharger(1, 3, 20, 0)
    cm.AddStation(2)
    cm.AddCharger(2, 3, 0, 10)

    // List chargers for station 1
    if chargers, ok := cm.ListChargers(1); ok {
        fmt.Println("Station 1 chargers:")
        for _, c := range chargers {
            // fmt.Printf("ChargerID: %d, Status up: %d, down:%d \n", c.ChargerID, c.UpTime, c.DownTime)
            fmt.Printf("ChargerID: %d, Status %d\n", c.ChargerID, cm.CalcUptime(1, c.ChargerID))
        }
    }
}
