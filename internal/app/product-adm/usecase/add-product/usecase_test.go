package addproduct_test

import (
	"testing"

	"github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/repository"
	addproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/usecase/add-product"
	"gorm.io/gorm"
)

func TestAddProduct(t *testing.T) {
	tests := []struct {
		name     string
		input    addproduct.AddProductInputDTO
		expected bool
		err      error
	}{{
		name: "should add a product",
		input: addproduct.AddProductInputDTO{
			Name:          "teste",
			Description:   "teste",
			PurchasePrice: 10,
			Stock:         10,
		},
		expected: false,
		err:      nil,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testRepository := repository.NewProductRepository(&gorm.DB{})
			useCaseTest := addproduct.NewAddProductUseCase(testRepository)
			err := useCaseTest.Execute(tt.input)
			if (err != nil) != tt.expected {
				t.Fatalf("AddProduct() error = %v, wantErr %v", err, tt.expected)
			}
			if err == nil {
				t.Fatalf("AddProduct() did not return an error but expected one: %s", tt.err)
			}
		})
	}
}
