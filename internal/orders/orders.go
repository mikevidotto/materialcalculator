package orders

import (
    "fmt"
    "materialcalculator/internal/materialcalculator"
)
/*
        name: formData.get("name").trim(),
        email: formData.get("email").trim(),
        phone: formData.get("phone").trim(),
        length: Number(formData.get("length")),
        width: Number(formData.get("width")),
        height: Number(formData.get("height")),
        notes: formData.get("notes").trim()
*/
type OrderData struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Phone string `json:"phone"`
    Length int `json:"length"`
    Width int `json:"width"`
    Height int `json:"height"`
    RoofType string `json:"roofType"`
    Notes string `json:"notes"`
    LumberData materialcalculator.LumberData `json:"lumberData"`
}

func NewOrder(order OrderData) (OrderData, error) {
    if order.Length <= 0 || order.Width <= 0 || order.Height <= 0 {
        return order, fmt.Errorf("length, width, and height must be positive")
    }
    if order.RoofType != "gable" && order.RoofType != "lean-to" {
        return order, fmt.Errorf("roofType must be %q or %q", "gable", "lean-to")
    }

    //convert feet to inches
    order.Length = order.Length * 12
    order.Width = order.Width * 12
    order.Height = order.Height * 12

    shedData := materialcalculator.ShedData{
        Length: order.Length,
        Width: order.Width,
        Height: order.Height,
        RoofType: order.RoofType,
    }

    order.LumberData = materialcalculator.CalculateShedMaterials(shedData)

    return order, nil
}
