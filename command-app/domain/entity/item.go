package entity

import vo "github.com/offlinebrain/cqrs-shop/command-app/domain/valueobject"

type Item struct {
	ID   vo.ID
	Name vo.Name
}
