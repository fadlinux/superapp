package model

type WarehouseData struct {
	Slots    []int
	Products map[int]Product
}

type Product struct {
	SKU     string
	ExpDate string
	SlotNo  int
}
