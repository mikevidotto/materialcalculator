package main

/* 
1. write function     []
2. write server       []
3. write API functino []
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
    OSBply int `json:"osbply"`
}


func main() {
    fmt.Println("sup")
    fmt.Println("Materials for an 8'x20' shed with height of 8'")
    CalculateMaterials(8, 20, 8)
}

//accepts a JSON payload with shed data.
//calculate required materials given the length, width and height in feet
//if shed is 8x20, then
// base needs to find least required material to build shed
func CalculateMaterials(length int, width int, height int) {
    longer := -1
    shorter := -1

    //floor base
    fmt.Println("longer: ", longer, "\nshorter: ", shorter)
    
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
    if measurement >= 16 {
        return 16
    } else if measurement >= 12 {
        return 12
    } else if measurement >= 10 {
        return 10
    } else {
        return 8
    }
}

//find studs assuming 16 inch centers
func findBaseStuds(longer int, shorter int) (map[int]int) {
    studMap := make(map[int]int)
    
    if (longer > 16) {
        fmt.Println("length (", longer, ") longer than 16 feet, calculated ", (((longer%16)*12) / 16) + 1, " studs, and subtracting ", (longer%16), " from total length")
        fmt.Println((longer%16), " feet = ", (longer%16)*12, " inches.")
        fmt.Println((longer%16)*12, " / 16 + 1 = ", ((longer%16)*12 / 16)+1)
        studMap[findMinimum(shorter)]+=(((longer%16)*12) / 16) +1
        if (findMinimum((longer%16)) / (longer%16) ) >= 2 {
            studMap[findMinimum((longer%16))]+=1
        } else {
            studMap[findMinimum((longer%16))]+=2
        }

        longer -= (longer%16)
    }

    fmt.Println("studMap: ", studMap)

    fmt.Println("# of studs required for ", longer, " feet:")
    //8 * 12 = 96 / 16, 32, 48, 64, 80, 96
    fmt.Println(longer, " feet = ", longer*12, " inches.")
    fmt.Println(longer*12, " / 16 = ", (longer*12 / 16))
    fmt.Println("1:", ((longer * 12) / 16)+1)

    studMap[findMinimum(shorter)] += ((longer * 12) / 16) + 1
        if (findMinimum(longer) / longer) >= 2 {
            studMap[findMinimum(longer)]+=1
        } else {
            studMap[findMinimum(longer)]+=2
        }

    fmt.Println("studMap: ", studMap)
    return studMap 
}

//find studs assuming 24 inch centers
//find studs for long walls first, then shorter walls.
func findWallStuds(longer int, shorter int) (map[int]int) {
    studMap := make(map[int]int)
    
    if (longer > 24) {
        fmt.Println("length (", longer, ") longer than 16 feet, calculated ", (((longer%24)*12) / 24) + 1, " studs, and subtracting ", (longer%24), " from total length")
        fmt.Println((longer%24), " feet = ", (longer%24)*12, " inches.")
        fmt.Println((longer%24)*12, " / 24 + 1 = ", ((longer%24)*12 / 24)+1)
        studMap[findMinimum(shorter)]+=(((longer%24)*12) / 24) +1
        if (findMinimum((longer%24)) / (longer%24) ) >= 2 {
            studMap[findMinimum((longer%24))]+=1
        } else {
            studMap[findMinimum((longer%24))]+=2
        }

        longer -= (longer%24)
    }

    fmt.Println("studMap: ", studMap)

    fmt.Println("# of studs required for ", longer, " feet:")
    //8 * 12 = 96 / 24, 32, 48, 64, 80, 96
    fmt.Println(longer, " feet = ", longer*12, " inches.")
    fmt.Println(longer*12, " / 24 = ", (longer*12 / 24))
    fmt.Println("1:", ((longer * 12) / 24)+1)

    studMap[findMinimum(shorter)] += ((longer * 12) / 24) + 1
        if (findMinimum(longer) / longer) >= 2 {
            studMap[findMinimum(longer)]+=1
        } else {
            studMap[findMinimum(longer)]+=2
        }

    fmt.Println("studMap: ", studMap)
    return studMap 
}
