package materialcalculator

import (
    "math"
)

type ShedData struct {
    Length int `json:"length"`
    Width int `json:"width"`
    Height int `json:"height"`
    RoofType string `json:"roofType"`
}

type LumberData struct {
    Eight int `json:"eight"`
    Ten int`json:"ten"`
    Twelve int `json:"twelve"`
    Sixteen int `json:"sixteen"`
    Floorply int `json:"floorply"`
    Finishedply int `json:"finishedply"`
    Roofply int `json:"roofply"`
}

func CalculateShedMaterials(Shed ShedData) LumberData {
	MaterialMap := make(map[int]int)
    longer, shorter := SetLongerShorter(Shed.Length, Shed.Width)
    
    //floor base
	tempMap := findBaseStuds(longer, shorter)
	for studlength, quantity:= range tempMap {
		MaterialMap[studlength] += quantity
	}
    //fmt.Printf("%#v\n", tempMap)

    switch Shed.RoofType {
        //gable roof
        case "gable":
            //wall
            tempMap = findGableWallStuds(longer, shorter, Shed.Height)
            for studlength, quantity:= range tempMap {
                MaterialMap[studlength] += quantity
            }

            //roof
			tempMap = findGableRoofStuds(longer, shorter)
			for studlength, quantity:= range tempMap {
				MaterialMap[studlength] += quantity
	        }

            //roof sheets
            tempMap = findGableRoofPly(longer, shorter)
            for sheetIndex, quantity := range tempMap {
                MaterialMap[sheetIndex] += quantity
            }
        //leanto roof
        case "lean-to":
            slope := 5.0/12.0
            rise := slope * float64(Shed.Height)
            TallStudHeight := Shed.Height + int(rise)

            //wall
            tempMap = findLeanToWallStuds(longer, shorter, Shed.Height, TallStudHeight)
            for studlength, quantity:= range tempMap {
                MaterialMap[studlength] += quantity
            }

            //roof
			tempMap = findLeanToRoofStuds(longer, shorter)
			for studlength, quantity:= range tempMap {
				MaterialMap[studlength] += quantity
	        }
        
            //roof sheets
            tempMap = findLeanToRoofPly(longer, shorter)
            for sheetIndex, quantity := range tempMap {
                MaterialMap[sheetIndex] += quantity
            }
            
    }

    //floor sheets 
    tempMap = findFloorPly(longer, shorter)
    for sheetIndex, quantity := range tempMap {
        MaterialMap[sheetIndex] += quantity
    }

    //finished sheets
    tempMap = findFinishedPly(longer, shorter, Shed.Height)
    for sheetIndex, quantity := range tempMap {
        MaterialMap[sheetIndex] += quantity
    }


    Lumber := LumberData {}

    for key, quantity := range MaterialMap {
        switch key {
        case 1:
            Lumber.Floorply = quantity
        case 2:
            Lumber.Finishedply= quantity
        case 3:
            Lumber.Roofply= quantity
        case 96:
            Lumber.Eight = quantity
        case 120:
            Lumber.Ten = quantity
        case 144:
            Lumber.Twelve = quantity
        case 196:
            Lumber.Sixteen = quantity
        }
    }

    return Lumber
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
    if measurement > 144 {
        return 196
    } else if measurement > 120 {
        return 144
    } else if measurement > 96 {
        return 120
    } else {
		return 96
	}
}

//find studs assuming 16 inch centers
func findBaseStuds(longer int, shorter int) (map[int]int) {
    studMap := make(map[int]int)
    
    if (longer > 192) {
        remainder := longer % 192
        if remainder != 0 {
            studMap[findMinimum(shorter)]+=(remainder / 16) +1
            if (findMinimum(remainder) / remainder ) >= 2 {
                studMap[findMinimum(remainder)]+=1
            } else {
                studMap[findMinimum(remainder)]+=2
            }
        }
        longer -= remainder
    }

    studMap[findMinimum(shorter)] += ((longer ) / 16) + 1

    if (findMinimum(longer) / longer) >= 2 {
        studMap[findMinimum(longer)]+=1
    } else {
        studMap[findMinimum(longer)]+=2
    }
    return studMap 
}

