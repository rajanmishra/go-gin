package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"rajan/employee/data"
	"rajan/employee/model"
	"rajan/employee/service"
)

func TestCreateEmployee(t *testing.T) {

	router := gin.Default()
	router.POST("/employee", service.CreateEmployee)

	// Create a new employee to send in the request
	newEmployee := model.Employee{
		ID:       "5",
		Name:     "Test Employee",
		Position: "Engineer 1",
		Salary:   30.0,
	}

	// Marshal the new employee into JSON format
	payload, _ := json.Marshal(newEmployee)

	// Create a POST request with the JSON payload
	req, err := http.NewRequest("POST", "/employee", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	// Set the request Content-Type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Check the status code and response body
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), newEmployee.Name)
}

func TestGetEmployeeByID(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	router.GET("/employee/:id", service.GetEmployeeByID)

	// Create a new request with a specific employee ID
	id := "1"
	req, err := http.NewRequest("GET", "/employee/"+id, nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, rec.Code)

	var emp model.Employee
	err = json.Unmarshal(rec.Body.Bytes(), &emp)
	assert.NoError(t, err)
	assert.Equal(t, id, emp.ID)
}

func TestGetEmployees(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	router.GET("/employee", service.GetEmployees)

	// Create a GET request with page and limit query parameters
	req, err := http.NewRequest("GET", "/employee?page=0&limit=2", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, rec.Code)

	var empList []model.Employee
	err = json.Unmarshal(rec.Body.Bytes(), &empList)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(empList)) // Check if 2 employees are returned based on the limit
}

func TestUpdateEmployee(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	router.PUT("/employee/:id", service.UpdateEmployee)

	// Create an updated employee data
	updatedEmployee := model.Employee{
		ID:       "1",
		Name:     "A Updated",
		Position: "Engineer 1 Updated",
		Salary:   55.0,
	}

	// Marshal the updated employee into JSON format
	payload, _ := json.Marshal(updatedEmployee)

	// Create a PUT request with the JSON payload
	req, err := http.NewRequest("PUT", "/employee/1", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	// Set the request Content-Type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, rec.Code)

	var emp model.Employee
	err = json.Unmarshal(rec.Body.Bytes(), &emp)
	assert.NoError(t, err)
	assert.Equal(t, updatedEmployee, emp) // Check if the employee was updated correctly
}

func TestDeleteEmployee(t *testing.T) {
	// Create a new Gin router
	var initial = len(data.Employees)
	router := gin.Default()
	router.DELETE("/employee/:id", service.DeleteEmployee)

	// Create a DELETE request for deleting the employee with ID 1
	req, err := http.NewRequest("DELETE", "/employee/1", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Check the status code and response body
	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Equal(t, initial-1, len(data.Employees)) // Check if the employee was deleted from the slice
}
