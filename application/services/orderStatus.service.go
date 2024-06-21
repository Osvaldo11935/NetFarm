package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type OrderStatusService struct {
	OrderStatusRepo irepositories.IOrderStatusRepository
}

func NewOrderStatusService(OrderStatusRepo irepositories.IOrderStatusRepository) iservices.IOrderStatusService {
	return &OrderStatusService{
		OrderStatusRepo: OrderStatusRepo,
	}
}

func (s *OrderStatusService) CreateOrderStatus(OrderStatus *entities.OrderStatus) *common.ErrorResponse {

	createdErr := s.OrderStatusRepo.Create(OrderStatus)

	if createdErr != nil {
		return errors.NewOrderStatusUnknownError(createdErr.Error(), "falha ao cadastrar estado")
	}

	return nil
}

func (s *OrderStatusService) FindAllOrderStatus(pagination *common.Pagination) ([]entities.OrderStatus, *common.ErrorResponse) {
	var orderStatus []entities.OrderStatus

	query := s.OrderStatusRepo.Query()

	query = query.Scopes(extensions.Paginate(&orderStatus, pagination, query))

	findErr := query.Find(&orderStatus).Error

	if findErr != nil {
		return orderStatus, errors.NewOrderStatusUnknownError(findErr.Error(), "falha ao buscar estado")
	}

	pagination.Data = orderStatus

	return orderStatus, nil
}

func (s *OrderStatusService) FindOrderStatusById(OrderStatusId string) (entities.OrderStatus, *common.ErrorResponse) {

	var orderStatus entities.OrderStatus

	findErr := s.OrderStatusRepo.Query().First(&orderStatus, "Id", OrderStatusId).Error

	if findErr != nil {
		return orderStatus, errors.NewOrderStatusUnknownError(findErr.Error(), "falha ao buscar estado")
	}

	return orderStatus, nil
}

func (s *OrderStatusService) UpdateOrderStatus(OrderStatusId string, name *string, description *string) *common.ErrorResponse {

	orderStatus, err := s.FindOrderStatusById(OrderStatusId)

	if err != nil {
		return err
	}

	orderStatus.Update(name, description)

	updatedErr := s.OrderStatusRepo.Update(orderStatus)

	if updatedErr != nil {
		return errors.NewOrderStatusUnknownError(updatedErr.Error(), "falha ao atualizar estado")
	}

	return nil
}

func (s *OrderStatusService) DeleteOrderStatus(OrderStatusId string) *common.ErrorResponse {

	orderStatus, err := s.FindOrderStatusById(OrderStatusId)

	if err != nil {
		return err
	}

	deletedErr := s.OrderStatusRepo.Delete(orderStatus)

	if deletedErr != nil {
		return errors.NewOrderStatusUnknownError(deletedErr.Error(), "falha ao deletar estado")
	}

	return nil
}
