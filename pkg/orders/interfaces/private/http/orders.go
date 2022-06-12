package http

import (
	"github.com/go-chi/chi"
	"monolith-t0-microservice-project/pkg/orders/application"
	"monolith-t0-microservice-project/pkg/orders/domain/orders"
	"net/http"
)

func AddRoutes(router *chi.Mux, service application.OrdersService, repository orders.Repository) {
	resource := ordersResource{service, repository}
	router.Post("/orders/{id}/paid", resource.PostPaid)
}

type ordersResource struct {
	service    application.OrdersSerivce
	repository orders.Repository
}

func (o ordersResource) PostPaid(w http.ResponseWriter, r *http.Request) {
	cmd := application.MarkOrderAsPaidCommand{
		OrderID: orders.ID(chi.URLParam(r, "id")),
	}

	if err := o.service.MarkOrderAsPaid(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
