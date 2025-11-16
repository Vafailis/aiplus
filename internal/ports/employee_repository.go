package ports

import (
	"aiplus_golang/internal/core/domain"
	"context"
)

type EmployeeRepository interface {
	Save(ctx context.Context, emp *domain.Employee) error
}
