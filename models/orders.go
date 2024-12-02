package orders

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type Orders struct {
	ID        int    `json:"id"`
	BuyerName string `json:"buyer_name"`
	StoreName string `json:"store_name"`
	ItemName  int    `json:"item_name"`
	ItemQty   int    `json:"item_qty"`
	CreatedAt int    `json:"created_at"`
}

type OrdersRepository struct {
	DB *sql.DB
}

func NewOrdersRepository(db *sql.DB) *OrdersRepository {
	return &OrdersRepository{DB: db}
}

func (or *OrdersRepository) GetAll() ([]Orders, error) {

	rows, err := or.DB.Query("SELECT id, buyer_name, store_name, item_name, item_qty, created_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Orders
	for rows.Next() {
		var o Orders
		if err := rows.Scan(&o.ID, &o.BuyerName, &o.StoreName, &o.ItemName, &o.ItemQty, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}

func (or *OrdersRepository) GetByID(id int) (*Orders, error) {
	var o Orders
	err := or.DB.QueryRow("SELECT id, buyer_name, store_name, item_name, item_qtym created_at FROM orders where id = ?", id).Scan(&o.ID, &o.BuyerName, &o.StoreName, &o.ItemName, &o.ItemQty, &o.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("Order not found")
	} else if err != nil {
		return nil, err
	}

	return &o, nil
}

func (or *OrdersRepository) Create(o *Orders) error {
	result, err := or.DB.Exec(
		"INSERT INTO products (buyer_name, store_name, item_name, item_qtym created_at) (?, ?, ?, ?, NOW())",
		o.BuyerName, o.StoreName, o.ItemName, o.ItemQty,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	o.ID = int(id)
	return nil
}

func (or *OrdersRepository) Update(o *Orders) error {
	_, err := or.DB.Exec(
		"UPDATE orders SET buyer_name = ?, store_name = ?, item_name = ?, item_qty = ? WHERE id = ?",
		o.BuyerName, o.StoreName, o.ItemName, o.ItemQty, o.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (or *OrdersRepository) Delete(id int) error {
	result, err := or.DB.Exec("DELETE FROM orders WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (or *OrdersRepository) JSONdecode(body []byte) (*Orders, error) {
	var o Orders
	err := json.Unmarshal(body, &o)
	return &o, err
}

func (r *OrdersRepository) JSONencode(p *Orders) ([]byte, error) {
	return json.Marshal(p)
}
