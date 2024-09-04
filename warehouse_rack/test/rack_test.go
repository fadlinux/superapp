package warehouse

import (
	"fadlinux/superapp/warehouse_rack/internal/repository"
	usecases "fadlinux/superapp/warehouse_rack/internal/usecase"

	"testing"
)

func TestWarehouseRack(t *testing.T) {
	w := repository.NewWarehouseRepository(6)
	usecase := *usecases.NewWarehouseUsecase(repository.WarehouseRepository{
		Slots:    w.Slots,
		Products: w.Products,
	})

	slot, err := usecase.CreateRack("ZG11AQA", "2024-02-28")
	if err != nil || slot != 1 {
		t.Errorf("Expected slot 1, got %d, err: %v", slot, err)
	}

	slot, err = usecase.CreateRack("SD92349WW", "2024-02-28")
	if err != nil || slot != 2 {
		t.Errorf("Expected slot 2, got %d, err: %v", slot, err)
	}

	slot, err = usecase.CreateRack("ZG748WDG", "2024-03-15")
	if err != nil || slot != 3 {
		t.Errorf("Expected slot 3, got %d, err: %v", slot, err)
	}

	slot, err = usecase.CreateRack("KA887YHJ", "2024-01-21")
	if err != nil || slot != 4 {
		t.Errorf("Expected slot 4, got %d, err: %v", slot, err)
	}

	slot, err = usecase.CreateRack("KA888YHH", "2024-04-01")
	if err != nil || slot != 5 {
		t.Errorf("Expected slot 5, got %d, err: %v", slot, err)
	}

	slot, err = usecase.CreateRack("DG8789YH", "2024-03-15")
	if err != nil || slot != 6 {
		t.Errorf("Expected slot 6, got %d, err: %v", slot, err)
	}

	err = usecase.RemoveRack(4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

}
