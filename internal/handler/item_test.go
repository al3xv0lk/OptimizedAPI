package handler

import (
	"fmt"
	"testing"
)

func mockOrders(count int) []Order {
	orders := make([]Order, count)
	for i := 0; i < count; i++ {
		orders[i] = Order{
			ID:     fmt.Sprintf("ord_%04d", i),
			Amount: "149.99",
		}
	}
	return orders
}

func BenchmarkProcessOrders_Base(b *testing.B) {
	orders := mockOrders(100000)

	b.ReportAllocs()

	for b.Loop() {
		ProcessOrdersBase(orders)
	}
}

func BenchmarkProcessOrders_Optimized(b *testing.B) {
	orders := mockOrders(100000)

	b.ReportAllocs()

	for b.Loop() {
		ProcessOrdersOptimized(orders)
	}
}
