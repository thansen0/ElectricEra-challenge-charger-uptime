package main

import (
//    "fmt"
    "sort"
)

type Charger struct {
    ChargerID uint32
    UpTime uint64
    DownTime uint64
}

type Station struct {
    StationID uint32
    Chargers []Charger
}

type ChargingMonitor struct {
    Stations map[uint32]Station
    ChargerToStat map[uint32]uint32
}

// Creates nwe charging station map
func NewChargingMonitor() *ChargingMonitor {
    return &ChargingMonitor{
        Stations: make(map[uint32]Station),
        ChargerToStat: make(map[uint32]uint32),
    }
}

// AddStation adds a new station if it doesn’t exist
func (cm *ChargingMonitor) AddStation(stationID uint32) {
    if _, exists := cm.Stations[stationID]; !exists {
        cm.Stations[stationID] = Station{
            StationID: stationID,
            Chargers:  []Charger{},
        }
    }
}

// AddCharger adds a charger to a station
func (cm *ChargingMonitor) AddCharger(stationID uint32, chargerID uint32, uptime uint64, downtime uint64) {
    station, exists := cm.Stations[stationID]
    if !exists {
        cm.AddStation(stationID)
        station = cm.Stations[stationID]
    }

    // insert into reverse lookup as well
    cm.ChargerToStat[chargerID] = stationID

    for charger_index, charger := range station.Chargers {
        if charger.ChargerID == chargerID {
            // add uptime and downtime to struct
            charger.UpTime += uptime
            charger.DownTime += downtime
            // fmt.Printf("Found charger %d %d \n", charger.UpTime, charger.DownTime)
            // write to original pointer
            cm.Stations[stationID].Chargers[charger_index] = charger
            return
        }
    }

    // no charger exists, add one
    station.Chargers = append(station.Chargers, Charger{
        ChargerID: chargerID,
        UpTime:    uptime,
        DownTime:  downtime,
    })
    cm.Stations[stationID] = station
}

func (cm *ChargingMonitor) GetStationID(chargerID uint32) uint32 {
    return cm.ChargerToStat[chargerID]
}

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

func (cm *ChargingMonitor) CalcStationUptime(stationID uint32) uint64 {
    chargers, exists := cm.ListChargers(stationID)    
    if !exists {
        // doesn't exist, so 0 according to problem statement
        return 0
    }

    var tot_time uint64 = 0
    var up_time uint64 = 0
    for _, c := range chargers {
        u, t := cm.CalcUptime(stationID, c.ChargerID)
        tot_time += t
        up_time += u
    }

    if tot_time == 0 {
        // doesn't exist, so 0 according to problem statement
        return tot_time
    }

    return (up_time * 100) / tot_time
}

func (cm *ChargingMonitor) CalcUptime(stationID uint32, chargerID uint32) (uint64, uint64) {
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

