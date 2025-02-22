/**  _____ _           _        _        _____          
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
    "os"
    "bufio"
    "strings"
    "strconv"
)

func ParseID(fields []string, line string) (stationID uint32) {
    stationID64, err := strconv.ParseUint(fields[0], 10, 32)
    if err != nil {
        fmt.Printf("ERROR: invalid station ID in %s: %v", line, err)
        os.Exit(1)
    }

    // ParseUint returns a 64 by default
    return uint32(stationID64)
}

func main() {
    fmt.Println("Electric Era Interview Problem")
    cm := NewChargingMonitor()

    if len(os.Args) != 2 {
        fmt.Println("ERROR: Usage: go run main.go ChargingStation.go <input_file.txt>")
        os.Exit(1)
    }
    var in_file string = os.Args[1]
    // fmt.Println(in_file)

    // read in lines from file until empty line
    file, err := os.Open(in_file)
    if err != nil {
        fmt.Printf("ERROR: Failed to open %s, %v\n", in_file, err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var curSel string
    // maintain charger/station ID association
    //var charToStat map[uint32]uint32
    // charToStat = make(map[uint32]uint32)

    for (scanner.Scan()) {
        line := strings.TrimSpace(scanner.Text())
        // so you can comment in files with the # prefix
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
            curSel = line[1 : len(line)-1] // Remove brackets
            continue
        }

        // prepare input fields
        fields := strings.Fields(line) // Split by whitespace
        if len(fields) == 0 {
            continue
        }

        switch curSel {
        case "Stations":

            // can have a station with no charger
            if len(fields) < 1 {
                fmt.Printf("ERROR: invalid station line: %s", line)
                os.Exit(1)
            }

            // get the station ID
            stationID := ParseID(fields, line)

            cm.AddStation(stationID)

            for _, field := range fields[1:] {
                chargerID64, err := strconv.ParseUint(field, 10, 32)
                if err != nil {
                    fmt.Printf("ERROR: invalid charger ID in %s: %v", line, err)
                    os.Exit(1)
                }
                chargerID := uint32(chargerID64)

                cm.AddCharger(stationID, chargerID, 0, 0)
                // charToStat[chargerID] = stationID
            }

        case "Charger Availability Reports":
            if len(fields) != 4 {
                fmt.Printf("ERROR: invalid report line: %s", line)
                os.Exit(1)
            }
            
            chargerID := ParseID(fields, line)
            startTime, err := strconv.ParseUint(fields[1], 10, 64)
            if err != nil {
                fmt.Printf("ERROR: invalid start time in %s: %v", line, err)
                os.Exit(1)
            }
            endTime, err := strconv.ParseUint(fields[2], 10, 64)
            if err != nil {
                fmt.Printf("ERROR: invalid end time in %s: %v", line, err)
                os.Exit(1)
            }
            available, err := strconv.ParseBool(fields[3])
            if err != nil {
                fmt.Printf("ERROR: invalid availability in %s: %v", line, err)
                os.Exit(1)
            }

            if available {
                cm.AddCharger(cm.GetStationID(chargerID), chargerID, endTime - startTime, 0)
            } else {
                cm.AddCharger(cm.GetStationID(chargerID), chargerID, 0, endTime - startTime)
            }
        }

    }

    // Phase 2: writing out to file
    var out_filename string = "out.txt"

    out_file, err := os.OpenFile(out_filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("ERROR: Failed to open %s, %v\n", out_filename, err)
        os.Exit(1)
    }
    defer out_file.Close()

    stationIDs := cm.ListStations()

    for _, sid := range stationIDs {
        var uptime uint64 = cm.CalcStationUptime(sid)

        line := fmt.Sprintf("%d %d\n", sid, uptime)
        fmt.Print(line)
        _, err = out_file.WriteString(line)
    }

}

