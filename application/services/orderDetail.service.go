package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
	"NetFarm/shared/models/responses"
)

type OrderDetailService struct {
	OrderDetailRepo irepositories.IMedicineRepository
}

func NewOrderDetailService(OrderDetailRepo irepositories.IOrderDetailRepository) iservices.IOrderDetailService {
	return &OrderDetailService{
		OrderDetailRepo: OrderDetailRepo,
	}
}

func (s *OrderDetailService) CreateOrderDetail(orderDetail *entities.OrderDetail) *common.ErrorResponse {

	createdErr := s.OrderDetailRepo.Create(orderDetail)

	if createdErr != nil {
		return errors.NewOrderDetailUnknownError(createdErr.Error(), "falha ao cadastrar despesa")
	}

	return nil
}

func (s *OrderDetailService) FindAllOrderDetail(pagination *common.Pagination) ([]entities.OrderDetail, *common.ErrorResponse) {

	var orderDetails []entities.OrderDetail

	query := s.OrderDetailRepo.Query()

	query = query.Scopes(extensions.Paginate(&orderDetails, pagination, query))

	findErr := query.Find(&orderDetails).Error

	if findErr != nil {
		return orderDetails, errors.NewOrderDetailUnknownError(findErr.Error(), "falha ao buscar detalhe do pedido")
	}

	pagination.Data = orderDetails

	return orderDetails, nil
}

func (s *OrderDetailService) Calculate(items []entities.OrderItem) responses.OrderDetailResponse {

	var amountTax = 0.0
	var amountPaid = 0.0
	var deliveryAmount = 0.0

	var imposed entities.Imposed
	var expense entities.Expense
	var medicine entities.Medicine
	var orderItems []responses.OrderItemDetailResponse

	query := s.OrderDetailRepo.Query()

	imposedErr := query.First(&imposed)

	if imposedErr != nil {

	}

	expenseErr := query.First(&expense)

	if expenseErr != nil {

	}

	for _, item := range items {

		medicineErr := query.First(&medicine, "Id", item.MedicineId)

		if medicineErr != nil {

		}

		amountPaid += medicine.Price * float64(item.Quantity)

		orderItem := responses.OrderItemDetailResponse{
			Price:        medicine.Price,
			Quantity:     item.Quantity,
			MedicineName: medicine.Name,
			MedicineId:   medicine.Id,
		}

		orderItems = append(orderItems, orderItem)
	}

	amountTax += amountPaid * (imposed.Rate / 100)
	deliveryAmount += amountPaid * (expense.Rate / 100)

	amountPaid = (amountPaid + amountTax) + deliveryAmount

	return responses.OrderDetailResponse{
		AmountPaid:     amountPaid,
		TaxAmount:      amountTax,
		DeliveryAmount: deliveryAmount,
		ImposedId:      imposed.Id,
		ExpenseId:      expense.Id,
		OrderItems:     orderItems,
	}

}
