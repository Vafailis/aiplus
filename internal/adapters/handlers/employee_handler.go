package handlers

import (
	domain "aiplus_golang/internal/core/domain"
	ports "aiplus_golang/internal/ports"
	"encoding/json"
	"net/http"
)

func NewEmployeeHandler(service ports.EmployeeService) ports.EmployeeHandler {
	return &EmployeeHandler{service: service}
}

type EmployeeHandler struct {
	service ports.EmployeeService
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	domainEmp := &domain.Employee{}

	if err := json.NewDecoder(r.Body).Decode(&domainEmp); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateEmployee(r.Context(), domainEmp); err != nil {
		if err == domain.ErrInvalidEmployee {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee created"})
}
