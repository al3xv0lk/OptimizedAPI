package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// RegisterRoutes registers all application endpoints onto the provided ServeMux.
func RegisterRoutes(mux *http.ServeMux) {
	// Health check
	mux.HandleFunc("GET /health", HealthHandler)
	mux.HandleFunc("POST /base", BaseHandler)
	mux.HandleFunc("POST /optimized", OptimizedHandler)
}

type Order struct {
	ID     string `json:"id"`
	Amount string `json:"amount"`
}

// Logic - Base
func ProcessOrdersBase(orders []Order) (float64, []string) {
	var total float64
	var processedIDs []string // Cresce dinamicamente (aloca várias vezes)

	for _, o := range orders {
		cleanAmount := strings.TrimSpace(o.Amount)
		val, _ := strconv.ParseFloat(cleanAmount, 64)
		total += val
		processedIDs = append(processedIDs, o.ID) // Repassa o ponteiro, não cria nova string
	}

	return total, processedIDs
}

// Logic - Optimized
func ProcessOrdersOptimized(orders []Order) (float64, []string) {
	n := len(orders)
	var total float64

	// ALOCAÇÃO ÚNICA: cria o array base com o tamanho exato na memória
	processedIDs := make([]string, 0, n)

	for i := 0; i < n; i++ {
		val, _ := strconv.ParseFloat(orders[i].Amount, 64)
		total += val
		processedIDs = append(processedIDs, orders[i].ID)
	}

	return total, processedIDs
}

// Handlers
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	var orders []Order
	if err := json.NewDecoder(r.Body).Decode(&orders); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	total, processedIDs := ProcessOrdersBase(orders)

	resp := map[string]any{
		"total":           total,
		"processed_count": len(processedIDs),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	var orders []Order
	if err := json.NewDecoder(r.Body).Decode(&orders); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	total, processedIDs := ProcessOrdersBase(orders)

	resp := map[string]any{
		"total":           total,
		"processed_count": len(processedIDs),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func OptimizedHandler(w http.ResponseWriter, r *http.Request) {
	var orders []Order
	if err := json.NewDecoder(r.Body).Decode(&orders); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	total, processedIDs := ProcessOrdersOptimized(orders)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"total":` + strconv.FormatFloat(total, 'f', 2, 64) + `,"processed_count":` + strconv.Itoa(len(processedIDs)) + `}`))
}
