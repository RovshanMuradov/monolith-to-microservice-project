package intraprocess

import (
	"monolith-t0-microservice-project/pkg/orders/application"
	"monolith-t0-microservice-project/pkg/orders/domain/orders"
)

type OrdersInterface struct {
	service application.OrdersSerivce
}

func NewOrdersInterface(service application.OrdersSerivce) OrdersInterface {
	return OrdersInterface{service}
}

func (p OrdersInterface) MarkOrderAsPaid(orderID string) error {
	return p.service.MarkOrderAsPaid(application.MarkOrderAsPaidCommand{orders.ID(orderID)})
}
