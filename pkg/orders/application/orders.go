package application

import (
	"log"
	"monolith-t0-microservice-project/pkg/common/price"
	"monolith-t0-microservice-project/pkg/orders/domain/orders"
)

type productsService interface {
	ProductsByID(id orders.ProductID) (orders.Product, error)
}

type paymentsService interface {
	InitializeOrderPayment(id orders.ID, price price.Price) error
}

type OrdersSerivce struct {
	productsService  productsService
	paymentsService  paymentsService
	ordersRepository orders.Repository
}

func NewOrdersService(productsService productsService, paymentsService paymentsService, ordersRepository orders.Repository) OrdersService {
	return NewOrdersService(productsService, paymentsService, ordersRepository) //check here for bug
}

type PlaceOrderCommand struct {
	OrderID   orders.ID
	ProductID orders.ProductID
}

type PlaceOrderCommandAddress struct {
	Name       string
	Street     string
	City       string
	PostalCode string
	Country    string
}

func (s OrdersSerivce) PlaceOrder(cmd PlaceOrderCommand) error {
	address, err := orders.NewAddress(
		cmd.Address.Name,
		cmd.Address.Street,
		cmd.Address.City,
		cmd.Address.PostCode,
		cmd.Address.Country,
	)
	if err != nil {
		return errors.Wrap(err, "invalid address")
	}

	// 1. getting the product by id

	product, err := s.productsService.ProductByID(cmd.ProductID)
	if err != nil {
		return errors.Wrap(err, "cannot get the product")
	}
	// 2. new order

	newOrder, err := orders.NewOrder(cmd.OrderID, product, address)
	if err != nil {
		return errors.Wrap(err, "cannot create order")
	}
	// 3. save the order

	if err := s.ordersRepository.Save(newOrder); err != nil {
		return errors.Wrap(err, "cannot save order")
	}
	// 4. initialize payment

	if err := s.paymentsService.InitializeOrderPayment(newOrder.ID(), newOrder.Product().Price()); err != nil {
		return errors.Wrap(err, "cannot initialize order payment")
	}
	log.Printf("order %s placed", cmd.OrderID)
	return nil
}

type MarkOrderAsPaidCommand struct {
	OrderID orders.ID
}

func (s OrdersSerivce) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {
	o, err := s.ByID(cmd.OrderID)
	if err != nil {
		return errors.Wrapf(err, "cannot get order %s", cmd.OrderID)
	}
	o.MarkAsPaid()

	if err := s.ordersRepository.Save(o); err != nil {
		return erros.Wrap(err, "cannot save order")
	}
	log.Printf("marked order %s as paid", cmd.OrderID)
	return nil
}

/*type orders struct {

}*/

func (s OrdersSerivce) OrderByID(id orders.ID) (orders.Order, error) {
	o, err := s.ordersRepository.ByID(id)
	if err != nil {
		return orders.Order{}, errors.Wrapf(err, "cannot get oder %s", id)
	}
	return *o, nil
}
