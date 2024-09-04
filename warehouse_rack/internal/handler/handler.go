package handler

import (
	"bufio"
	"fadlinux/superapp/warehouse_rack/internal/repository"
	usecases "fadlinux/superapp/warehouse_rack/internal/usecase"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NewWarehouseHandler(input *os.File) {
	scanner := bufio.NewScanner(input)
	var w *repository.WarehouseRepository
	var usecase usecases.WarehouseUsecase

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		command := parts[0]

		switch command {

		case "create_warehouse_rack":
			if len(parts) != 2 {
				fmt.Println("Invalid input")
				continue
			}

			size, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid size")
				continue
			}

			w = repository.NewWarehouseRepository(size)
			usecase = *usecases.NewWarehouseUsecase(repository.WarehouseRepository{
				Slots:    w.Slots,
				Products: w.Products,
			})

			fmt.Printf("Created a warehouse rack with %d slots\n", size)

		default:
			fmt.Println("Invalid command")
		}
	}
}
