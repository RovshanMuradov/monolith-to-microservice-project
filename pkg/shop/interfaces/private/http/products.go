package http

import (
	"github.com/go-chi/chi"
	"monolith-t0-microservice-project/pkg/common/price"
	"net/http"
)

func AddRoutes(router *chi.Mux, repo products_domain.Repository) {
	resource := productsResource{repo}
	router.Get("/products/{id}", resource.Get)
}

type productsResource struct {
	repo products_domain.Repository
}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func priceViewFromPrice(p price.Price) PriceView {
	return PriceView{p.Cents(), p.Currency()}
}

func (p productsResource) Get(w http.ResponseWriter, r *http.Request) {

	products, err := p.repo.ByID(products_domain.ID(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	render.Respond(w, r, ProductView{
		string(product.ID()),
		product.Name(),
		product.Description(),
		priceViewFromPrice(product.Price()),
	})
}
