package ports

import "net/http"

type EmployeeHandler interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request)
}
