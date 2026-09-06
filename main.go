package main

/* 
1. write all functions                      [x]
    ->function1: find studs for base         ->[x]
    ->function2: find studs for all walls    ->[x]
    ->function3: find studs for roof         ->[x]
    ->function4: find total floorply         ->[x]
    ->function5: find total finishedply      ->[x]
    ->function6: find total roofply          ->[x]

2. write server/routes                      [ ]
3. write API functino                       [ ] 
*/

import (
    "fmt"
    "materialcalculator/internal/materialcalculator"
    "materialcalculator/internal/server"
)

//accepts a JSON payload with shed data.
//calculate required materials given the length, width and height in INCHES. not feet
//if shed is 8'x20'x8' (96"x240"x96"), then
// base needs to find least required material to build shed
func main() {
    fmt.Println("Materials for an 8'x20' (96\"x240\"x96\") shed with height of 8'")
    Shed := materialcalculator.ShedData {
        Length: 96,
        Width: 240,
        Height: 96,
    }

    Lumber := materialcalculator.CalculateMaterials(Shed)

    fmt.Printf("%#v\n", Lumber)

    server.StartServer()
}

