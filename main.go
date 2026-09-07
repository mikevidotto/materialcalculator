package main

/* 
1. write all functions                      [x]
    ->function1: find studs for base         ->[x]
    ->function2: find studs for all walls    ->[x]
    ->function3: find studs for roof         ->[x]
    ->function4: find total floorply         ->[x]
    ->function5: find total finishedply      ->[x]
    ->function6: find total roofply          ->[x]

2. write server/routes                      [x]
3. write API functino                       [x] 
*/

import (
    "materialcalculator/internal/server"
)

//accepts a JSON payload with shed data.
//calculate required materials given the length, width and height in INCHES. not feet
//if shed is 8'x20'x8' (96"x240"x96"), then
// base needs to find least required material to build shed
func main() {
    server.StartServer()
}

