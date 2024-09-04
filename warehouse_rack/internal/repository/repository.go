package repository

import (
	"fadlinux/superapp/warehouse_rack/internal/model"
)

type WarehouseRepository struct {
	Slots    []int
	Products map[int]model.Product
}

func NewWarehouseRepository(nDataSlots int) *WarehouseRepository {
	return &WarehouseRepository{
		Slots:    make([]int, nDataSlots),
		Products: make(map[int]model.Product),
	}
}
