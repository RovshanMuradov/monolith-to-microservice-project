package products

import "errors"

var ErrNotFound = errors.New("product not found")

type Repository interface {
	Save(*product) error
	ByID(ID) (*Product, error)
}
