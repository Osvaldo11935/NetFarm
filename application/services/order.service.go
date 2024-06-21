package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type OrderService struct {
	orderRepo irepositories.IOrderRepository
}

func NewOrderService(orderRepo irepositories.IOrderRepository) iservices.IOrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) CreateOrder(order *entities.Order) *common.ErrorResponse {

	createdErr := s.orderRepo.Create(order)

	if createdErr != nil {
		return errors.NewOrderUnknownError(createdErr.Error(), "falha ao cadastrar pedido")
	}

	return nil
}

func (s *OrderService) FindAllOrder(pagination *common.Pagination) ([]entities.Order, *common.ErrorResponse) {
	var orders []entities.Order

	query := s.orderRepo.Query()

	query = query.Scopes(extensions.Paginate(&orders, pagination, query))

	findErr := query.Find(&orders).Error

	if findErr != nil {
		return orders, errors.NewOrderUnknownError(findErr.Error(), "falha ao buscar pedido")
	}

	pagination.Data = orders

	return orders, nil
}

func (s *OrderService) FindOrderByUserId(userId string, pagination *common.Pagination) ([]entities.Order, *common.ErrorResponse) {
	var orders []entities.Order

	query := s.orderRepo.Query().Where("UserId", userId)

	query = query.Scopes(extensions.Paginate(&orders, pagination, query))

	findErr := query.Find(&orders).Error

	if findErr != nil {
		return orders, errors.NewOrderUnknownError(findErr.Error(), "falha ao buscar pedido")
	}

	pagination.Data = orders

	return orders, nil
}

func (s *OrderService) FindOrderById(orderId string) (entities.Order, *common.ErrorResponse) {

	var order entities.Order

	findErr := s.orderRepo.Query().First(&order, "Id", orderId).Error

	if findErr != nil {
		return order, errors.NewOrderUnknownError(findErr.Error(), "falha ao buscar pedido")
	}

	return order, nil
}

func (s *OrderService) UpdateOrder(orderId string, statusId *string) *common.ErrorResponse {

	order, err := s.FindOrderById(orderId)

	if err != nil {
		return err
	}

	order.UpdateStatus(statusId)

	updatedErr := s.orderRepo.Update(order)

	if updatedErr != nil {
		return errors.NewOrderUnknownError(updatedErr.Error(), "falha ao atualizar pedido")
	}

	return nil
}

func (s *OrderService) DeleteOrder(OrderId string) *common.ErrorResponse {

	Order, err := s.FindOrderById(OrderId)

	if err != nil {
		return err
	}

	deletedErr := s.orderRepo.Delete(Order)

	if deletedErr != nil {
		return errors.NewOrderUnknownError(deletedErr.Error(), "falha ao deletar pedido")
	}
	return nil
}
