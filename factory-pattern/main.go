package main

import (
	"fmt"

	"github.com/mrrahbarnia/Design-Patterns-In-Golang/factory-pattern/product"
)

func main() {
	factory := product.Product{}
	product := factory.New()
	fmt.Println("The product was created at", product.CreatedAt.UTC())

}
