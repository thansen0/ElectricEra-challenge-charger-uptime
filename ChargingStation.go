package main

import (
    "sort"
    "fmt"
    "os"
)

// struct defining charger, associated
// with a station
type Charger struct {
    ChargerID uint32
    PrevTime uint64
    UpTime uint64
    DownTime uint64
}

// contains a slice of chargers, we assume the number
// of chargers per station is relatively small
type Station struct {
    StationID uint32
    Chargers []Charger
}

// used as instance for object, contains fast
// lookup for stations and station ID's
type ChargingMonitor struct {
    Stations map[uint32]*Station
    ChargerToStat map[uint32]uint32
}

// Creates nwe charging station map
func NewChargingMonitor() *ChargingMonitor {
    return &ChargingMonitor{
        Stations: make(map[uint32]*Station),
        ChargerToStat: make(map[uint32]uint32),
    }
}

// AddStation adds a new station if it doesn’t exist
func (cm *ChargingMonitor) AddStation(stationID uint32) {
    if _, exists := cm.Stations[stationID]; !exists {
        cm.Stations[stationID] = &Station{
            StationID: stationID,
            Chargers:  []Charger{},
        }
    }
}

// AddCharger adds a charger to a station
func (cm *ChargingMonitor) AddCharger(stationID uint32, chargerID uint32, start_time uint64, end_time uint64, available bool) {
    var uptime uint64
    var downtime uint64
    // set uptime and downtime variables, simplifies logic
    if available {
        uptime = end_time - start_time
        downtime = 0
    } else {
        uptime = 0
        downtime = end_time - start_time
    }

    // makes sure station exists, if not, add it
    station, exists := cm.Stations[stationID]
    if !exists {
        fmt.Printf("Creating new station: %d\n", stationID)
        cm.AddStation(stationID)
        station, _ = cm.Stations[stationID]
    }

    // insert into reverse lookup as well
    cm.ChargerToStat[chargerID] = stationID

    // find if charger exists, if so update
    for i := 0; i < len(station.Chargers); i++ {
        // found current charger in list, thus it exists
        if station.Chargers[i].ChargerID == chargerID {
            // check for relative order
            if station.Chargers[i].PrevTime > start_time {
                fmt.Println("ERROR: Up and down time from input file is out of order.")
                os.Exit(1)
            }

            if station.Chargers[i].PrevTime > 0 && station.Chargers[i].PrevTime < start_time {
                // update downtime with missing sequential time
                station.Chargers[i].DownTime += start_time - station.Chargers[i].PrevTime
            }
            station.Chargers[i].PrevTime = end_time

            // update uptime and downtime
            station.Chargers[i].UpTime += uptime
            station.Chargers[i].DownTime += downtime
            return
        }
    }

    // no charger exists, add one
    station.Chargers = append(station.Chargers, Charger{
        ChargerID: chargerID,
        PrevTime:  end_time,
        UpTime:    uptime,
        DownTime:  downtime,
    })

    // assign station back into object struct
    // cm.Stations[stationID] = station
}

// Gets station ID from charger ID
func (cm *ChargingMonitor) GetStationID(chargerID uint32) uint32 {
    return cm.ChargerToStat[chargerID]
}

// Creates a slice of station IDs in O(n) time
func (cm *ChargingMonitor) ListStations() []uint32 {
    var keys []uint32
    for stationID := range cm.Stations {
        keys = append(keys, stationID)
    }

    // slice comes unsorted, but we prefer it sorted on printout
    sort.Slice(keys, func(i, j int) bool {
        return keys[i] < keys[j]
    })

    return keys
}

// ListChargers returns all chargers for a station
func (cm *ChargingMonitor) ListChargers(stationID uint32) ([]Charger, bool) {
    station, exists := cm.Stations[stationID]
    if !exists {
        return nil, false
    }
    return station.Chargers, true
}

// Calculates station uptime as a truncated percent out of 100
func (cm *ChargingMonitor) CalcStationUptime(stationID uint32) uint64 {
    chargers, exists := cm.ListChargers(stationID)    
    if !exists {
        // doesn't exist, so 0 according to problem statement
        return 0
    }

    var tot_time uint64 = 0
    var up_time uint64 = 0
    for _, c := range chargers {
        u, t := cm.calcChargerUptime(stationID, c.ChargerID)
        tot_time += t
        up_time += u
    }

    if tot_time == 0 {
        // doesn't exist, so 0 according to problem statement
        return tot_time
    }

    return (up_time * 100) / tot_time
}

// returns the uptime and total_time of a particular charger at a particular
// station, returns (0,0) if the charger is invalid
func (cm *ChargingMonitor) calcChargerUptime(stationID uint32, chargerID uint32) (uint64, uint64) {
    station, exists := cm.Stations[stationID]
    if !exists {
        return 0, 0
    }

    for _, charger := range station.Chargers {
        if charger.ChargerID == chargerID {
            // fmt.Printf("Charger Status: %d %d \n", charger.UpTime, charger.DownTime)
            var tot_time uint64 = charger.UpTime + charger.DownTime

            // add uptime and downtime to struct
            return charger.UpTime, tot_time
        }
    }

    // no data is 0 according to problem statement
    return 0, 0
}

