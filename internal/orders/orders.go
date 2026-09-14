package orders

import (
    "fmt"
    "materialcalculator/internal/materialcalculator"
    "materialcalculator/internal/list"
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
    RoofType string `json:"rooftype"`
    Notes string `json:"notes"`
    LumberData materialcalculator.LumberData
}

func NewOrder(order OrderData) error {
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

    err := list.CreateListFile(order.LumberData)
    if err != nil {
        return err
    }
    fmt.Printf("lumber data: %#v", order.LumberData)

    return nil
}
