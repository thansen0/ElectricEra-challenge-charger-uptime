package main

import (
    "fmt"
)

type Charger struct {
    ChargerID uint32
    UpTime int
    DownTime int
}

type Station struct {
    StationID uint32
    Chargers []Charger
}

type ChargingMonitor struct {
    Stations map[uint32]Station
}

// Creates nwe charging station map
func NewChargingMonitor() *ChargingMonitor {
    return &ChargingMonitor{
        Stations: make(map[uint32]Station),
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
func (cm *ChargingMonitor) AddCharger(stationID uint32, chargerID uint32, uptime int, downtime int) {
    station, exists := cm.Stations[stationID]
    if !exists {
        cm.AddStation(stationID)
        station = cm.Stations[stationID]
    }

    for charger_index, charger := range station.Chargers {
        if charger.ChargerID == chargerID {
            // add uptime and downtime to struct
            charger.UpTime += uptime
            charger.DownTime += downtime
            fmt.Printf("Found charger %d %d \n", charger.UpTime, charger.DownTime)
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

// ListChargers returns all chargers for a station
func (cm *ChargingMonitor) ListChargers(stationID uint32) ([]Charger, bool) {
    station, exists := cm.Stations[stationID]
    if !exists {
        return nil, false
    }
    return station.Chargers, true
}
