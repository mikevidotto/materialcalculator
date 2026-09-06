package main

/* 
1. write all functions                      [ ]
    ->function1: find studs for base         ->[x]
    ->function2: find studs for all walls    ->[ ]
    ->function3: find studs for roof         ->[ ]
    ->function4: find total floorply         ->[ ]
    ->function5: find total finishedply      ->[ ]
    ->function6: find total roofply          ->[ ]

2. write server/routes                      [ ]
3. write API functino                       [ ] 
*/

import (
    "fmt"
)

const ROOFANGLE = 22.5;

type ShedData struct {
    Length int `json:"length"`
    Width int `json:"width"`
    Height int `json:"height"`
}

type Lumber struct {
    Eight int `json:"eight"`
    Ten int`json:"ten"`
    Twelve int `json:"twelve"`
    Sixteen int `json:"sixteen"`
    Floorply int `json:"floorply"`
    Finishedply int `json:"finishedply"`
    Roofply int `json:"roofply"`
}


func main() {
    fmt.Println("sup")
    fmt.Println("Materials for an 8'x20' shed with height of 8'")
    //CalculateMaterials(8, 20, 8)
    CalculateMaterials(96, 240, 96)
}

//accepts a JSON payload with shed data.
//calculate required materials given the length, width and height in INCHES. not feet
//if shed is 8'x20'x8' (96"x240"x96"), then
// base needs to find least required material to build shed
func CalculateMaterials(length int, width int, height int) {
    //floor base
    fmt.Println("BASE STUDS: ", findBaseStuds(SetLongerShorter(length, width)))
    //wall
    fmt.Println("LONG WALL STUDS: ", findWallStuds(SetLongerShorter(length, width)))
    //roof
}

func SetLongerShorter(num1, num2 int) (longer, shorter int) {
    if num1 > num2 {
        longer = num1
        shorter = num2
    } else {
        longer = num2
        shorter = num1
    }

    return longer, shorter
}

func findMinimum(measurement int) int {
    if measurement >= 192 {
        return 192
    } else if measurement >= 144 {
        return 144
    } else if measurement >= 120 {
        return 120
    } else {
        return 96
    }
}

//find studs assuming 16 inch centers
func findBaseStuds(longer int, shorter int) (map[int]int) {
    studMap := make(map[int]int)
    
    if (longer > 192) {
        fmt.Println("length (", longer, ") longer than 16 feet (192), calculated ", (((longer%192)) / 16) + 1, " studs, and subtracting ", (longer%192), " from total length")
        fmt.Println((longer%192), " feet = ", (longer%192), " inches.")
        fmt.Println((longer%192), " / 16 + 1 = ", ((longer%192) / 16)+1)
        studMap[findMinimum(shorter)]+=(((longer%192)) / 16) +1
        if (findMinimum((longer%192)) / (longer%192) ) >= 2 {
            studMap[findMinimum((longer%192))]+=1
        } else {
            studMap[findMinimum((longer%192))]+=2
        }

        longer -= (longer%192)
    }

    fmt.Println("studMap: ", studMap)

    fmt.Println("# of studs required for ", longer, " feet:")
    //8 * 12 = 96 / 16, 32, 48, 64, 80, 96
    fmt.Println(longer, " feet = ", longer, " inches.")
    fmt.Println(longer, " / 16 = ", (longer / 16))
    fmt.Println("1:", ((longer ) / 16)+1)

    studMap[findMinimum(shorter)] += ((longer ) / 16) + 1
        if (findMinimum(longer) / longer) >= 2 {
            studMap[findMinimum(longer)]+=1
        } else {
            studMap[findMinimum(longer)]+=2
        }

    fmt.Println("studMap: ", studMap)
    return studMap 
}

//find studs assuming 24 inch centers
//find studs for all four walls.
//find studs for long walls (do one and just "x2" it.)
//find studs for shorter walls (remember to subtract 7 inches from total inches)
func findWallStuds(longer int, shorter int) (map[int]int) {
    studMap := make(map[int]int)
    
    if (longer > 192) {
        fmt.Println("length (", longer, ") longer than 16 feet (192), calculated ", (((longer%192)) / 24) + 1, " studs, and subtracting ", (longer%192), " from total length")
        fmt.Println((longer%192), " feet = ", (longer%192), " inches.")
        fmt.Println((longer%192), " / 24 + 1 = ", ((longer%192) / 24)+1)
        studMap[findMinimum(shorter)]+=(((longer%192)) / 24) +1
        if (findMinimum((longer%192)) / (longer%192) ) >= 2 {
            studMap[findMinimum((longer%192))]+=1
        } else {
            studMap[findMinimum((longer%192))]+=2
        }

        longer -= (longer%192)
    }

    fmt.Println("studMap: ", studMap)

    fmt.Println("# of studs required for ", longer, " feet:")
    //8 * 12 = 96 / 24, 32, 48, 64, 80, 96
    fmt.Println(longer, " feet = ", longer, " inches.")
    fmt.Println(longer, " / 24 = ", (longer / 24))
    fmt.Println("1:", ((longer ) / 24)+1)

    studMap[findMinimum(shorter)] += ((longer ) / 24) + 1
        if (findMinimum(longer) / longer) >= 2 {
            studMap[findMinimum(longer)]+=1
        } else {
            studMap[findMinimum(longer)]+=2
        }

    fmt.Println("studMap: ", studMap)
    return studMap 
}
