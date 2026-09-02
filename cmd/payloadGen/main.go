package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Order struct {
	ID     string `json:"id"`
	Amount string `json:"amount"`
}

func main() {
	count := 100000
	orders := make([]Order, 0, count)

	for i := 1; i <= count; i++ {
		// Generate random amounts between 1.00 and 500.00
		amount := rand.Float64()*499.0 + 1.0
		orders = append(orders, Order{
			ID:     fmt.Sprintf("ord_%04d", i),
			Amount: fmt.Sprintf("%.2f", amount),
		})
	}

	// Create and write directly to payload.json
	file, err := os.Create("payload.json")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Indented for readability

	if err := encoder.Encode(orders); err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}

	fmt.Printf("Successfully generated payload.json with %d items!\n", count)
}
