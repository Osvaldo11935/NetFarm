package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type OrderFileService struct {
	OrderFileRepo irepositories.IOrderFileRepository
}

func NewOrderFileService(OrderFileRepo irepositories.IOrderFileRepository) iservices.IOrderFileService {
	return &OrderFileService{
		OrderFileRepo: OrderFileRepo,
	}
}

func (s *OrderFileService) CreateOrderFile(OrderFile *entities.OrderFile) *common.ErrorResponse {

	createdErr := s.OrderFileRepo.Create(OrderFile)

	if createdErr != nil {
		return errors.NewOrderFileUnknownError(createdErr.Error(), "falha ao cadastrar ficheiro do pedido")
	}

	return nil
}

func (s *OrderFileService) FindOrderFileByOrderId(orderId string, pagination *common.Pagination) ([]entities.OrderFile, *common.ErrorResponse) {
	var orderFiles []entities.OrderFile

	query := s.OrderFileRepo.Query().Where("OrderId", orderId)

	query = query.Scopes(extensions.Paginate(&orderFiles, pagination, query))

	findErr := query.Find(&orderFiles).Error

	if findErr != nil {
		return orderFiles, errors.NewOrderFileUnknownError(findErr.Error(), "falha ao buscar ficheiro do pedido")
	}

	pagination.Data = orderFiles

	return orderFiles, nil
}

func (s *OrderFileService) FindOrderFileById(orderFileId string) (entities.OrderFile, *common.ErrorResponse) {

	var orderFile entities.OrderFile

	findErr := s.OrderFileRepo.Query().First(&orderFile, "Id", orderFileId).Error

	if findErr != nil {
		return orderFile, errors.NewOrderFileUnknownError(findErr.Error(), "falha ao buscar ficheiro do pedido")
	}

	return orderFile, nil
}

func (s *OrderFileService) DeleteOrderFile(orderFileId string) *common.ErrorResponse {

	orderFile, err := s.FindOrderFileById(orderFileId)

	if err != nil {
		return err
	}

	deletedErr := s.OrderFileRepo.Delete(orderFile)

	if deletedErr != nil {
		return errors.NewOrderFileUnknownError(deletedErr.Error(), "falha ao deleter ficheiro do pedido")
	}

	return nil
}
