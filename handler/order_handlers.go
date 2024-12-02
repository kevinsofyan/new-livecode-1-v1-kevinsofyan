package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	orders "orders/models"
	"strconv"
	"strings"
)

type OrdersHandler struct {
	Repo *orders.OrdersRepository
}

func (h *OrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	var id int
	if len(parts) > 2 && parts[2] != "" {
		parsedID, err := strconv.Atoi(parts[2])
		if err == nil {
			id = parsedID
		}
	}

	switch r.Method {
	case http.MethodGet:
		if id > 0 {
			h.GetOrdersByID(w, r, id)
		} else {
			h.GetAllOrders(w, r)
		}
	case http.MethodPost:
		h.CreateOrders(w, r)
	case http.MethodPut:
		h.UpdateOrders(w, r, id)
	case http.MethodDelete:
		h.DeleteOrders(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrdersHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.Repo.GetAll()
	if err != nil {
		h.handleError(w, "internal server error", err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *OrdersHandler) GetOrdersByID(w http.ResponseWriter, r *http.Request, id int) {
	orders, err := h.Repo.GetByID(id)
	if err != nil {
		h.handleError(w, "order not found", err, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *OrdersHandler) CreateOrders(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, "error reading request body", err, http.StatusBadRequest)
		return
	}

	var order orders.Orders
	if err := json.Unmarshal(body, &order); err != nil {
		h.handleError(w, "error parsing orders data", err, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if order.BuyerName == "" {
		h.handleError(w, "buyer_name is required", fmt.Errorf("buyer_name is required"), http.StatusBadRequest)
		return
	}
	if order.StoreName == "" {
		h.handleError(w, "store_name is required", fmt.Errorf("store_name is required"), http.StatusBadRequest)
		return
	}
	if order.ItemName == "" {
		h.handleError(w, "item_name is required", fmt.Errorf("item_name is required"), http.StatusBadRequest)
		return
	}
	if order.ItemQty == 0 {
		h.handleError(w, "item_qty is required", fmt.Errorf("item_qty is required"), http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(&order); err != nil {
		h.handleError(w, "internal server error", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrdersHandler) UpdateOrders(w http.ResponseWriter, r *http.Request, id int) {
	if id == 0 {
		h.handleError(w, "invalid order ID", fmt.Errorf("invalid order ID"), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, "error reading request body", err, http.StatusBadRequest)
		return
	}

	orders, err := h.Repo.JSONdecode(body)
	if err != nil {
		h.handleError(w, "error parsing order data", err, http.StatusBadRequest)
		return
	}

	orders.ID = id
	if err := h.Repo.Update(orders); err != nil {
		h.handleError(w, "internal server error", err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *OrdersHandler) DeleteOrders(w http.ResponseWriter, r *http.Request, id int) {
	if id == 0 {
		h.handleError(w, "invalid order ID", fmt.Errorf("invalid order ID"), http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		h.handleError(w, "internal server error", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Order deleted successfully"}`)
}

func (h *OrdersHandler) handleError(w http.ResponseWriter, message string, err error, statusCode int) {
	log.Print(err)
	w.WriteHeader(statusCode)
	response := map[string]interface{}{
		"message": message,
	}
	if statusCode == http.StatusInternalServerError {
		response["detail"] = err.Error()
	}
	json.NewEncoder(w).Encode(response)
}
