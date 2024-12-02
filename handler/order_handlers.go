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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *OrdersHandler) GetOrdersByID(w http.ResponseWriter, r *http.Request, id int) {
	orders, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *OrdersHandler) CreateOrders(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	orders, err := h.Repo.JSONdecode(body)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error parsing product data", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(orders); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(orders)
}

func (h *OrdersHandler) UpdateOrders(w http.ResponseWriter, r *http.Request, id int) {
	if id == 0 {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	orders, err := h.Repo.JSONdecode(body)
	if err != nil {
		http.Error(w, "Error parsing product data", http.StatusBadRequest)
		return
	}

	orders.ID = id
	if err := h.Repo.Update(orders); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *OrdersHandler) DeleteOrders(w http.ResponseWriter, r *http.Request, id int) {
	if id == 0 {
		http.Error(w, "Invalid Orders ID", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Order deleted successfully"}`)
}
