package main

import (
	"fmt"
	"log"
	"monolith-t0-microservice-project/pkg/common/cmd"
	"os"
)

func main(){
	log.Println("starting payments microservice")

	defer log.Println("closing payments microservice")

	ctx := cmd.Contet()

	paymentsInterface := createPaymentsMicroservice()
	if er := paymentsInterface.Run(ctx); err!=nil{
		panic(err)
	}
}

func createPaymentsMicroservice() amqp.paymentsInterface{
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)

	paymentsInterface, err := amqp.NewPaymentsInferface(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR"))
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		paymentsService,
		)

	if err != nil{
		panic(err)
	}

	return paymentsInterface
}