package materialcalculator

import "testing"

// TestCalculateShedMaterials_ExactMultipleOfSixteenFeet guards against a
// divide-by-zero panic that occurred whenever the longer wall dimension
// was an exact multiple of 192" (16ft), such as a 32ft or 48ft long shed.
func TestCalculateShedMaterials_ExactMultipleOfSixteenFeet(t *testing.T) {
	cases := []ShedData{
		{Length: 384, Width: 96, Height: 96, RoofType: "gable"},
		{Length: 96, Width: 384, Height: 96, RoofType: "gable"},
		{Length: 384, Width: 96, Height: 96, RoofType: "lean-to"},
		{Length: 576, Width: 120, Height: 96, RoofType: "gable"},
	}

	for _, shed := range cases {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("CalculateShedMaterials(%+v) panicked: %v", shed, r)
				}
			}()
			CalculateShedMaterials(shed)
		}()
	}
}

func TestCalculateShedMaterials_ProducesMaterials(t *testing.T) {
	shed := ShedData{Length: 240, Width: 96, Height: 96, RoofType: "gable"}
	lumber := CalculateShedMaterials(shed)

	if lumber.Floorply == 0 {
		t.Errorf("expected floor plywood sheets to be calculated, got 0")
	}
	if lumber.Finishedply == 0 {
		t.Errorf("expected finished plywood sheets to be calculated, got 0")
	}
	if lumber.Roofply == 0 {
		t.Errorf("expected roof plywood sheets to be calculated, got 0")
	}
	if lumber.Eight+lumber.Ten+lumber.Twelve+lumber.Sixteen == 0 {
		t.Errorf("expected some studs to be calculated, got none")
	}
}
