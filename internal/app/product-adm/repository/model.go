package productadm_repository

import (
	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/domain"
)

type Product struct {
	productDomain.Product `gorm:"embedded"`
}

func (p *Product) TableName() string {
	return "products"
}
