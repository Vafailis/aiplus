package services

import (
	"aiplus_golang/internal/core/domain"
	"context"
	"errors"
	"testing"
)

type mockEmployeeRepository struct {
	saveErr error
}

func (m *mockEmployeeRepository) Save(ctx context.Context, emp *domain.Employee) error {
	return m.saveErr
}

func TestCreateEmployee_Success(t *testing.T) {
	repo := &mockEmployeeRepository{}
	service := NewEmployeeService(repo)

	emp := &domain.Employee{
		FullName: "Test Name",
		Phone:    "+123456789",
		Address:  "Test Address",
		Email:    "test@example.com",
	}

	err := service.CreateEmployee(context.Background(), emp)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestCreateEmployee_InvalidEmployee(t *testing.T) {
	repo := &mockEmployeeRepository{}
	service := NewEmployeeService(repo)

	emp := &domain.Employee{} // все поля пустые

	err := service.CreateEmployee(context.Background(), emp)
	if err != domain.ErrInvalidEmployee {
		t.Errorf("expected ErrInvalidEmployee, got %v", err)
	}
}

func TestCreateEmployee_SaveError(t *testing.T) {
	repo := &mockEmployeeRepository{saveErr: errors.New("db error")}
	service := NewEmployeeService(repo)

	emp := &domain.Employee{
		FullName: "Test Name",
		Phone:    "+123456789",
		Address:  "Test Address",
		Email:    "test@example.com",
	}

	err := service.CreateEmployee(context.Background(), emp)
	if err == nil || err.Error() != "db error" {
		t.Errorf("expected db error, got %v", err)
	}
}

func TestValidateEmployee(t *testing.T) {
	service := &EmployeeService{}

	valid := &domain.Employee{
		FullName: "Test Name",
		Phone:    "+123456789",
		Address:  "Test Address",
		Email:    "test@example.com",
	}
	if service.validateEmployee(valid) {
		t.Error("expected valid employee to pass validation")
	}

	invalid := &domain.Employee{}
	if !service.validateEmployee(invalid) {
		t.Error("expected invalid employee to fail validation")
	}
}
