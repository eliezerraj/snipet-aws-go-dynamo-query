package service

import (
	"context"

)

type Repository interface {
	AddInvoice(context.Context, interface{}) error
	QueryInvoice(context.Context, string, string) (interface{}, error) 
	QueryInvoiceGsi(context.Context, string, string) (interface{}, error) 
	UpdateInvoiceTransaction(context.Context, string, string, string ,float32) (interface{}, error) 
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddInvoice(ctx context.Context, item interface{}) error {
	return s.repo.AddInvoice(ctx, item)
}

func (s *Service) QueryInvoice(ctx context.Context, pk string, sk string) (interface{}, error)  {
	return s.repo.QueryInvoice(ctx, pk, sk)
}

func (s *Service) QueryInvoiceGsi(ctx context.Context, pk string, sk string) (interface{}, error)  {
	return s.repo.QueryInvoiceGsi(ctx, pk, sk)
}

func (s *Service) UpdateInvoiceTransaction(ctx context.Context, pk string, sk string, order_id string ,amount float32)  (interface{}, error)  {
	return s.repo.UpdateInvoiceTransaction(ctx, pk, sk, order_id, amount)
}