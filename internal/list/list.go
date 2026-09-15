package list

import (
	"fmt"
	"materialcalculator/internal/materialcalculator"
    "os"
)

func CreateListFile(data materialcalculator.LumberData) error {
    body := fmt.Sprintf("Studs\n\n2x4x8: %d\n2x4x10: %d\n2x4x12: %d\n2x4x16: %d\n\nSheets\n\n3/4\": %d\n1/2\": %d\nSmartSide: %d\n", data.Eight, data.Ten, data.Twelve, data.Sixteen, data.Floorply, data.Roofply, data.Finishedply)
    fmt.Println(body)
    name := "list.txt"

    file, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(body))
	if err != nil {
		return err
	}

    return nil
}
