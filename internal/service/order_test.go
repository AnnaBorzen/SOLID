package service

import (
	"SOLID/internal/repository/model"
	"testing"
)

// Мок RepositoryWriter
type MockRepositoryWriter struct {
	SaveOrderFunc func(order *model.Order) error
}

// Мок Notifier
type MockNotifier struct {
	SendFunc func(customer string)
}

func (m *MockRepositoryWriter) SaveOrder(order *model.Order) error {
	if m.SaveOrderFunc != nil {
		return m.SaveOrderFunc(order)
	}
	return nil // По умолчанию успех
}

func (m *MockNotifier) Send(customer string) {
	if m.SendFunc != nil {
		m.SendFunc(customer)
	}
}

func TestServiceOrder_Success(t *testing.T) {
	mockRepo := &MockRepositoryWriter{}
	mockNotifier := &MockNotifier{}

	service := NewOrderService(mockRepo, mockNotifier)
	order := &model.Order{
		ID:       21,
		Customer: "test",
		Products: "test",
		Total:    12,
		Status:   "test",
	}

	err := service.CreateOrder(order)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}