func findLeanToWallStuds(longer, shorter, height1, height2 int) (map[int]int) {
    studMap := make(map[int]int)
	//each short wall will lose 7 inches on each side since they will join with each existing long wall.
	shorter = shorter-7
    remainderwall := (longer%192)

	//if wall is longer than 16feet, find the materials for the remainder (eg. 20ft wall would need to calculate materials for a 4 foot wall and a 16 foot wall)
    if (longer > 192) {
        if remainderwall != 0 {
            studMap[findMinimum(height1)]+=((remainderwall) / 24) +1
            //fmt.Println((remainderwall / 24) + 1)
            studMap[findMinimum(height2)]+=((remainderwall) / 24) +1
            //fmt.Println((remainderwall / 24) + 1)
            //if the minimum plate required divided by the length needed is greater than 2, we can use one stud for the same plate, so only add one.
            if (findMinimum(remainderwall) / remainderwall ) >= 2 {
                //fmt.Println("2")
                studMap[findMinimum(remainderwall)]+=2
            } else {
                //fmt.Println("4")
                studMap[findMinimum(remainderwall)]+=4
            }
        }
        longer -= remainderwall
    }

	//calculate first wall, then duplicate
    studMap[findMinimum(height1)] += ((longer ) / 24) + 1
    //fmt.Println((longer / 24) +1)
	
	studMap[findMinimum(height2)] += ((longer) / 24) + 1
    //fmt.Println((longer / 24) +1)
	
        if (findMinimum(longer) / longer) >= 2 {
            studMap[findMinimum(longer)]+=2
        //fmt.Println("2")
        } else {
            studMap[findMinimum(longer)]+=4
        //fmt.Println("4")
        }

	//calculate second wall, then duplicate
	studMap[findMinimum(height1)] += int(math.Ceil((float64(shorter) / 24)) +1)
    //fmt.Println((shorter / 24) +1)
    //fmt.Println("math.Ceil:", math.Ceil((float64(shorter) / 24)) +1)
	
	studMap[findMinimum(height1)] += int(math.Ceil((float64(shorter) / 24)) +1)
    //fmt.Println((shorter / 24) +1)
    //fmt.Println("math.Ceil:", math.Ceil((float64(shorter) / 24)) +1)
        if (findMinimum(shorter) / shorter) >= 2 {
            studMap[findMinimum(shorter)]+=2
            //fmt.Println("2")
        } else {
            studMap[findMinimum(shorter)]+=4
            //fmt.Println("4")
        }

    return studMap 
}

//find studs assuming 24 inch centers
//find studs for all four walls.
//find studs for long walls (do one and just "x2" it.)
//find studs for shorter walls (remember to subtract 7 inches from total inches)
func findGableWallStuds(longer, shorter, height int) (map[int]int) {
    studMap := make(map[int]int)
	//each short wall will lose 7 inches on each side since they will join with each existing long wall.
	shorter = shorter-7
    remainderwall := (longer%192)

	//if wall is longer than 16feet, find the materials for the remainder (eg. 20ft wall would need to calculate materials for a 4 foot wall and a 16 foot wall)
    if (longer > 192) {
        if remainderwall != 0 {
            studMap[findMinimum(height)]+=((remainderwall) / 24) +1
            //fmt.Println((remainderwall / 24) + 1)
            studMap[findMinimum(height)]+=((remainderwall) / 24) +1
            //fmt.Println((remainderwall / 24) + 1)
            //if the minimum plate required divided by the length needed is greater than 2, we can use one stud for the same plate, so only add one.
            if (findMinimum(remainderwall) / remainderwall ) >= 2 {
                //fmt.Println("2")
                studMap[findMinimum(remainderwall)]+=2
            } else {
                //fmt.Println("4")
                studMap[findMinimum(remainderwall)]+=4
            }
        }
        longer -= remainderwall
    }

	//calculate first wall, then duplicate
    studMap[findMinimum(height)] += ((longer ) / 24) + 1
    //fmt.Println((longer / 24) +1)
	
	studMap[findMinimum(height)] += ((longer) / 24) + 1
    //fmt.Println((longer / 24) +1)
	
        if (findMinimum(longer) / longer) >= 2 {
            studMap[findMinimum(longer)]+=2
        //fmt.Println("2")
        } else {
            studMap[findMinimum(longer)]+=4
        //fmt.Println("4")
        }

	//calculate second wall, then duplicate
	studMap[findMinimum(height)] += int(math.Ceil((float64(shorter) / 24)) +1)
    //fmt.Println((shorter / 24) +1)
    //fmt.Println("math.Ceil:", math.Ceil((float64(shorter) / 24)) +1)
	
	studMap[findMinimum(height)] += int(math.Ceil((float64(shorter) / 24)) +1)
    //fmt.Println((shorter / 24) +1)
    //fmt.Println("math.Ceil:", math.Ceil((float64(shorter) / 24)) +1)
        if (findMinimum(shorter) / shorter) >= 2 {
            studMap[findMinimum(shorter)]+=2
            //fmt.Println("2")
        } else {
            studMap[findMinimum(shorter)]+=4
            //fmt.Println("4")
        }

    return studMap 
}

