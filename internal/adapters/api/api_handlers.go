package apiHandlers

import (
	"aiplus_golang/internal/adapters/handlers"
	"aiplus_golang/internal/ports"
	http "net/http"
)

func AddHanledrs(service ports.EmployeeService) *http.ServeMux {
	handler := handlers.NewEmployeeHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/employees", handler.CreateEmployee)

	return mux
}
