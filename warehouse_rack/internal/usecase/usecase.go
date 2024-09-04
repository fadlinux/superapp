package usecases

import (
	"errors"
	"fadlinux/superapp/warehouse_rack/internal/model"
	"fadlinux/superapp/warehouse_rack/internal/repository"
	"fmt"
	"strings"
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

func (uc *WarehouseUsecase) RemoveRack(slotNo int) error {
	if _, ok := uc.repo.Products[slotNo]; ok {
		delete(uc.repo.Products, slotNo)
		uc.repo.Slots[slotNo-1] = 0
		return nil
	}
	return errors.New("slot is empty")
}

func (uc *WarehouseUsecase) PrintStatus() string {
	var result []string
	for i, occupied := range uc.repo.Slots {
		if occupied == 1 {
			product := uc.repo.Products[i+1]
			result = append(result, fmt.Sprintf("%d %s %s", product.SlotNo, product.SKU, product.ExpDate))
		}
	}
	return strings.Join(result, "\n\n")
}
