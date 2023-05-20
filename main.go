package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartItem struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func (item *CartItem) CalculateTotalPricePerItem() float64 {
	return item.Price * float64(item.Quantity)
}

type Cart struct {
	Items []CartItem
}

func (c *Cart) CalculateTotalPrice() float64 {
	totalPrice := 0.0
	for _, item := range c.Items {
		totalPrice += item.CalculateTotalPricePerItem()
	}
	return totalPrice
}

func NewCart() *Cart {
	return &Cart{
		Items: []CartItem{
			{ID: 1, Name: "Product 1", Price: 10.0, Quantity: 2},
			{ID: 2, Name: "Product 2", Price: 15.0, Quantity: 1},
		},
	}
}

func totalPriceHandler(ctx *gin.Context) {
	cart := NewCart()
	totalPrice := cart.CalculateTotalPrice()
	ctx.JSON(http.StatusOK, gin.H{"total price": totalPrice})
}

func main() {
	router := gin.New()
	router.GET("/total", totalPriceHandler)
	router.Run(":9090")
}
