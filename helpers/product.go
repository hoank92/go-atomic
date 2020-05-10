package helpers

import (
	"go-atomic/main/models"
	"sync/atomic"
)

var BaseProduct *models.Product

var quantity *int64

func InitData() {
	BaseProduct = &models.Product{Id: 10, Name: "hoankTest", Quantity: 10000000, Version: 0}
	quantity = &BaseProduct.Quantity
}

func GetProduct() *models.Product {
	BaseProduct.Quantity = atomic.LoadInt64(quantity)
	return BaseProduct
}

func DecQuantity() {
	atomic.AddInt64(quantity, -1)
}
