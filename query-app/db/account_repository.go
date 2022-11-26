package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"query-app/db/entity"
	"query-app/infrastructure/metrics"
)

type AccountRepository interface {
	Get(id string) (entity.Account, error)
}

type SpanAccountRepository struct {
	AccountRepository
}

func NewSpanAccountRepository(accountRepository AccountRepository) *SpanAccountRepository {
	return &SpanAccountRepository{AccountRepository: accountRepository}
}

func (s *SpanAccountRepository) Get(id string) (entity.Account, error) {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		metrics.TotalDbAccessDuration.WithLabelValues("read", "AccountRepository.Get").Observe(f)
	}))
	defer timer.ObserveDuration()
	return s.AccountRepository.Get(id)
}
