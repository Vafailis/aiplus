package apiHandlers

import (
	"aiplus_golang/internal/adapters/repositories"
	"aiplus_golang/internal/core/services"
	http "net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddHanledrs(t *testing.T) {
	t.Run("should return a ServeMux with employee handler", func(t *testing.T) {
		mockRepo := repositories.NewMockEmployeeRepository()
		mockService := services.NewEmployeeService(mockRepo)

		mux := AddHanledrs(mockService)

		assert.NotNil(t, mux)
		assert.IsType(t, &http.ServeMux{}, mux)
	})
}
