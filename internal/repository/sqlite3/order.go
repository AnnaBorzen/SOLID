package sqlite3

import (
	"SOLID/internal/repository/model"
	"database/sql"
	"fmt"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{
		db: db,
	}
}

func (r *SQLiteRepo) SaveOrder(order *model.Order) error {
	if order == nil {
		return fmt.Errorf("order is nil")
	}

	_, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		order.Customer, order.Products, order.Total, order.Status,
	)
	if err != nil {
		return err
	}
	return nil
}
