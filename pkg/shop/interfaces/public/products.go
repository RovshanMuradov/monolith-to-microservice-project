package http

import (
	"github.com/go-chi/chi"
	"monolith-t0-microservice-project/pkg/common/price"
	products "monolith-t0-microservice-project/pkg/shop/domain"
	"net/http"
)

func AddRoutes(router *chi.Mux, productsReadModel productsReadModel) {
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

type productsReadModel interface {
	AllProducts() ([]products.Product, error)
}

type productsResource struct {
	readModel productsReadModel
}

type productsView struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       priceView `json:"price"`
}

type priceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func priceViewFromPrice(p price.Price) priceView {
	return priceView{p.Cents(), p.Currency()}
}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request) {

	products, err := p.readModel.AllProducts()
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	view := []productsView{}
	for _, product := range products {
		view = append(view, productsView{
			string(product.ID()),
			product.Name(),
			product.Description(),
			priceViewFromPrice(product.Price()),
		})
	}
	render.Respond(w, r, view)
}
