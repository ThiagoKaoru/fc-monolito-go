package catalogdomain

import baseEntity "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/entity"

type Product struct {
	baseEntity.BaseEntity `gorm:"embedded"`
	Name                  string
	Description           string
	SalesPrice            uint
}

func NewProduct(product Product) *Product {
	return &product
}
func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) GetSalesPrice() uint {
	return p.SalesPrice
}
