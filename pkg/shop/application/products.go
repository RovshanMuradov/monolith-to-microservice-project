package application

import (
	"monolith-t0-microservice-project/pkg/common/price"
	products "monolith-t0-microservice-project/pkg/shop/domain"
)

type productReadModel interface {
	AllProducts() ([]products.Product, err)
}

type ProductsService struct {
	repo      products.Repository
	readModel productReadModel
}

func NewProductsService(repo products.Repository, readModel productReadModel) ProductsService {

	return ProductsService{repo, readModel}
}

func (s ProductsService) AllProducts() ([]products.Product, error) {

	return s.readModel.AllProducts()
}

type AddProductCommand struct {
	ID          string
	Name        string
	Description string
	PriceCents  uint
	Price       string
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {

	price, err := price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)
	if err != nil {
		return errors.Wrap(err, "invalid product price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)

	if err != nil {
		return errors.Wrap(err, "can not create product")
	}

	if err := s.repo.Save(p); err != nil {
		return errors.Wrap(err, "can not save product")
	}

	return nil
}
