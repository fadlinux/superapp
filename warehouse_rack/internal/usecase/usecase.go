package usecases

import (
	"errors"
	"fadlinux/superapp/warehouse_rack/internal/model"
	"fadlinux/superapp/warehouse_rack/internal/repository"
)

type WarehouseUsecase struct {
	repo repository.WarehouseRepository
}

func NewWarehouseUsecase(repo repository.WarehouseRepository) *WarehouseUsecase {
	return &WarehouseUsecase{repo: repo}
}

func (uc *WarehouseUsecase) CreateRack(sku string, expDate string) (int, error) {
	for i := 0; i < len(uc.repo.Slots); i++ {
		if uc.repo.Slots[i] == 0 {
			uc.repo.Slots[i] = 1
			uc.repo.Products[i+1] = model.Product{SKU: sku, ExpDate: expDate, SlotNo: i + 1}
			return i + 1, nil
		}
	}
	return 0, errors.New("rack is full")
}
