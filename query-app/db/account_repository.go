package db

import "query-app/db/entity"

type AccountRepository interface {
	Get(id string) (entity.Account, error)
}
