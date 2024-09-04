package main

import (
	"fmt"
	"os"

	"fadlinux/superapp/warehouse_rack/internal/handler"
)

func main() {
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		handler.NewWarehouseHandler(file)
	}
}
