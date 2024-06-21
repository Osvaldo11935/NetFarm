package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type OrderItemService struct {
	OrderItemRepo irepositories.IOrderItemRepository
}

func NewOrderItemService(OrderItemRepo irepositories.IOrderItemRepository) iservices.IOrderItemService {
	return &OrderItemService{
		OrderItemRepo: OrderItemRepo,
	}
}

func (s *OrderItemService) CreateOrderItem(OrderItem *entities.OrderItem) *common.ErrorResponse {

	createdErr := s.OrderItemRepo.Create(OrderItem)

	if createdErr != nil {
		return errors.NewOrderStatusUnknownError(createdErr.Error(), "falha ao cadastrar itens do pedido")
	}

	return nil
}

func (s *OrderItemService) FindOrderItemByOrderId(orderId string, pagination *common.Pagination) ([]entities.OrderItem, *common.ErrorResponse) {
	var orderItems []entities.OrderItem

	query := s.OrderItemRepo.Query().Where("OrderId", orderId)

	query = query.Scopes(extensions.Paginate(&orderItems, pagination, query))

	findErr := query.Find(&orderItems).Error

	if findErr != nil {
		return orderItems, errors.NewOrderItemUnknownError(findErr.Error(), "falha ao cadastrar itens do pedido")
	}

	pagination.Data = orderItems

	return orderItems, nil
}

func (s *OrderItemService) FindOrderItemById(orderItemId string) (entities.OrderItem, *common.ErrorResponse) {

	var orderItem entities.OrderItem

	findErr := s.OrderItemRepo.Query().First(&orderItem, "Id", orderItemId).Error

	if findErr != nil {
		return orderItem, errors.NewOrderItemUnknownError(findErr.Error(), "falha ao cadastrar itens do pedido")
	}

	return orderItem, nil
}

func (s *OrderItemService) UpdateOrderItem(orderItemId string,
	quantity *int) *common.ErrorResponse {

	orderItem, err := s.FindOrderItemById(orderItemId)

	if err != nil {
		return err
	}

	orderItem.Update(quantity)

	updatedErr := s.OrderItemRepo.Update(orderItem)

	if updatedErr != nil {
		return errors.NewOrderItemUnknownError(updatedErr.Error(), "falha ao atualizar itens do pedido")
	}

	return nil
}

func (s *OrderItemService) DeleteOrderItem(orderItemId string) *common.ErrorResponse {

	orderItem, err := s.FindOrderItemById(orderItemId)

	if err != nil {
		return err
	}

	deletedErr := s.OrderItemRepo.Delete(orderItem)

	if deletedErr != nil {
		return errors.NewOrderItemUnknownError(deletedErr.Error(), "falha ao deletar itens do pedido")
	}

	return nil
}

func (s *OrderItemService) OrderDetail(orderItems []entities.OrderItem) {

}
