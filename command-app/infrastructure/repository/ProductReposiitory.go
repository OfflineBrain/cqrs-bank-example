package repository

import "github.com/offlinebrain/cqrs-shop/command-app/domain/aggregate"

type ProductRepository interface {
	Get(id string) (aggregate.Product, error)
	Update(product aggregate.Product) error
}
