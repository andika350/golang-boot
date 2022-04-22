package controllers

import (
	"encoding/json"
	"net/http"
)

// GetOrders godoc

// @Summary Get details of all orders

// @Description Get details of all orders

// @Tags orders

// @Accept json

// @Produce json

// @Success 200{array} Order

// @Router /orders[get]

var orders []string

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}