package main

import (
	"SOLID/internal/notification"
	"SOLID/internal/repository/model"
	"SOLID/internal/repository/sqlite3"
	"SOLID/internal/service"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Обычно нужны миграции, но задание не предполагало их создание
func migrationCreateTable(db *sql.DB) {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        customer TEXT NOT NULL,
        products TEXT NOT NULL,
        total REAL NOT NULL,
        status TEXT NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	migrationCreateTable(db)

	orderTest := &model.Order{
		ID:       21,
		Customer: "string",
		Products: "string",
		Total:    12,
		Status:   "string",
	}

	repo := sqlite3.NewSQLiteRepo(db)

	emailSender := &notification.EmailSender{}
	smsSender := &notification.SMSSender{}

	orderService := service.NewOrderService(repo, emailSender)
	orderServiceSms := service.NewOrderService(repo, smsSender)

	errOrder := orderService.CreateOrder(orderTest)
	if errOrder != nil {
		log.Fatal("Failed to create order:", errOrder)
	}

	errOrderSMS := orderServiceSms.CreateOrder(orderTest)
	if errOrderSMS != nil {
		log.Fatal("Failed to create order:", errOrderSMS)
	}

}
