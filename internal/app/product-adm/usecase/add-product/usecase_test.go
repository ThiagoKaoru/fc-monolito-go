package addproduct

import (
	"testing"

	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/domain"
)

func TestAddProduct(t *testing.T) {
	tests := []struct {
		name     string
		input    *productDomain.Product
		expected bool
		err      error
	}{{
		name: "should add a product",
		input: &productDomain.Product{
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
			err := Add(*tt.input)
			if (err != nil) != tt.expected {
				t.Fatalf("AddProduct() error = %v, wantErr %v", err, tt.expected)
			}
			if err == nil {
				t.Fatalf("AddProduct() did not return an error but expected one: %s", tt.err)
			}
		})
	}
}
