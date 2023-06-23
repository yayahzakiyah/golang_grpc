package manager

import "lopei-grpc-client/usecase"

type UseCaseManager interface {
	CheckBalanceUseCase() usecase.CheckBalanceUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CheckBalanceUseCase() usecase.CheckBalanceUseCase {
	return usecase.NewCheckBalanceUseCase(u.repoManager.CustomerRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}