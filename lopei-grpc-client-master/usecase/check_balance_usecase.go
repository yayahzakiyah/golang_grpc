package usecase

import "lopei-grpc-client/repository"

type CheckBalanceUseCase interface {
	GetBalance(lopeId int32) (float32, error)
}

type checkBalanceUseCase struct {
	repo repository.CustomerRepository
}

func (c *checkBalanceUseCase) GetBalance(lopeId int32) (float32, error) {
	return c.repo.CheckBalance(lopeId)
}

func NewCheckBalanceUseCase(repo repository.CustomerRepository) CheckBalanceUseCase {
	return &checkBalanceUseCase{repo: repo}
}