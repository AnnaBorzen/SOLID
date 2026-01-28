package repository

import "SOLID/internal/repository/model"

type RepositoryWriter interface {
	SaveOrder(order *model.Order) error
}

type Notifier interface {
	Send(customer string)
}
