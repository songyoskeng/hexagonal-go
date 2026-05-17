package core

import "errors"

// Represent to business logic core
type orderService struct {
	OrderService
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(order Order) error {
	// Business logic function
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}