func findGableRoofStuds(longer, shorter int) (map[int]int) {
	studMap := make(map[int]int)
    rakeLengthFloat := CalculateGableRoofRakeLength(shorter)
	rakeLengthInt := int(rakeLengthFloat)

	//find number of trusses required for every 16 inches - use longer
	//rakelonger / 16
	studMap[findMinimum(rakeLengthInt * 2)] += longer / 16
	return studMap
}

func findLeanToRoofStuds(longer, shorter int) (map[int]int) {
	studMap := make(map[int]int)
    rakeLengthFloat := CalculateLeanToRoofRakeLength(shorter)
	rakeLengthInt := int(rakeLengthFloat)

	//find number of trusses required for every 16 inches - use longer
	//rakelonger / 16
	studMap[findMinimum(rakeLengthInt * 2)] += longer / 16
	return studMap
}

func CalculateGableRoofRakeLength(length int) float64 {
	//find rake length (eg. 8ft width is 52) - use shorter
	//find run (span / 2)
	run := float64(length) / 2.0
	//find pitch (5/12 or 22.5 degrees)
	pitch := 5.0/12.0
	//find verticalrise (pitch * run)
	verticalRise := float64(pitch * run)
	//find rake length (pythagorean theorem)
	rakeLengthFloat:= math.Hypot(float64(verticalRise), float64(run))
    return rakeLengthFloat
}

func CalculateLeanToRoofRakeLength(length int) float64 {
	//find rake length (eg. 8ft width is 52) - use shorter
	//find run (span / 2)
	run := float64(length)
	//find pitch (5/12 or 22.5 degrees)
	pitch := 5.0/12.0
	//find verticalrise (pitch * run)
	verticalRise := float64(pitch * run)
	//find rake length (pythagorean theorem)
	rakeLengthFloat:= math.Hypot(float64(verticalRise), float64(run))
    return rakeLengthFloat
}

func findFloorPly(length, width int) map[int]int {
    floorPlyMap:= make(map[int]int)
    //all ply comes in 4'x8' (48"x96")
    //for sheds with a width of 8' or less, we simply divide the length by 4 to get the number of sheets
    //1. find the area of the shed.
    shedArea := length * width
    //2. calculate area of sheet
    sheetArea := 48 * 96
    //3. find number of sheets required based on area.
    floorPlyMap[1] += int(math.Ceil(float64(shedArea) / float64(sheetArea)))
    return floorPlyMap 
}

func findFinishedPly(length, width, height int) map[int]int {
    finishedPlyMap := make(map[int]int)
    sheetArea := 48 * 96

    WallArea1 := length * height
    WallArea2 := length * height
    WallArea3 := width* height
    WallArea4 := width* height

    TotalArea := WallArea1 + WallArea2 + WallArea3 + WallArea4

    finishedPlyMap[2] += int(math.Ceil(float64(TotalArea) / float64(sheetArea)))

    return finishedPlyMap
}

func findGableRoofPly(length, width int) map[int]int {
    roofPlyMap := make(map[int]int)
    sheetArea := 48 * 96
    rakeLengthFloat := CalculateGableRoofRakeLength(width)
   

    roofArea := (length * int(rakeLengthFloat)) * 2
   

    roofPlyMap[3] += int(math.Ceil(float64(roofArea) / float64(sheetArea)))

    return roofPlyMap
}

func findLeanToRoofPly(length, width int) map[int]int {
    roofPlyMap := make(map[int]int)
    sheetArea := 48 * 96
    rakeLengthFloat := CalculateLeanToRoofRakeLength(width)

    roofArea := (length * int(rakeLengthFloat)) * 2
   

    roofPlyMap[3] += int(math.Ceil(float64(roofArea) / float64(sheetArea)))

    return roofPlyMap
}
