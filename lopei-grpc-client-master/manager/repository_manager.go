package manager

import "lopei-grpc-client/repository"

type RepositoryManager interface {
	CustomerRepo() repository.CustomerRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infraManager.LogicClientConn())
}

func NewRepositoryManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{infraManager: infraManager}
}
