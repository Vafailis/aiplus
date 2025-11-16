package ports

import (
	"aiplus_golang/internal/core/domain"
	"context"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, emp *domain.Employee) error
}
