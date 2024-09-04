package handler

import (
	"bufio"
	"fadlinux/superapp/warehouse_rack/internal/repository"
	usecases "fadlinux/superapp/warehouse_rack/internal/usecase"
	"fmt"
	"os"
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

		default:
			fmt.Println("Invalid command")
		}
	}
}
