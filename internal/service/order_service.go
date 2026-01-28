package service

import (
	"SOLID/internal/repository"
	"SOLID/internal/repository/model"
)

type OrderService struct {
	repo     repository.RepositoryWriter
	notifier repository.Notifier
}

func NewOrderService(repo repository.RepositoryWriter, notifier repository.Notifier) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: notifier,
	}
}

// CreateOrder создание клиента, с отправкой уведомления
func (s *OrderService) CreateOrder(order *model.Order) error {
	if err := s.repo.SaveOrder(order); err != nil {
		return err
	}

	s.notifier.Send("Тестовый клиент")
	return nil
}
