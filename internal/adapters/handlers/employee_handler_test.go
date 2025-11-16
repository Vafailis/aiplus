package handlers

import (
	"aiplus_golang/internal/core/domain"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockEmployeeService struct {
	createErr error
}

func (m *mockEmployeeService) CreateEmployee(ctx context.Context, emp *domain.Employee) error {
	return m.createErr
}

func TestCreateEmployee_Success(t *testing.T) {
	service := &mockEmployeeService{}
	handler := NewEmployeeHandler(service)

	emp := domain.Employee{
		FullName: "Test Name",
		Phone:    "+123456789",
		Address:  "Test Address",
		Email:    "test@example.com",
	}
	body, _ := json.Marshal(emp)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateEmployee(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestCreateEmployee_InvalidJSON(t *testing.T) {
	service := &mockEmployeeService{}
	handler := NewEmployeeHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader([]byte("{invalid json")))
	w := httptest.NewRecorder()

	handler.CreateEmployee(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCreateEmployee_InvalidEmployee(t *testing.T) {
	service := &mockEmployeeService{createErr: domain.ErrInvalidEmployee}
	handler := NewEmployeeHandler(service)

	emp := domain.Employee{}
	body, _ := json.Marshal(emp)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateEmployee(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCreateEmployee_InternalError(t *testing.T) {
	service := &mockEmployeeService{createErr: errors.New("some error")}
	handler := NewEmployeeHandler(service)

	emp := domain.Employee{
		FullName: "Test Name",
		Phone:    "+123456789",
		Address:  "Test Address",
		Email:    "test@example.com",
	}
	body, _ := json.Marshal(emp)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateEmployee(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestCreateEmployee_MethodNotAllowed(t *testing.T) {
	service := &mockEmployeeService{}
	handler := NewEmployeeHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	w := httptest.NewRecorder()

	handler.CreateEmployee(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}
