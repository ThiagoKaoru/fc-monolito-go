package productdomain

import (
	baseEntity "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/entity"
)

type Product struct {
	baseEntity.BaseEntity `gorm:"embedded"`
	Name                  string
	Description           string
	PurchasePrice         uint
	Stock                 int
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

func (p *Product) GetPurchasePrice() uint {
	return p.PurchasePrice
}

func (p *Product) GetStock() int {
	return p.Stock
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) SetDescription(description string) {
	p.Description = description
}

func (p *Product) SetPurchasePrice(price uint) {
	p.PurchasePrice = price
}

func (p *Product) SetStock(stock int) {
	p.Stock = stock
}
