package productrepository

import (
	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/domain"
)

type Product struct {
	productDomain.Product `gorm:"embedded"`
}

func (p *Product) TableName() string {
	return "products"
}
