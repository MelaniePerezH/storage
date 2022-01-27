package product

import (
	"context"
	"errors"

	"github.com/extlurosell/meli_bootcamp_go_w3-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("product not found")
)

type Service interface {
	Get(ctx context.Context, name string) (domain.Product, error)
	Save(ctx context.Context, name, productType string, count int, price int) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Update(ctx context.Context, p domain.Product) (domain.Product, error)
}

type service struct {
	repository Repository
}

func (s *service) Get(ctx context.Context, name string) (domain.Product, error) {
	product, err := s.repository.Get(ctx, name)
	if err != nil {
		return domain.Product{}, errors.New("product doesn't exists")
	}
	return product, nil
}
func (s *service) Save(ctx context.Context, name, productType string, count int, price int) (domain.Product, error) {
	p := domain.Product{}
	p.Name = name
	p.Tipo = productType
	p.Count = count
	p.Price = price

	return s.repository.Store(ctx, p)
}

func (r *service) GetAll() ([]domain.Product, error) {
	return r.repository.GetAll()
}
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
func (s *service) Update(ctx context.Context, p domain.Product) (domain.Product, error) {
	producto, err := s.repository.Get(ctx, p.Name)
	if err != nil {
		return domain.Product{}, errors.New("product doesn't exists'")
	}

	return producto, s.repository.Update(ctx, producto)
}
