package service

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"rajan/employee/data"
	"rajan/employee/model"
)

var mu sync.Mutex // Mutex for synchronizing access to employees slice

// GetEmployees responds with the list of all employee as JSON.
func GetEmployees(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Page should be a valid number"})
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Limit should be a valid number"})
		return
	}

	startIndex := page * limit
	endIndex := (page + 1) * limit
	if len(data.Employees) >= endIndex {
		c.IndentedJSON(http.StatusOK, data.Employees[startIndex:endIndex])
		return
	}

	remainingRecord := len(data.Employees) - startIndex
	if remainingRecord > 0 {
		endIndex = startIndex + remainingRecord
		c.IndentedJSON(http.StatusOK, data.Employees[startIndex:endIndex])
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
}

// Adds a new employee to the database or store.
func CreateEmployee(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var emp model.Employee
	if err := c.BindJSON(&emp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	data.Employees = append(data.Employees, emp)
	c.IndentedJSON(http.StatusCreated, emp)
}

// Retrieves an employee from the database or store by ID.
func GetEmployeeByID(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id := c.Param("id")
	for _, emp := range data.Employees {
		if emp.ID == id {
			c.IndentedJSON(http.StatusOK, emp)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
}

// Updates the details of an existing employee.
func UpdateEmployee(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id := c.Param("id")
	var emp model.Employee
	if err := c.BindJSON(&emp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	for index, e := range data.Employees {
		if e.ID == id {
			data.Employees[index] = emp
			c.IndentedJSON(http.StatusOK, emp)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
}

// Deletes an employee from the database or store by ID.
func DeleteEmployee(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id := c.Param("id")
	for index, emp := range data.Employees {
		if emp.ID == id {
			data.Employees = append(data.Employees[:index], data.Employees[index+1:]...)
			c.IndentedJSON(http.StatusNoContent, nil)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
}
