package repositories

import (
	domain "aiplus_golang/internal/core/domain"
	ports "aiplus_golang/internal/ports"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewEmployeeRepository(db *pgxpool.Pool) ports.EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func NewMockEmployeeRepository() ports.EmployeeRepository {
	return &EmployeeRepository{}
}

type EmployeeRepository struct {
	db *pgxpool.Pool
}

func (r *EmployeeRepository) Save(ctx context.Context, emp *domain.Employee) error {
	query := `INSERT INTO employees (FullName, Email, Phone, Address) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, emp.FullName, emp.Email, emp.Phone, emp.Address)
	return err
}
