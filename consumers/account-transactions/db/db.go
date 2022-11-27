package db

import (
	"account-transactions/infrastructure/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type AccountRepository interface {
	Save(account Account) error
	IncreaseBalance(id string, amount uint64) error
	DecreaseBalance(id string, amount uint64) error
	SetInactive(id string) error
	Delete(id string) error
}

type NoOpAccountRepository struct {
}

func (m *NoOpAccountRepository) SetInactive(_ string) error {
	return nil
}

func (m *NoOpAccountRepository) IncreaseBalance(_ string, _ uint64) error {
	return nil
}

func (m *NoOpAccountRepository) DecreaseBalance(_ string, _ uint64) error {
	return nil
}

func (m *NoOpAccountRepository) Save(_ Account) error {
	return nil
}

func (m *NoOpAccountRepository) Delete(string) error {
	return nil
}

type PromAccountRepository struct {
	AccountRepository
}

func (s *PromAccountRepository) Save(account Account) error {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		metrics.TotalDbAccessDuration.WithLabelValues("write", "AccountRepository.Save").Observe(f)
	}))
	defer timer.ObserveDuration()
	return s.AccountRepository.Save(account)
}

func (s *PromAccountRepository) IncreaseBalance(id string, amount uint64) error {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		metrics.TotalDbAccessDuration.WithLabelValues("write", "AccountRepository.IncreaseBalance").Observe(f)
	}))
	defer timer.ObserveDuration()
	return s.AccountRepository.IncreaseBalance(id, amount)
}

func (s *PromAccountRepository) DecreaseBalance(id string, amount uint64) error {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		metrics.TotalDbAccessDuration.WithLabelValues("write", "AccountRepository.DecreaseBalance").Observe(f)
	}))
	defer timer.ObserveDuration()
	return s.AccountRepository.DecreaseBalance(id, amount)
}

func (s *PromAccountRepository) SetInactive(id string) error {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		metrics.TotalDbAccessDuration.WithLabelValues("write", "AccountRepository.SetInactive").Observe(f)
	}))
	defer timer.ObserveDuration()
	return s.AccountRepository.SetInactive(id)
}

func (s *PromAccountRepository) Delete(id string) error {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		metrics.TotalDbAccessDuration.WithLabelValues("write", "AccountRepository.Delete").Observe(f)
	}))
	defer timer.ObserveDuration()
	return s.AccountRepository.Delete(id)
}

func NewPromAccountRepository(accountRepository AccountRepository) *PromAccountRepository {
	return &PromAccountRepository{AccountRepository: accountRepository}
}
