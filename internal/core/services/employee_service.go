package services

import (
	domain "aiplus_golang/internal/core/domain"
	ports "aiplus_golang/internal/ports"
	"context"
)

type EmployeeService struct {
	repo ports.EmployeeRepository
}

func NewEmployeeService(repo ports.EmployeeRepository) ports.EmployeeService {
	return &EmployeeService{repo: repo}
}

func (service *EmployeeService) CreateEmployee(ctx context.Context, emp *domain.Employee) error {
	if service.validateEmployee(emp) {
		return domain.ErrInvalidEmployee
	}
	return service.repo.Save(ctx, emp)
}

func (service *EmployeeService) validateEmployee(emp *domain.Employee) bool {
	if emp.FullName == "" || emp.Phone == "" || emp.Address == "" || emp.Email == "" {
		return true
	}
	return false
}
