package aggregate

import (
	"github.com/google/uuid"
	"github.com/offlinebrain/cqrs-shop/command-app/domain/entity"
	vo "github.com/offlinebrain/cqrs-shop/command-app/domain/valueobject"
)

type Product struct {
	item   entity.Item
	price  vo.Price
	amount vo.Amount
	status vo.ProductStatus
}

func NewProduct(name vo.Name) *Product {
	return &Product{
		item: entity.Item{
			ID: vo.ID(uuid.New()), Name: name,
		},
		amount: 0,
	}
}

func (p *Product) ChangePrice(price vo.Price) {
	p.price = price
}

func (p *Product) IncreaseAmount(amount vo.Amount) {
	if amount < 1 {
		return
	}
	p.amount += amount
}

func (p *Product) RemoveItems(amount vo.Amount) {
	if p.amount < amount {
		return
	}
	p.amount -= amount
}

func (p *Product) StartSales() {
	p.status = vo.OnSale
}

func (p *Product) StopSales() {
	p.status = vo.Unavailable
}

func (p *Product) Status() vo.ProductStatus {
	if p.status == vo.Unavailable {
		return vo.Unavailable
	} else if p.amount == 0 {
		return vo.SoldOut
	}
	return vo.OnSale
}
