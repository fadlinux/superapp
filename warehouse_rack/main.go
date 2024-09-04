package main

import (
	"fmt"

	mWarehouse "superapp/warehouse_rack/internal/model"
)

func main() {

	//init model
	nData := 3
	tempModel := make(map[int]mWarehouse.WarehouseRackData, nData)

	tempModel[1] = mWarehouse.WarehouseRackData{
		SlotNo:      1,
		SKU:         "sku",
		ExpiredData: "xxx",
	}

	tempModel[2] = mWarehouse.WarehouseRackData{
		SlotNo:      2,
		SKU:         "2sku",
		ExpiredData: "2xxx",
	}

	tempModel[3] = mWarehouse.WarehouseRackData{
		SlotNo:      3,
		SKU:         "3sku",
		ExpiredData: "4xxx",
	}

	for x, y := range tempModel {
		fmt.Println("x,y : ", x, y)

		if x >= nData {
			break
		}

	}
	fmt.Println("tempModel : ", tempModel)

}
